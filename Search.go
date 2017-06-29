package shoboi

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/parnurzeal/gorequest"
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
	FirstCh       string `json:"FirstCh"`
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
func SearchAnime(title string) (tid string, err error) {
	title = strings.ToLower(title)
	searchResult := &SearchResult{}
	resp, _, errs := gorequest.New().Get("http://cal.syoboi.jp/json?Req=TitleSearch&Search=" + title + "&Limit=15").EndStruct(searchResult)

	if len(errs) > 0 {
		return "", errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Invalid status code: " + strconv.Itoa(resp.StatusCode))
	}

	if searchResult == nil || searchResult.Titles == nil {
		return "", errors.New("Invalid data: Titles is nil")
	}

	for _, anime := range searchResult.Titles {
		if strings.ToLower(anime.TitleJapanese) == title || strings.ToLower(anime.TitleEnglish) == title {
			return anime.TID, nil
		}
	}

	return "", nil
}
