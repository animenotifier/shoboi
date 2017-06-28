package shoboi

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/fatih/color"
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
)

func TestEpisodes(t *testing.T) {
	tid := "4528"

	titleFull := &TitleFull{}
	resp, _, errs := gorequest.New().Get("http://cal.syoboi.jp/json.php?Req=TitleFull&TID=" + tid).EndStruct(titleFull)

	assert.Empty(t, errs)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	assert.NotNil(t, titleFull)
	assert.NotNil(t, titleFull.Titles)
	assert.NotNil(t, titleFull.Titles[tid])

	anime := titleFull.Titles[tid]
	for _, episode := range anime.Episodes() {
		fmt.Printf("Episode %d: %s\n", episode.Number, color.GreenString(episode.TitleJapanese))
		fmt.Println(episode.AiringDate())
		return
	}

	// for _, comment := range strings.Split(anime.Comment, "\r\n") {
	// 	fmt.Println(comment)
	// }
}
