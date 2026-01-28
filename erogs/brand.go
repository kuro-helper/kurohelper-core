package erogs

import (
	"encoding/json"
	"fmt"

	kurohelpercore "kurohelper-core"
)

type (
	Brand struct {
		ID            int    `json:"id"`
		BrandName     string `json:"brandname"`
		BrandFurigana string `json:"brandfurigana"`
		URL           string `json:"url"`
		Kind          string `json:"kind"`
		Lost          bool   `json:"lost"`
		DirectLink    bool   `json:"directlink"` // 網站可不可用
		Median        int    `json:"median"`     // 該品牌的遊戲評分中位數(一天更新一次)
		Twitter       string `json:"twitter"`
		Count2        int    `json:"count2"`
		CountAll      int    `json:"count_all"`
		Average2      int    `json:"average2"`
		Stdev         int    `json:"stdev"` // 標準偏差值(更新週期官方沒寫明確)
		GameList      []struct {
			ID       int    `json:"id"`
			GameName string `json:"gamename"`
			Furigana string `json:"furigana"`
			SellDay  string `json:"sellday"`
			Model    string `json:"model"`
			Median   int    `json:"median"` // 分數中位數(一天更新一次)
			Stdev    int    `json:"stdev"`  // 分數標準偏差值(一天更新一次)
			Count2   int    `json:"count2"` // 分數計算的樣本數
			VNDB     string `json:"vndb"`   // *string
		} `json:"gamelist"`
	}
)

func GetBrandByFuzzy(search string) (*Brand, error) {
	searchJP := kurohelpercore.ZhTwToJp(search)
	sql, err := buildBrandSQL(search, searchJP, false)
	if err != nil {
		return nil, err
	}

	jsonText, err := sendPostRequest(sql)
	if err != nil {
		return nil, err
	}

	var res Brand
	err = json.Unmarshal([]byte(jsonText), &res)
	if err != nil {
		fmt.Println(jsonText)
		return nil, err
	}

	return &res, nil
}

func GetBrandByID(id string) (*Brand, error) {
	sql, err := buildBrandSQL(id, "", true)
	if err != nil {
		return nil, err
	}

	jsonText, err := sendPostRequest(sql)
	if err != nil {
		return nil, err
	}

	var res Brand
	err = json.Unmarshal([]byte(jsonText), &res)
	if err != nil {
		fmt.Println(jsonText)
		return nil, err
	}

	return &res, nil
}

func buildBrandSQL(searchTW string, searchJP string, useID bool) (string, error) {
	buildString := ""
	if useID {
		buildString = fmt.Sprintf("WHERE id = '%s'", searchTW)
	} else {
		resultTW, err := buildSearchStringSQL(searchTW)
		if err != nil {
			return "", err
		}

		resultJP, err := buildSearchStringSQL(searchJP)
		if err != nil {
			return "", err
		}

		buildString = fmt.Sprintf("WHERE brandname ILIKE '%s' OR brandname ILIKE '%s'", resultTW, resultJP)
	}

	return fmt.Sprintf(`
WITH single_brand AS (
    SELECT
        id,
        brandname,
        brandfurigana,
        url,
        kind,
        lost,
        directlink,
        median,
        twitter,
        count2,
        count_all,
        average2,
        stdev
    FROM brandlist
    %s
    ORDER BY count2 DESC NULLS LAST, median DESC NULLS LAST
    LIMIT 1
)
SELECT row_to_json(r)
FROM (
    SELECT 
        A.id, 
        A.brandname, 
        A.brandfurigana, 
        A.url, 
        A.kind, 
        A.lost, 
        A.directlink, 
        A.median, 
        A.twitter, 
        A.count2, 
        A.count_all, 
        A.average2, 
        A.stdev,
        (
            SELECT json_agg(
                json_build_object(
                    'id', g.id,
                    'gamename', g.gamename,
                    'furigana', g.furigana,
                    'sellday', g.sellday,
                    'median', g.median,
                    'model', g.model,
                    'stdev', g.stdev,
                    'count2', g.count2,
                    'vndb', g.vndb
                ) ORDER BY g.sellday DESC
            )
            FROM gamelist g
            WHERE g.brandname = A.id
        ) AS gamelist
    FROM single_brand A
) r;
`, buildString), nil
}
