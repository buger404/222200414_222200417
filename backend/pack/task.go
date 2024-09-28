package pack

import (
	spiderModel "backend/biz/dal/model"
	"backend/biz/model/model"
	consts2 "backend/consts"
	"sort"
	"time"
)

// WrapOlympicsData 将 OlympicsData 转换为 AllMedalsResp
func WrapOlympicsData(olympicsData *spiderModel.OlympicsData, medalSorts int) *model.MedalRank {
	return &model.MedalRank{
		// 根据 Medals 结构体定义填充数据
		Details: convertToMedals(olympicsData, medalSorts),
	}
}

// convertToMedals 将 MedalsData 转换为 model.Medals 并进行排序
func convertToMedals(medalsData *spiderModel.OlympicsData, medalSorts int) []*model.Medal {
	// 创建存储奖牌条目的切片
	var entries []*model.Medal

	// 使用 map 来存储每个国家的奖牌信息
	countryMedalsMap := make(map[string]*model.Medals)

	for _, medalEntry := range medalsData.Props.InitialMedals.MedalsTable {
		countryName := medalEntry.Description // 获取国家名

		// 初始化或获取国家对应的 medal 信息
		if _, exists := countryMedalsMap[countryName]; !exists {
			countryMedalsMap[countryName] = &model.Medals{}
		}
		countryMedals := countryMedalsMap[countryName]

		// 累加该国家在所有 discipline 下的奖牌信息
		for _, discipline := range medalEntry.Disciplines {
			countryMedals.Gold += int64(discipline.Gold)
			countryMedals.Sliver += int64(discipline.Silver)
			countryMedals.Bronze += int64(discipline.Bronze)
			countryMedals.Total += int64(discipline.Total)
		}
	}

	// 将 map 中的数据转换为 Medal 结构体，并添加到 entries 切片中
	for country, medals := range countryMedalsMap {
		entry := &model.Medal{
			CountryName: country,
			Flag:        consts2.CountryMap[country],
			List:        medals,
		}
		entries = append(entries, entry)
	}

	// 根据 medalSorts 进行排序
	switch medalSorts {
	case 0:
		// 按金牌、银牌、铜牌顺序排序
		sort.Slice(entries, func(i, j int) bool {
			if entries[i].List.Gold != entries[j].List.Gold {
				return entries[i].List.Gold > entries[j].List.Gold
			}
			if entries[i].List.Sliver != entries[j].List.Sliver {
				return entries[i].List.Sliver > entries[j].List.Sliver
			}
			return entries[i].List.Bronze > entries[j].List.Bronze
		})

	case 1:
		// 按总奖牌数排序
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].List.Total > entries[j].List.Total
		})

	case 2:
		// 按国家首字母排序
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Flag < entries[j].Flag
		})
	}

	return entries
}

func parseTime(startDate string) string {
	// 定义 ISO 8601 时间格式
	layout := "2006-01-02T15:04:05-07:00"

	// 解析时间
	t, _ := time.Parse(layout, startDate)

	// 提取时和分，并格式化为字符串
	return t.Format("15:04")
}

func ConvertEvent(events []*spiderModel.Event) *model.DailyEvent {
	ReverseCountryMap := make(map[string]string)
	for k, v := range consts2.CountryMap {
		ReverseCountryMap[v] = k
	}

	var modelEvents []*model.Event

	for _, event := range events {
		var countries []*model.Country
		if event.Competitors == nil {
			continue
		}
		// 转换 Competitor 为 Country
		for _, competitor := range event.Competitors {
			name, ok := ReverseCountryMap[competitor.Name]
			if !ok {
				name = competitor.Name
			}
			countries = append(countries, &model.Country{
				Name:   name,
				Rating: competitor.Rating,
				Flag:   competitor.Flag, // 如果有旗帜信息，可以在此处填入
			})
		}

		var res = 16
		if event.Competitors[0].WinnerLoserTie == "W" {
			res = 1
		} else {
			res = 2
		}

		// 创建并添加转换后的 Event 到 modelEvents 切片
		modelEvents = append(modelEvents, &model.Event{
			Id:        event.ID,
			Time:      parseTime(event.StartDate),
			Event:     event.PhaseName,
			Group:     event.DisciplineName,
			Countries: countries,
			Winner:    int64(res),
		})
	}

	return &model.DailyEvent{
		Events: modelEvents,
	}
}

func WrapEventList(eventTypeData []*spiderModel.EventList) *model.EventTypeLists {
	// 创建最终返回的 EventTypeLists 对象
	eventTypeLists := &model.EventTypeLists{}
	// 使用 map 来跟踪已经处理的事件名称
	mergedEventTypeLists := make(map[string]*model.EventTypeList)

	// 遍历 eventTypeData 并填充 EventTypeLists
	for _, eventTypeMap := range eventTypeData {
		// 检查事件名称是否已经存在于 mergedEventTypeLists 中
		eventTypeList, exists := mergedEventTypeLists[eventTypeMap.Name]
		if !exists {
			// 如果不存在，则创建新的 EventTypeList 对象
			eventTypeList = &model.EventTypeList{
				Name:  eventTypeMap.Name,    // 从 eventTypeMap 中获取事件名称
				Types: []*model.EventType{}, // 初始化 Types 列表
			}
			// 将新创建的 EventTypeList 添加到 mergedEventTypeLists 中
			mergedEventTypeLists[eventTypeMap.Name] = eventTypeList
		}

		// 遍历每个 EventTypeList 中的 Types 数据并填充
		for _, event := range eventTypeMap.List {
			id, ok1 := event["id"]
			name, ok2 := event["name"]
			if !ok1 || !ok2 {
				continue // 如果数据缺失，则跳过
			}

			// 创建一个 EventType 对象并填充数据
			eventType := &model.EventType{
				Id:   id,
				Name: name,
			}

			// 将 EventType 添加到 EventTypeList 的 Types 字段
			eventTypeList.Types = append(eventTypeList.Types, eventType)
		}
	}

	// 使用 map 去除 eventTypeLists.List 中重复的 ID
	idSet := make(map[string]struct{})

	// 将合并好的 EventTypeList 添加到最终的 EventTypeLists
	for _, eventTypeList := range mergedEventTypeLists {
		uniqueTypes := make([]*model.EventType, 0)

		// 过滤掉重复的 ID
		for _, eventType := range eventTypeList.Types {
			if _, exists := idSet[eventType.Id]; !exists {
				uniqueTypes = append(uniqueTypes, eventType)
				idSet[eventType.Id] = struct{}{}
			}
		}

		// 将唯一的类型添加到最终的 EventTypeLists
		eventTypeList.Types = uniqueTypes
		eventTypeLists.List = append(eventTypeLists.List, eventTypeList)
	}

	return eventTypeLists
}

func ConvertEventTable(eventTable1 []*spiderModel.EventTable) *model.EventTable {
	var tables []*model.Table

	// 遍历所有 EventTable
	for _, et := range eventTable1 {
		// 创建新 Table
		if et.Title == "A组" || et.Title == "B组" || et.Title == "C组" || et.Title == "D组" {
			continue
		}
		table := &model.Table{
			Title:   et.Title,
			Period:  et.Period,
			Special: et.Special,
		}

		table.Winner = 2
		// 转换 Competitors 为 Countries
		var countries []*model.Country
		for i, competitor := range et.Competitors {
			country := &model.Country{
				Name:   competitor.Name,
				Rating: competitor.Rating,
				Flag:   consts2.CountryMap[competitor.Name], // 从 CountryMap 获取国家标志
			}
			countries = append(countries, country)

			// 假设第一个 Competitor 是 Winner
			if i == 0 && competitor.WinnerLoserTie == "W" {
				table.Winner = int64(i + 1) // 设置 Winner 为第一个 Competitor 的索引（+1 因为 Winner 索引从1开始）
			}
		}

		// 将 countries 添加到 Table
		table.Countries = countries

		// 将 Table 添加到 Tables 列表中
		tables = append(tables, table)
	}

	// 排序 Tables，按指定顺序
	sort.Slice(tables, func(i, j int) bool {
		return getOrder(tables[i].Title) < getOrder(tables[j].Title)
	})

	// 创建最终的 EventTable2
	eventTable2 := &model.EventTable{
		Tables: tables,
	}

	return eventTable2
}

// getOrder 返回标题的排序值
func getOrder(title string) int {
	switch title {
	case "1/4决赛":
		return 3
	case "半决赛":
		return 6
	case "铜牌赛":
		return 6
	case "金牌赛":
		return 7
	default:
		return 8 // 未知标题，放在最后
	}
}

func ConvertContestListToEventList(contestList []*spiderModel.ContestList) *model.EventLists {
	// 创建一个 EventList 切片来存储多个 EventList
	var eventLists []*model.EventList

	for _, contest := range contestList {
		// 创建新的 Contest 列表
		var contests []*model.Contest

		for _, c := range contest.Competitors {
			// 转换 Competitors 为 Countries
			var countries []*model.Country
			code := 2 // 默认 code 为 2
			for i, competitor := range c.Country {
				if i == 0 && competitor.WinnerLoserTie == "w" {
					code = 1 // 如果第一个 Competitor 是赢家，设置 code 为 1
				}
				country := &model.Country{
					Name:   competitor.Name,
					Rating: competitor.Rating,
					Flag:   consts2.CountryMap[competitor.Name], // 从 CountryMap 获取国旗
				}
				countries = append(countries, country)
			}

			// 创建新的 Contest 并添加到 contests 列表中
			newContest := &model.Contest{
				ContestId: c.ID,         // Contest ID
				Date:      contest.Date, // Contest 日期
				Countries: countries,    // 转换后的国家列表
				Status:    int64(code),  // 使用 code 作为状态
			}
			contests = append(contests, newContest)
		}

		// 创建新的 EventList 实例并添加到 eventLists 列表中
		eventList := &model.EventList{
			GroupId:  contest.Title, // 使用 ContestList 的标题作为 EventList 的 ID
			Contests: contests,      // 将 contests 列表添加到 EventList
		}

		eventLists = append(eventLists, eventList)
	}

	// 返回包含多个 EventList 的切片
	return &model.EventLists{
		Event: eventLists,
	}
}
