package model

// 定义国家和奖牌信息的结构体
type MedalWinner struct {
	CompetitorType        string `json:"competitorType"`
	CompetitorDisplayName string `json:"competitorDisplayName"`
	Date                  string `json:"date"`
	MedalType             string `json:"medalType"`
}

type Discipline struct {
	Name         string        `json:"name"`
	Total        int           `json:"total"`
	Gold         int           `json:"gold"`
	Silver       int           `json:"silver"`
	Bronze       int           `json:"bronze"`
	MedalWinners []MedalWinner `json:"medalWinners"`
}

type MedalEntry struct {
	Description string       `json:"description"`
	Disciplines []Discipline `json:"disciplines"`
}

type NocListEntry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MedalsData struct {
	MedalsTable []MedalEntry `json:"medalsTable"`
}

type PageProps struct {
	NocList       []NocListEntry `json:"nocList"`
	InitialMedals MedalsData     `json:"initialMedals"`
}

type OlympicsData struct {
	Props PageProps `json:"pageProps"`
}

type Competitor struct {
	Name           string `json:"name"`
	WinnerLoserTie string `json:"winnerLoserTie"`
	Rating         string `json:"mark"`
	Flag           string `json:"flag"`
}

type Event struct {
	StartDate      string        `json:"startDate"`
	PhaseName      string        `json:"phaseName"`
	ID             string        `json:"id"`
	DisciplineName string        `json:"disciplineName"`
	Competitors    []*Competitor `json:"competitors"`
}

type EventList struct {
	Name string        `json:"name"`
	List EventTypeList `json:"list"`
}

type EventTypeList []map[string]string

type EventTable struct {
	Title       string        `json:"title"`
	Period      string        `json:"period"`
	Special     string        `json:"special"`
	Competitors []*Competitor `json:"competitors"`
}

type ContestList struct {
	Title       string     `json:"title"`
	Date        string     `json:"date"`
	Competitors []*Contest `json:"competitors"`
}

type Contest struct {
	ID      string        `json:"ID"`
	Country []*Competitor `json:"country"`
}
