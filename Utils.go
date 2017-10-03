package shoboi

import (
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/parnurzeal/gorequest"
)

func get(url string) (resp gorequest.Response, body []byte, errs []error) {
	const maxTries = 10

	tryCount := 0
	tryDelay := 5 * time.Second

	for {
		resp, body, errs = gorequest.New().Get(url).EndBytes()

		if resp.StatusCode == http.StatusOK {
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

	return resp, body, errs
}
