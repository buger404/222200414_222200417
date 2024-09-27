package consts

const (
	Url         = "https://olympics.com/zh/paris-2024/medals"
	FootballUrl = "https://olympics.com/OG2024/data/GLO_EventGames~comp=OG2024~event=FBLMTEAM11------------~lang=CHI.json"
	ErrorBind   = 10001
	ERROR       = 10000
)

var (
	CountryMap = map[string]string{
		"阿尔巴尼亚":          "ALB",
		"阿尔及利亚":          "ALG",
		"阿富汗":            "AFG",
		"阿根廷":            "ARG",
		"阿联酋":            "UAE",
		"阿鲁巴":            "ARU",
		"阿曼":             "OMA",
		"阿塞拜疆":           "AZE",
		"埃及":             "EGY",
		"埃塞俄比亚":          "ETH",
		"爱尔兰":            "IRL",
		"爱沙尼亚":           "EST",
		"安道尔":            "AND",
		"安哥拉":            "ANG",
		"安提瓜和巴布达":        "ANT",
		"奥地利":            "AUT",
		"奥林匹克难民代表团":      "EOR",
		"澳大利亚":           "AUS",
		"巴巴多斯":           "BAR",
		"巴布亚新几内亚":        "PNG",
		"巴哈马":            "BAH",
		"巴基斯坦":           "PAK",
		"巴拉圭":            "PAR",
		"巴勒斯坦":           "PLE",
		"巴林":             "BRN",
		"巴拿马":            "PAN",
		"巴西":             "BRA",
		"百慕大":            "BER",
		"保加利亚":           "BUL",
		"北马其顿":           "MKD",
		"贝宁":             "BEN",
		"比利时":            "BEL",
		"冰岛":             "ISL",
		"波多黎各":           "PUR",
		"波黑":             "BIH",
		"波兰":             "POL",
		"玻利维亚":           "BOL",
		"伯利兹":            "BIZ",
		"博茨瓦纳":           "BOT",
		"不丹":             "BHU",
		"布基纳法索":          "BUR",
		"布隆迪":            "BDI",
		"朝鲜":             "PRK",
		"赤道几内亚":          "GEQ",
		"丹麦":             "DEN",
		"德国":             "GER",
		"东帝汶":            "TLS",
		"多哥":             "TOG",
		"多米尼加":           "DOM",
		"多米尼克":           "DMA",
		"厄瓜多尔":           "ECU",
		"厄立特里亚":          "ERI",
		"法国":             "FRA",
		"菲律宾":            "PHI",
		"斐济":             "FIJ",
		"芬兰":             "FIN",
		"佛得角":            "CPV",
		"冈比亚":            "GAM",
		"刚果（布）":          "CGO",
		"刚果（金）":          "COD",
		"哥伦比亚":           "COL",
		"哥斯达黎加":          "CRC",
		"格林纳达":           "GRN",
		"格鲁吉亚":           "GEO",
		"古巴":             "CUB",
		"关岛":             "GUM",
		"圭亚那":            "GUY",
		"哈萨克斯坦":          "KAZ",
		"海地":             "HAI",
		"韩国":             "KOR",
		"荷兰":             "NED",
		"黑山":             "MNE",
		"洪都拉斯":           "HON",
		"基里巴斯":           "KIR",
		"吉布提":            "DJI",
		"吉尔吉斯斯坦":         "KGZ",
		"几内亚":            "GUI",
		"几内亚比绍":          "GBS",
		"加拿大":            "CAN",
		"加纳":             "GHA",
		"加蓬":             "GAB",
		"柬埔寨":            "CAM",
		"捷克":             "CZE",
		"津巴布韦":           "ZIM",
		"喀麦隆":            "CMR",
		"卡塔尔":            "QAT",
		"开曼群岛":           "CAY",
		"科摩罗":            "COM",
		"科索沃":            "KOS",
		"科特迪瓦":           "CIV",
		"科威特":            "KUW",
		"克罗地亚":           "CRO",
		"肯尼亚":            "KEN",
		"库克群岛":           "COK",
		"拉脱维亚":           "LAT",
		"莱索托":            "LES",
		"老挝":             "LAO",
		"黎巴嫩":            "LBN",
		"立陶宛":            "LTU",
		"利比里亚":           "LBR",
		"利比亚":            "LBA",
		"列支敦士登":          "LIE",
		"卢森堡":            "LUX",
		"卢旺达":            "RWA",
		"罗马尼亚":           "ROU",
		"马达加斯加":          "MAD",
		"马尔代夫":           "MDV",
		"马耳他":            "MLT",
		"马拉维":            "MAW",
		"马来西亚":           "MAS",
		"马里":             "MLI",
		"马绍尔群岛":          "MHL",
		"毛里求斯":           "MRI",
		"毛里塔尼亚":          "MTN",
		"美国":             "USA",
		"美属萨摩亚":          "ASA",
		"美属维尔京群岛":        "ISV",
		"蒙古":             "MGL",
		"孟加拉国":           "BAN",
		"秘鲁":             "PER",
		"密克罗尼西亚":         "FSM",
		"缅甸":             "MYA",
		"摩尔多瓦":           "MDA",
		"摩洛哥":            "MAR",
		"摩纳哥":            "MON",
		"莫桑比克":           "MOZ",
		"墨西哥":            "MEX",
		"纳米比亚":           "NAM",
		"南非":             "RSA",
		"南苏丹":            "SSD",
		"瑙鲁":             "NRU",
		"尼泊尔":            "NEP",
		"尼加拉瓜":           "NCA",
		"尼日尔":            "NIG",
		"尼日利亚":           "NGR",
		"挪威":             "NOR",
		"帕劳":             "PLW",
		"葡萄牙":            "POR",
		"日本":             "JPN",
		"瑞典":             "SWE",
		"瑞士":             "SUI",
		"萨尔瓦多":           "ESA",
		"萨摩亚":            "SAM",
		"塞尔维亚":           "SRB",
		"塞拉利昂":           "SLE",
		"塞内加尔":           "SEN",
		"塞浦路斯":           "CYP",
		"塞舌尔":            "SEY",
		"沙特阿拉伯":          "KSA",
		"圣多美和普林西比":       "STP",
		"圣基茨和尼维斯":        "SKN",
		"圣卢西亚":           "LCA",
		"圣马力诺":           "SMR",
		"圣文森特和格林纳丁斯":     "VIN",
		"斯里兰卡":           "SRI",
		"斯洛伐克":           "SVK",
		"斯洛文尼亚":          "SLO",
		"斯威士兰":           "SWZ",
		"苏丹":             "SUD",
		"苏里南":            "SUR",
		"所罗门群岛":          "SOL",
		"索马里":            "SOM",
		"塔吉克斯坦":          "TJK",
		"泰国":             "THA",
		"坦桑尼亚":           "TAN",
		"汤加":             "TGA",
		"Chinese Taipei": "TPE",
		"特立尼达和多巴哥":       "TTO",
		"突尼斯":            "TUN",
		"图瓦卢":            "TUV",
		"土耳其":            "TUR",
		"土库曼斯坦":          "TKM",
		"瓦努阿图":           "VAN",
		"危地马拉":           "GUA",
		"委内瑞拉":           "VEN",
		"文莱":             "BRU",
		"乌干达":            "UGA",
		"乌克兰":            "UKR",
		"乌拉圭":            "URU",
		"乌兹别克斯坦":         "UZB",
		"西班牙":            "ESP",
		"希腊":             "GRE",
		"新加坡":            "SGP",
		"新西兰":            "NZL",
		"匈牙利":            "HUN",
		"叙利亚":            "SYR",
		"牙买加":            "JAM",
		"亚美尼亚":           "ARM",
		"也门":             "YEM",
		"伊拉克":            "IRQ",
		"伊朗":             "IRI",
		"以色列":            "ISR",
		"意大利":            "ITA",
		"印度":             "IND",
		"印度尼西亚":          "INA",
		"英国":             "GBR",
		"英属维尔京群岛":        "IVB",
		"约旦":             "JOR",
		"越南":             "VIE",
		"赞比亚":            "ZAM",
		"乍得":             "CHA",
		"泽西岛":            "JER",
		"中国":             "CHN",
		"中国香港":           "HKG",
		"智利":             "CHI",
	}
)
