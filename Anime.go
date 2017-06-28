package shoboi

import (
	"regexp"
	"strings"
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
