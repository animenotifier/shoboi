package shoboi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/fatih/color"
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
)

func TestAnime(t *testing.T) {
	tid := "4528"
	resp, body, errs := gorequest.New().Get("http://cal.syoboi.jp/json.php?Req=TitleFull&TID=" + tid).EndBytes()

	if len(errs) > 0 {
		panic(errs[0])
	}

	assert.Equal(t, resp.StatusCode, http.StatusOK)

	titleFull := &TitleFull{}
	json.Unmarshal(body, titleFull)

	assert.NotNil(t, titleFull)
	assert.NotNil(t, titleFull.Titles)
	assert.NotNil(t, titleFull.Titles[tid])

	anime := titleFull.Titles[tid]
	for i, episode := range anime.EpisodeNames() {
		fmt.Printf("Episode %d: %s\n", i+1, color.GreenString(episode))
	}
}
