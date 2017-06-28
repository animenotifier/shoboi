package shoboi

import (
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/parnurzeal/gorequest"
)

// Episode ...
type Episode struct {
	Number        int
	TitleJapanese string
	anime         *Anime
}

// AiringDate ...
func (episode *Episode) AiringDate() *AiringDate {
	programList := &ProgramList{}
	resp, _, errs := gorequest.New().Get("http://cal.syoboi.jp/json.php?Req=ProgramByCount&TID=" + episode.anime.ID + "&Count=" + strconv.Itoa(episode.Number)).EndStruct(programList)

	if len(errs) > 0 {
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	// Find earliest airing date
	var endTimeErr error
	earliestStartTime := int64(math.MaxInt64)
	earliestEndTime := int64(math.MaxInt64)

	for _, airingDateInfo := range programList.Programs {
		startTime, err := strconv.ParseInt(airingDateInfo.StartTime, 10, 64)

		if err == nil && startTime != 0 && startTime < earliestStartTime {
			earliestStartTime = startTime
			earliestEndTime, endTimeErr = strconv.ParseInt(airingDateInfo.EndTime, 10, 64)
		}
	}

	airingDate := &AiringDate{
		Start: time.Unix(earliestStartTime, 0).In(time.UTC).Format(time.RFC3339),
	}

	if endTimeErr != nil || earliestEndTime == 0 {
		airingDate.End = ""
	} else {
		airingDate.End = time.Unix(earliestEndTime, 0).In(time.UTC).Format(time.RFC3339)
	}

	return airingDate
}
