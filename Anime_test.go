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
		fmt.Printf("Episode %d: %s\n", episode.Number, color.GreenString(episode.TitleJapanese))
		fmt.Println(episode.AiringDate())
	}
}
