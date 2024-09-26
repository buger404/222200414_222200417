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
)

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
