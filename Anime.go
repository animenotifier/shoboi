package shoboi

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/parnurzeal/gorequest"
)

var episodeNameRegex = regexp.MustCompile(`\*(\d+)\*(.*)`)

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
	TitleJapanese string `xml:"Title" json:"Title"`
	TitleEnglish  string `xml:"TitleEN" json:"TitleEN"`
	TitleHiragana string `xml:"TitleYomi" json:"TitleYomi"`
	TitleFlag     string `xml:"TitleFlag" json:"TitleFlag"`
	UpdatedAt     string `xml:"LastUpdate" json:"LastUpdate"`
	UserPoint     string `xml:"UserPoint" json:"UserPoint"`
	UserPointRank string `xml:"UserPointRank" json:"UserPointRank"`
}

// GetAnime ...
func GetAnime(tid string) (*Anime, error) {
	const maxTries = 5
	tryCount := 0
	titleFull := &TitleFull{}
	tryDelay := 10 * time.Second

	var resp gorequest.Response
	var errs []error

	for {
		resp, _, errs = gorequest.New().Get("http://cal.syoboi.jp/json.php?Req=TitleFull&TID=" + tid).EndStruct(titleFull)

		if resp.StatusCode == http.StatusOK {
			break
		}

		tryCount++

		color.Red("Shoboi status code %d (#%d)", resp.StatusCode, tryCount)

		if tryCount > maxTries {
			break
		}

		time.Sleep(tryDelay)

		// Exponential falloff
		tryDelay *= 2
	}

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
	for _, episode := range episodeNames {
		episode = strings.TrimSpace(episode)
		matches := episodeNameRegex.FindStringSubmatch(episode)

		if len(matches) > 2 {
			episodeNumber, _ := strconv.Atoi(matches[1])

			episodes = append(episodes, &Episode{
				Number:        episodeNumber,
				TitleJapanese: matches[2],
				anime:         anime,
			})
		}
	}

	return episodes
}
