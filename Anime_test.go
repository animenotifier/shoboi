package shoboi

import (
	"strings"
	"testing"

	"github.com/akyoto/assert"
)

func TestEpisodesWithTitles(t *testing.T) {
	anime, err := GetAnime("663")
	assert.Nil(t, err)

	episodes := anime.Episodes()
	assert.Equal(t, len(episodes), 24)

	for _, episode := range episodes {
		assert.NotEqual(t, episode.Number, 0)
		assert.NotEqual(t, episode.TitleJapanese, "")

		airingDate := episode.AiringDate

		assert.NotEqual(t, airingDate.Start, "")
		assert.True(t, strings.HasPrefix(airingDate.Start, "2006"))
		assert.NotEqual(t, airingDate.End, "")
		assert.True(t, strings.HasPrefix(airingDate.End, "2006"))
	}
}

func TestEpisodesWithoutTitles(t *testing.T) {
	anime, err := GetAnime("4689")
	assert.Nil(t, err)

	episodes := anime.Episodes()
	assert.Equal(t, len(episodes), 12)

	for _, episode := range episodes {
		assert.NotEqual(t, episode.Number, 0)

		airingDate := episode.AiringDate

		assert.NotEqual(t, airingDate.Start, "")
		assert.True(t, strings.HasPrefix(airingDate.Start, "2017"))
		assert.NotEqual(t, airingDate.End, "")
		assert.True(t, strings.HasPrefix(airingDate.End, "2017"))
	}
}

func TestEpisodesStartingWithZero(t *testing.T) {
	anime, err := GetAnime("3500")
	assert.Nil(t, err)

	episodes := anime.Episodes()
	assert.Equal(t, len(episodes), 13)
	assert.Equal(t, episodes[0].Number, 0)

	for _, episode := range episodes {
		assert.True(t, episode.Number <= 12)
		assert.NotEqual(t, episode.TitleJapanese, "")
	}
}

func TestNonExistingAnime(t *testing.T) {
	anime, err := GetAnime("999999999999")

	assert.Nil(t, anime)
	assert.NotNil(t, err)
}
