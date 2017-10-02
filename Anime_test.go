package shoboi

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEpisodesWithTitles(t *testing.T) {
	anime, err := GetAnime("1536")
	assert.NoError(t, err)

	episodes := anime.Episodes()
	assert.Len(t, episodes, 7)

	for _, episode := range episodes {
		assert.NotZero(t, episode.Number)
		assert.NotEmpty(t, episode.TitleJapanese)

		airingDate := episode.AiringDate

		assert.NotEmpty(t, airingDate.Start)
		assert.True(t, strings.HasPrefix(airingDate.Start, "2008") || strings.HasPrefix(airingDate.Start, "2009") || strings.HasPrefix(airingDate.Start, "2010"))
		assert.NotEmpty(t, airingDate.End)
		assert.True(t, strings.HasPrefix(airingDate.End, "2008") || strings.HasPrefix(airingDate.End, "2009") || strings.HasPrefix(airingDate.End, "2010"))
	}
}

func TestEpisodesWithoutTitles(t *testing.T) {
	anime, err := GetAnime("4689")
	assert.NoError(t, err)

	episodes := anime.Episodes()
	assert.Len(t, episodes, 12)

	for _, episode := range episodes {
		assert.NotZero(t, episode.Number)

		airingDate := episode.AiringDate

		assert.NotEmpty(t, airingDate.Start)
		assert.True(t, strings.HasPrefix(airingDate.Start, "2017"))
		assert.NotEmpty(t, airingDate.End)
		assert.True(t, strings.HasPrefix(airingDate.End, "2017"))
	}
}

func TestEpisodesStartingWithZero(t *testing.T) {
	anime, err := GetAnime("3500")
	assert.NoError(t, err)

	episodes := anime.Episodes()
	assert.Len(t, episodes, 13)
	assert.Equal(t, episodes[0].Number, 0)

	for _, episode := range episodes {
		assert.True(t, episode.Number <= 12)
		assert.NotEmpty(t, episode.TitleJapanese)
	}
}

func TestNonExistingAnime(t *testing.T) {
	anime, err := GetAnime("999999999999")

	assert.Nil(t, anime)
	assert.Error(t, err)
}
