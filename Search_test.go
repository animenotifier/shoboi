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
