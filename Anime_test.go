package shoboi

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEpisodes(t *testing.T) {
	anime, err := GetAnime("1536")

	assert.Nil(t, err)

	for _, episode := range anime.Episodes() {
		assert.NotZero(t, episode.Number)
		assert.NotEmpty(t, episode.TitleJapanese)

		airingDate := episode.AiringDate()

		assert.NotEmpty(t, airingDate.Start)
		assert.True(t, strings.HasPrefix(airingDate.Start, "2008") || strings.HasPrefix(airingDate.Start, "2009") || strings.HasPrefix(airingDate.Start, "2010"))
		assert.NotEmpty(t, airingDate.End)
		assert.True(t, strings.HasPrefix(airingDate.End, "2008") || strings.HasPrefix(airingDate.End, "2009") || strings.HasPrefix(airingDate.End, "2010"))
	}
}

func TestNonExistingAnime(t *testing.T) {
	anime, err := GetAnime("999999999999")

	assert.Nil(t, anime)
	assert.Error(t, err)
}
