package ymgal

import (
	"encoding/json"
	"fmt"
	"strings"
)

type (
	SearchGameResp struct {
		Result   []Result `json:"result"`
		Total    int      `json:"total"`
		HasNext  bool     `json:"hasNext"`
		PageNum  int      `json:"pageNum"`
		PageSize int      `json:"pageSize"`
	}

	Result struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		ChineseName    string `json:"chineseName"`
		State          string `json:"state"`
		Weights        int    `json:"weights"`
		MainImg        string `json:"mainImg"`
		PublishVersion int    `json:"publishVersion"`
		PublishTime    string `json:"publishTime"`
		Publisher      int    `json:"publisher"`
		Score          string `json:"score"`
		OrgID          int    `json:"orgId"`
		OrgName        string `json:"orgName"`
		ReleaseDate    string `json:"releaseDate"`
		HaveChinese    bool   `json:"haveChinese"`
	}
)

func SearchGame(keyword string) (*SearchGameResp, error) {
	r, err := sendWithRetry(fmt.Sprintf("%s/open/archive/search-game?mode=list&pageNum=1&keyword=%s&pageSize=1", cfg.Endpoint, strings.TrimSpace(keyword)))
	if err != nil {
		return nil, err
	}

	var result basicResp[SearchGameResp]

	// 如果查不到資料就會直接回傳HTML，無法用status code判斷
	err = json.Unmarshal(r, &result)
	if err != nil {
		return nil, err
	}

	if !result.Success {
		return nil, ErrAPIFailed{Code: result.Code}
	}

	return &result.Data, nil
}
