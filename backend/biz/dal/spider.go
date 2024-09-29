package dal

import (
	"backend/biz/dal/model"
	"backend/consts"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

// GetAllMedals 获取所有奖牌信息
func GetAllMedals() (*model.OlympicsData, error) {
	pageContent, err := fetchPageContent(consts.Url)
	if err != nil {
		return nil, err
	}

	data, err := parseOlympicsData(pageContent)
	return data, err
}

// 发起HTTP请求并返回页面内容
func fetchPageContent(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// parseOlympicsData 解析页面内容并构建 OlympicsData 对象
func parseOlympicsData(pageContent string) (*model.OlympicsData, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageContent))
	if err != nil {
		return nil, err
	}

	// 查找包含数据的 <script> 标签
	scriptTag := doc.Find("script#__NEXT_DATA__").First()
	if scriptTag.Length() == 0 {
		return nil, fmt.Errorf("no script tag with JSON data found")
	}

	jsonData := scriptTag.Text()

	// 将 jsonData 解析为 map 类型
	var jsonMap map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &jsonMap)
	if err != nil {
		return nil, fmt.Errorf("failed to parse json data: %v", err)
	}

	// 导航至 props -> pageProps -> initialMedals -> medalStandings -> medalsTable
	medalsTable, ok := jsonMap["props"].(map[string]interface{})["pageProps"].(map[string]interface{})["initialMedals"].(map[string]interface{})["medalStandings"].(map[string]interface{})["medalsTable"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to extract medals table from JSON data")
	}

	// 使用 map 存储国家对应的 MedalEntry
	countryMedalsMap := make(map[string]model.MedalEntry)

	for _, entry := range medalsTable {
		entryMap := entry.(map[string]interface{})

		// 获取国家名称
		country := entryMap["description"].(string)

		// 初始化或获取国家对应的 MedalEntry
		medalEntry, exists := countryMedalsMap[country]
		if !exists {
			medalEntry = model.MedalEntry{
				Description: country,
				Disciplines: []model.Discipline{},
			}
		}

		// 获取 disciplines 数据
		disciplines := entryMap["disciplines"].([]interface{})
		for _, discipline := range disciplines {
			disciplineMap := discipline.(map[string]interface{})

			// 创建 Discipline 对象
			disciplineEntry := model.Discipline{
				Name:   disciplineMap["name"].(string),
				Total:  int(disciplineMap["total"].(float64)),
				Gold:   int(disciplineMap["gold"].(float64)),
				Silver: int(disciplineMap["silver"].(float64)),
				Bronze: int(disciplineMap["bronze"].(float64)),
			}

			// 添加学科条目到该国家的 MedalEntry
			medalEntry.Disciplines = append(medalEntry.Disciplines, disciplineEntry)
		}

		// 更新国家的 MedalEntry 到 map 中
		countryMedalsMap[country] = medalEntry
	}

	// 构建 OlympicsData 对象，将 map 中的 MedalEntry 转换为 MedalsTable
	var data model.OlympicsData
	for _, medalEntry := range countryMedalsMap {
		data.Props.InitialMedals.MedalsTable = append(data.Props.InitialMedals.MedalsTable, medalEntry)
	}

	return &data, nil
}

// DailyEvents 每日赛程
func DailyEvents(date string) ([]*model.Event, error) {
	url := "https://sph-s-api.olympics.com/summer/schedules/api/CHI/schedule/day/2024-" + date // 替换为目标网址
	data, err := fetchPageContent(url)
	if err != nil {
		return nil, err
	}
	return parseDailyEvents(data)
}

func parseDailyEvents(data string) ([]*model.Event, error) {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(data), &jsonMap)
	if err != nil {
		return nil, fmt.Errorf("failed to parse json data: %v", err)
	}

	groups, ok := jsonMap["units"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to assert data")
	}

	var events []*model.Event // 用于存储所有事件信息

	for _, group := range groups {
		groupMap, ok := group.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("failed to assert group data")
		}

		// 获取所需字段
		phaseName := groupMap["phaseName"].(string)
		id := groupMap["id"].(string)
		disciplineName := groupMap["disciplineName"].(string)
		startDate := groupMap["startDate"].(string)

		// 获取竞争者信息
		var competitors []*model.Competitor
		competitorsData, ok := groupMap["competitors"].([]interface{})
		if ok {
			for _, competitor := range competitorsData {

				competitorMap, ok := competitor.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("failed to assert competitor data")
				}
				name := competitorMap["name"].(string)
				flag := competitorMap["noc"].(string)
				results, ok := competitorMap["results"].(map[string]interface{})
				if !ok {
					continue
				}
				winnerLoserTie, ok := results["winnerLoserTie"].(string)
				if !ok {
					continue
				}
				rating := results["mark"].(string)

				competitors = append(competitors, &model.Competitor{ // 使用指针
					Name:           name,
					WinnerLoserTie: winnerLoserTie,
					Rating:         rating,
					Flag:           flag,
				})
			}
		}

		// 创建一个事件结构体并添加到事件列表
		events = append(events, &model.Event{ // 使用指针
			PhaseName:      phaseName,
			ID:             id,
			DisciplineName: disciplineName,
			Competitors:    competitors,
			StartDate:      startDate,
		})
	}

	return events, nil
}

func EventTypeList() ([]*model.EventList, error) {
	data, err := fetchPageContent(consts.Url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch page")
	}
	return extractEvents(data)
}

func extractEvents(data string) ([]*model.EventList, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	// 查找包含数据的 <script> 标签
	scriptTag := doc.Find("script#__NEXT_DATA__").First()
	if scriptTag.Length() == 0 {
		return nil, fmt.Errorf("no script tag with JSON data found")
	}

	jsonData := scriptTag.Text()

	// 将 jsonData 解析为 map 类型
	var jsonMap map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &jsonMap)
	if err != nil {
		return nil, fmt.Errorf("failed to parse json data: %v", err)
	}

	medalsTable, ok := jsonMap["props"].(map[string]interface{})["pageProps"].(map[string]interface{})["initialMedals"].(map[string]interface{})["medalStandings"].(map[string]interface{})["medalsTable"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to extract medals table from JSON data")
	}

	// 用于存储最终的事件列表
	var eventLists []*model.EventList

	for _, entry := range medalsTable {
		entryMap := entry.(map[string]interface{})

		// 获取 disciplines 数据
		disciplines := entryMap["disciplines"].([]interface{})
		for _, discipline := range disciplines {
			disciplineMap := discipline.(map[string]interface{})

			// 获取 eventCode 和 name
			disciplineName := disciplineMap["name"].(string)
			medalWinners, ok := disciplineMap["medalWinners"].([]interface{})
			if !ok {
				continue
			}

			// 构造事件类型列表
			var eventTypeList model.EventTypeList

			for _, winner := range medalWinners {
				winnerMap := winner.(map[string]interface{})
				eventCode := winnerMap["eventCode"].(string)
				description := winnerMap["eventDescription"].(string)

				// 构造所需的 map 数据
				event := map[string]string{
					"name": description,
					"id":   eventCode,
				}

				// 添加到事件类型列表
				eventTypeList = append(eventTypeList, event)
			}

			// 创建一个新的 EventList 并添加到事件列表数组中
			eventList := &model.EventList{
				Name: disciplineName,
				List: eventTypeList,
			}
			eventLists = append(eventLists, eventList)
		}
	}

	// 返回事件列表数组
	return eventLists, nil
}

func EventTable(eventID string) ([]*model.EventTable, error) {
	url := "https://olympics.com/OG2024/data/GLO_EventGames~comp=OG2024~event=" + eventID + "~lang=CHI.json"
	data, err := fetchPageContent(url)
	if err != nil {
		return nil, err
	}
	return parseEvents(data, url)
}

func parseEvents(data string, url string) ([]*model.EventTable, error) {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(data), &jsonMap)
	if err != nil {
		return nil, fmt.Errorf("failed to parse json data: %v, url:%v", err, url)
	}

	// 获取事件数据
	events, ok := jsonMap["event"].(map[string]interface{})["phases"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to assert data")
	}

	var parsedEvents []*model.EventTable

	for _, phase := range events {
		phaseMap, ok := phase.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("failed to assert phase data")
		}

		// 访问 'units' 字段
		units, ok := phaseMap["units"].([]interface{})
		if !ok {
			continue // 如果没有 'units' 字段，则跳过
		}
		shortDescription := "null"
		// 遍历每个单位
		for _, unit := range units {
			unitMap, ok := unit.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("failed to assert unit data")
			}

			// 从单位中提取所需字段
			unitDescription := unitMap["description"].(string)
			shortDescription, ok = unitMap["shortDescription"].(string)

			// 访问 'schedule' 中的 'items'
			schedule, ok := unitMap["schedule"].(map[string]interface{})
			if !ok {
				continue // 如果没有 'schedule' 字段，则跳过
			}

			result, ok := schedule["result"].(map[string]interface{})
			if !ok {
				continue // 如果没有 'result' 字段，则跳过
			}

			extends, ok := result["extendedInfos"].([]interface{})
			if !ok {
				continue
			}
			special := ""
			for _, extend := range extends {
				ex, ok := extend.(map[string]interface{})
				if !ok {
					continue
				}
				code := ex["ei_code"].(string)
				if code == "RES_CODE" {
					special = ex["ei_value"].(string)
					break
				}
			}

			items, ok := result["items"].([]interface{})
			if !ok {
				continue // 如果没有 'items' 字段，则跳过
			}
			var competitors []*model.Competitor
			for _, item := range items {
				i, ok := item.(map[string]interface{})
				if !ok {
					continue
				}
				rating := i["resultData"].(string)
				winnerLoserTie, ok := i["resultWLT"].(string)
				if !ok {

				}
				participant := i["participant"].(map[string]interface{})
				name := participant["name"].(string)
				organisation, ok := participant["organisation"].(map[string]interface{})
				if !ok {
					continue
				}
				flag := organisation["code"].(string)
				competitors = append(competitors, &model.Competitor{
					Name:           name,
					Rating:         rating,
					WinnerLoserTie: winnerLoserTie,
					Flag:           flag,
				})
			}
			// 创建事件结构体并添加到列表
			parsedEvents = append(parsedEvents, &model.EventTable{
				Title:       shortDescription,
				Period:      unitDescription,
				Competitors: competitors,
				Special:     special,
			})
		}
	}

	return parsedEvents, nil
}

func ContestList(eventID string) ([]*model.ContestList, error) {
	url := "https://olympics.com/OG2024/data/GLO_EventGames~comp=OG2024~event=" + eventID + "~lang=CHI.json"
	data, err := fetchPageContent(url)
	if err != nil {
		return nil, err
	}
	return parseContest(data)
}

func parseContest(data string) ([]*model.ContestList, error) {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(data), &jsonMap)
	if err != nil {
		return nil, fmt.Errorf("failed to parse json data: %v", err)
	}

	// 获取事件数据
	events, ok := jsonMap["event"].(map[string]interface{})["phases"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to assert data")
	}

	var contestList []*model.ContestList

	// 遍历所有 phases
	for _, phase := range events {
		phaseMap, ok := phase.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("failed to assert phase data")
		}

		// 访问 'units' 字段
		units, ok := phaseMap["units"].([]interface{})
		if !ok {
			continue // 如果没有 'units' 字段，则跳过
		}

		// 创建一个 ContestList（如果该 phase 有 units）
		var contest *model.ContestList
		for _, unit := range units {
			unitMap, ok := unit.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("failed to assert unit data")
			}
			id, ok := unitMap["code"].(string)
			// 从单位中提取所需字段
			shortDescription, ok := unitMap["shortDescription"].(string)
			if !ok {
				shortDescription = "null"
			}

			// 访问 'schedule' 中的 'items'
			schedule, ok := unitMap["schedule"].(map[string]interface{})
			if !ok {
				continue // 如果没有 'schedule' 字段，则跳过
			}

			startDate, ok := schedule["startDate"].(string)
			//unitNum, ok := schedule["unitNum"].(string)

			result, ok := schedule["result"].(map[string]interface{})
			if !ok {
				continue // 如果没有 'result' 字段，则跳过
			}

			items, ok := result["items"].([]interface{})
			if !ok {
				continue // 如果没有 'items' 字段，则跳过
			}

			// 创建 competitors 列表
			var competitors []*model.Competitor
			for _, item := range items {
				i, ok := item.(map[string]interface{})
				if !ok {
					continue
				}
				rating := i["resultData"].(string)
				winnerLoserTie, ok := i["resultWLT"].(string)
				participant := i["participant"].(map[string]interface{})
				organisation, ok := participant["organisation"].(map[string]interface{})
				if !ok {
					continue
				}
				country := organisation["description"].(string)
				competitors = append(competitors, &model.Competitor{
					Name:           country,
					Rating:         rating,
					WinnerLoserTie: winnerLoserTie,
				})
			}

			// 创建 Contest
			contestItem := &model.Contest{
				ID:      id,
				Country: competitors,
			}

			// 如果 contest 为空，则初始化 ContestList
			if contest == nil {
				contest = &model.ContestList{
					Title:       shortDescription,
					Date:        startDate,
					Competitors: []*model.Contest{contestItem},
				}
			} else {
				// 否则将 Contest 添加到现有的 Competitors 列表中
				contest.Competitors = append(contest.Competitors, contestItem)
			}
		}

		// 将生成的 ContestList 添加到 contestList 中
		if contest != nil {
			contestList = append(contestList, contest)
		}
	}

	return contestList, nil
}
