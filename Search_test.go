package shoboi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchAnime(t *testing.T) {
	tid, err := SearchAnime("Attack on Titan")
	assert.Nil(t, err)
	fmt.Println(tid)
}
