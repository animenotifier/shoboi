package shoboi

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// SearchResult ...
type SearchResult struct {
	Titles map[string]*AnimeSearchResult `json:"Titles"`
}

// AnimeSearchResult ...
type AnimeSearchResult struct {
	TID           string `json:"TID"`
	TitleJapanese string `json:"Title"`
	ShortTitle    string `json:"ShortTitle"`
	TitleHiragana string `json:"TitleYomi"`
	TitleEnglish  string `json:"TitleEN"`
	Cat           string `json:"Cat"`
	FirstChannel  string `json:"FirstCh"`
	FirstYear     string `json:"FirstYear"`
	FirstMonth    string `json:"FirstMonth"`
	FirstEndYear  string `json:"FirstEndYear"`
	FirstEndMonth string `json:"FirstEndMonth"`
	TitleFlag     string `json:"TitleFlag"`
	Comment       string `json:"Comment"`
	Search        int    `json:"Search"`
	Programs      []struct {
		PID         string      `json:"PID"`
		TID         string      `json:"TID"`
		StTime      string      `json:"StTime"`
		EdTime      string      `json:"EdTime"`
		ChID        string      `json:"ChID"`
		StOffset    string      `json:"StOffset"`
		Count       interface{} `json:"Count"`
		ProgComment string      `json:"ProgComment"`
		SubTitle    string      `json:"SubTitle"`
		ChName      string      `json:"ChName"`
	} `json:"Programs"`
}

// SearchAnime ...
func SearchAnime(title string) (anime *AnimeSearchResult, err error) {
	title = strings.ToLower(title)
	searchResult := &SearchResult{}
	resp, err := get("http://cal.syoboi.jp/json?Req=TitleSearch&Search=" + title + "&Limit=15")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New("Invalid status code: " + strconv.Itoa(resp.StatusCode()))
	}

	err = json.Unmarshal(resp.BodyBytes(), searchResult)

	if err != nil {
		return nil, err
	}

	if searchResult == nil || searchResult.Titles == nil {
		return nil, nil
	}

	for _, anime := range searchResult.Titles {
		if strings.ToLower(anime.TitleJapanese) == title || strings.ToLower(anime.TitleEnglish) == title {
			return anime, nil
		}
	}

	return nil, nil
}
