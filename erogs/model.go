package erogs

// 只抓一筆(LIMIT 1)
type FuzzySearchCreatorResponse struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	TwitterUsername string `json:"twitter_username"`
	Blog            string `json:"blog"`
	Pixiv           *int   `json:"pixiv"`
	Games           []Game `json:"games"` // 參與過的遊戲
}

type Game struct {
	Gamename string     `json:"gamename"`
	SellDay  string     `json:"sellday"`
	Median   int        `json:"median"`
	CountAll int        `json:"count2"`
	Shokushu []Shokushu `json:"shokushu"` // 有可能一個遊戲有多種身分
}

type Shokushu struct {
	Shubetu           int    `json:"shubetu"`
	ShubetuDetail     int    `json:"shubetu_detail"`
	ShubetuDetailName string `json:"shubetu_detail_name"` // *string
}

type FuzzySearchGameResponse struct {
	ID                               int              `json:"id"`
	BrandID                          int              `json:"brandid"`
	BrandName                        string           `json:"brandname"`
	Gamename                         string           `json:"gamename"`
	SellDay                          string           `json:"sellday"`
	Model                            string           `json:"model"`
	DMM                              string           `json:"dmm"` // dmm image
	Median                           string           `json:"median"`
	TokutenCount                     string           `json:"count2"`
	TotalPlayTimeMedian              string           `json:"total_play_time_median"`
	TimeBeforeUnderstandingFunMedian string           `json:"time_before_understanding_fun_median"`
	Okazu                            string           `json:"okazu"`
	Erogame                          string           `json:"erogame"`
	Genre                            string           `json:"genre"`
	BannerUrl                        string           `json:"banner_url"`
	SteamId                          string           `json:"steam"`
	VndbId                           string           `json:"vndb"`
	Shoukai                          string           `json:"shoukai"`
	Junni                            int              `json:"junni"`
	CreatorShubetu                   []Creatorshubetu `json:"shubetu_detail"`
}

type Creatorshubetu struct {
	ShubetuType       int    `json:"shubetu_type"`
	CreatorName       string `json:"creater_name"`
	ShubetuDetailType int    `json:"shubetu_detail_type"`
	ShubetuDetailName string `json:"shubetu_detail_name"`
}

type FuzzySearchListResponse struct {
	ID                               int    `json:"id"`
	Name                             string `json:"name"`
	DMM                              string `json:"dmm"` // dmm image
	Median                           string `json:"median"`
	TokutenCount                     string `json:"count2"`
	TotalPlayTimeMedian              string `json:"total_play_time_median"`
	TimeBeforeUnderstandingFunMedian string `json:"time_before_understanding_fun_median"`
	Category                         string `json:"category"`
	Model                            string `json:"model"`
}

type FuzzySearchCharacterResponse struct {
	ID            int    `json:"id"`
	CharacterName string `json:"name"`
	Sex           string `json:"sex"`
	BloodType     string `json:"bloodtype"`
	Birthday      string `json:"birthday"`
	GameName      string `json:"gamename"`
	URL           string `json:"url"`
	FormalExplain string `json:"formal_explanation"`
	Age           string `json:"age"`
	Bust          string `json:"bust"`
	Waist         string `json:"waist"`
	Hip           string `json:"hip"`
	Height        string `json:"height"`
	Weight        string `json:"weight"`
	Cup           string `json:"cup"`
	Role          int    `json:"role"`
	CreatorName   string `json:"creater_name"`
}
