package shoboi

import (
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/parnurzeal/gorequest"
)

func get(url string) (resp gorequest.Response, body []byte, errs []error) {
	const maxTries = 5

	tryCount := 0
	tryDelay := 10 * time.Second

	for {
		resp, body, errs = gorequest.New().Get(url).EndBytes()

		if resp.StatusCode == http.StatusOK {
			break
		}

		tryCount++

		color.Red("Shoboi status code %d (#%d)", resp.StatusCode, tryCount)

		if tryCount > maxTries {
			break
		}

		time.Sleep(tryDelay)

		// Exponential falloff
		tryDelay *= 2
	}

	return resp, body, errs
}
