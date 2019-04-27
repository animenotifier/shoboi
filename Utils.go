package shoboi

import (
	"net/http"
	"time"

	"github.com/aerogo/http/client"
	"github.com/akyoto/color"
)

func get(url string) (response *client.Response, err error) {
	const maxTries = 10

	tryCount := 0
	tryDelay := 5 * time.Second

	for {
		response, err = client.Get(url).End()

		if response.StatusCode() == http.StatusOK {
			break
		}

		tryCount++

		color.Red("Status code %d | Try #%d | %s", response.StatusCode(), tryCount, url)

		if tryCount > maxTries {
			break
		}

		time.Sleep(tryDelay)

		// Exponential falloff
		tryDelay += tryDelay / 2
	}

	return response, err
}
