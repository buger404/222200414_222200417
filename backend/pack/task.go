package pack

import (
	spiderModel "backend/biz/dal/model"
	"backend/biz/model/model"
	"sort"
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
