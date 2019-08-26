package shoboi

import (
	"testing"

	"github.com/akyoto/assert"
)

func TestSearchAnime(t *testing.T) {
	result, err := SearchAnime("Attack on Titan")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotEqual(t, result.TID, "")
	assert.NotEqual(t, result.TitleJapanese, "")
}

func TestSearchAnimeJP(t *testing.T) {
	result, err := SearchAnime("劇場版 進撃の巨人 前編～紅蓮の弓矢～")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotEqual(t, result.TID, "")
	assert.NotEqual(t, result.TitleJapanese, "")
}
