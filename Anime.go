package shoboi

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/parnurzeal/gorequest"
)

var episodeNameRegex = regexp.MustCompile(`\*\d{1,3}\*(.*)`)

// Anime ...
type Anime struct {
	ID            string `xml:"TID" json:"TID"`
	CategoryID    string `xml:"Cat" json:"Cat"`
	Comment       string `xml:"Comment" json:"Comment"`
	FirstChannel  string `xml:"FirstCh" json:"FirstCh"`
	FirstEndMonth string `xml:"FirstEndMonth" json:"FirstEndMonth"`
	FirstEndYear  string `xml:"FirstEndYear" json:"FirstEndYear"`
	FirstMonth    string `xml:"FirstMonth" json:"FirstMonth"`
	FirstYear     string `xml:"FirstYear" json:"FirstYear"`
	Keywords      string `xml:"Keywords" json:"Keywords"`
	ShortTitle    string `xml:"ShortTitle" json:"ShortTitle"`
	SubTitles     string `xml:"SubTitles" json:"SubTitles"`
	Title         string `xml:"Title" json:"Title"`
	TitleEnglish  string `xml:"TitleEN" json:"TitleEN"`
	TitleFlag     string `xml:"TitleFlag" json:"TitleFlag"`
	TitleYomi     string `xml:"TitleYomi" json:"TitleYomi"`
	UpdatedAt     string `xml:"LastUpdate" json:"LastUpdate"`
	UserPoint     string `xml:"UserPoint" json:"UserPoint"`
	UserPointRank string `xml:"UserPointRank" json:"UserPointRank"`
}

// GetAnime ...
func GetAnime(tid string) (*Anime, error) {
	titleFull := &TitleFull{}
	resp, _, errs := gorequest.New().Get("http://cal.syoboi.jp/json.php?Req=TitleFull&TID=" + tid).EndStruct(titleFull)

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Invalid status code: " + strconv.Itoa(resp.StatusCode))
	}

	if titleFull == nil || titleFull.Titles == nil || titleFull.Titles[tid] == nil {
		return nil, errors.New("Invalid data: Titles is nil")
	}

	return titleFull.Titles[tid], nil
}

// Episodes ...
func (anime *Anime) Episodes() []*Episode {
	if anime.SubTitles == "" {
		return nil
	}

	episodes := []*Episode{}

	episodeNames := strings.Split(anime.SubTitles, "\r\n")
	for i, episode := range episodeNames {
		episode = strings.TrimSpace(episode)
		matches := episodeNameRegex.FindStringSubmatch(episode)

		if len(matches) > 1 {
			episodes = append(episodes, &Episode{
				Number:        i + 1,
				TitleJapanese: matches[1],
				anime:         anime,
			})
		}
	}

	return episodes
}
