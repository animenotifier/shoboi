package shoboi

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"time"
)

// AiringDate ...
type AiringDate struct {
	Start     string `json:"startDate"`
	End       string `json:"endDate"`
	ProgramID string `json:"programId"`
	ChannelID string `json:"channelId"`
}

// GetAiringDate ...
func GetAiringDate(animeID string, episodeNumber int) *AiringDate {
	programList := &ProgramList{}
	resp, err := get("http://cal.syoboi.jp/json.php?Req=ProgramByCount&TID=" + animeID + "&Count=" + strconv.Itoa(episodeNumber))

	if err != nil {
		return nil
	}

	if resp.StatusCode() != http.StatusOK {
		return nil
	}

	err = json.Unmarshal(resp.BodyBytes(), programList)

	if err != nil {
		return nil
	}

	if programList.Programs == nil {
		return nil
	}

	// Find earliest airing date
	var endTimeErr error
	earliestPID := ""
	earliestChannelID := ""
	earliestStartTime := int64(math.MaxInt64)
	earliestEndTime := int64(math.MaxInt64)

	for pid, airingDateInfo := range programList.Programs {
		startTime, err := strconv.ParseInt(airingDateInfo.StartTime, 10, 64)

		if err == nil && startTime != 0 && startTime < earliestStartTime {
			earliestPID = pid
			earliestChannelID = airingDateInfo.ChannelID
			earliestStartTime = startTime
			earliestEndTime, endTimeErr = strconv.ParseInt(airingDateInfo.EndTime, 10, 64)
		}
	}

	airingDate := &AiringDate{
		ProgramID: earliestPID,
		ChannelID: earliestChannelID,
		Start:     time.Unix(earliestStartTime, 0).In(time.UTC).Format(time.RFC3339),
	}

	if endTimeErr != nil || earliestEndTime == 0 {
		airingDate.End = ""
	} else {
		airingDate.End = time.Unix(earliestEndTime, 0).In(time.UTC).Format(time.RFC3339)
	}

	return airingDate
}
