package request

import (
	"net/http"
	"time"
)

type Status int

const (
	SUCCESS Status = 0 + iota
	FAILURE Status = 0 + iota
)

type Result struct {
	Status        Status
	Url           string
	HttpStatus    int
	Nanoseconds   int64
	StatusMessage string
}

func Get(path string, host string, client *http.Client) Result {
	url := host + path

	start := time.Now()
	response, err := client.Get(url)
	requestTime := time.Since(start).Nanoseconds()

	if err != nil {
		return newFailureResult(err, url)
	}

	return newSuccessResult(url, response.StatusCode, requestTime)
}

func newFailureResult(err error, url string) Result {
	return Result{
		Status:        FAILURE,
		Url:           url,
		StatusMessage: err.Error(),
	}
}

func newSuccessResult(url string, httpStatus int, nanoseconds int64) Result {
	return Result{
		Status:      SUCCESS,
		Url:         url,
		HttpStatus:  httpStatus,
		Nanoseconds: nanoseconds,
	}
}

var statusStrings = [...]string{"S", "F"}

func (status Status) String() string {
	return statusStrings[status]
}
