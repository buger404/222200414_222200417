package pack

import (
	spiderModel "backend/biz/dal/model"
	"backend/biz/model/model"
	consts2 "backend/consts"
	"fmt"
	"sort"
	"time"
)

// WrapOlympicsData 将 OlympicsData 转换为 AllMedalsResp
func WrapOlympicsData(olympicsData *spiderModel.OlympicsData) *model.MedalRank {
	return &model.MedalRank{
		// 根据 Medals 结构体定义填充数据
		Details: convertToMedals(olympicsData),
	}
}

// convertToMedals 将 MedalsData 转换为 model.Medals 并进行排序
func convertToMedals(medalsData *spiderModel.OlympicsData) []*model.Medal {
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
			countryMedals.Sliver += int64(discipline.Silver) // 拼写修正：Sliver -> Silver
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

	// 根据 Total, Gold, Silver, Bronze 进行排序
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].List.Total != entries[j].List.Total {
			return entries[i].List.Total > entries[j].List.Total
		}
		if entries[i].List.Gold != entries[j].List.Gold {
			return entries[i].List.Gold > entries[j].List.Gold
		}
		if entries[i].List.Sliver != entries[j].List.Sliver {
			return entries[i].List.Sliver > entries[j].List.Sliver
		}
		return entries[i].List.Bronze > entries[j].List.Bronze
	})

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
	var modelEvents []*model.Event

	for _, event := range events {
		var countries []*model.Country
		if event.Competitors == nil {
			continue
		}
		// 转换 Competitor 为 Country
		for _, competitor := range event.Competitors {
			countries = append(countries, &model.Country{
				Name:   competitor.Name,
				Rating: competitor.Rating,
				Flag:   consts2.CountryMap[competitor.Name], // 如果有旗帜信息，可以在此处填入
			})
		}

		var res = 16
		fmt.Println(event.Competitors[0].WinnerLoserTie)
		if event.Competitors[0].WinnerLoserTie == "W" {
			res = 2
		} else {
			res = 1
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
