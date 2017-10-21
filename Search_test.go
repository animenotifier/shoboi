package shoboi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchAnime(t *testing.T) {
	result, err := SearchAnime("Attack on Titan")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.TID)
	assert.NotEmpty(t, result.TitleJapanese)
}

func TestSearchAnimeJP(t *testing.T) {
	result, err := SearchAnime("劇場版 進撃の巨人 前編～紅蓮の弓矢～")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.TID)
	assert.NotEmpty(t, result.TitleJapanese)
}
