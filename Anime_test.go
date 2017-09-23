package shoboi

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestEpisodes(t *testing.T) {
	anime, err := GetAnime("4515")

	assert.Nil(t, err)

	for _, episode := range anime.Episodes() {
		assert.NotEmpty(t, episode.TitleJapanese)

		fmt.Printf("Episode %d: %s\n", episode.Number, color.GreenString(episode.TitleJapanese))

		airingDate := episode.AiringDate()
		fmt.Println(airingDate)

		assert.NotEmpty(t, airingDate.Start)
		assert.NotEmpty(t, airingDate.End)
	}
}

func TestNonExistingAnime(t *testing.T) {
	anime, err := GetAnime("999999999999")

	assert.Nil(t, anime)
	assert.Error(t, err)
}
