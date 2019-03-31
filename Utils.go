package shoboi

import (
	"net/http"
	"time"

	"github.com/aerogo/http/client"
	"github.com/blitzprog/color"
)

func get(url string) (resp client.Response, err error) {
	const maxTries = 10

	tryCount := 0
	tryDelay := 5 * time.Second

	for {
		resp, err = client.Get(url).End()

		if resp.StatusCode() == http.StatusOK {
			break
		}

		tryCount++

		color.Red("Status code %d | Try #%d | %s", resp.StatusCode, tryCount, url)

		if tryCount > maxTries {
			break
		}

		time.Sleep(tryDelay)

		// Exponential falloff
		tryDelay += tryDelay / 2
	}

	return resp, err
}
