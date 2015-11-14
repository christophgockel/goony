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
	Time          time.Time
	Url           string
	HttpStatus    int
	Nanoseconds   int64
	StatusMessage string
}

func Get(path string, host string, client *http.Client) Result {
	url := host + path

	start := time.Now()
	response, err := client.Get(url)
	requestDuration := time.Since(start).Nanoseconds()

	if err != nil {
		return newFailureResult(err, start, url, requestDuration)
	}

	return newSuccessResult(start, url, response.StatusCode, requestDuration)
}

func newFailureResult(err error, time time.Time, url string, nanoseconds int64) Result {
	return Result{
		Status:        FAILURE,
		Time:          time,
		Url:           url,
		StatusMessage: err.Error(),
		Nanoseconds:   nanoseconds,
	}
}

func newSuccessResult(time time.Time, url string, httpStatus int, nanoseconds int64) Result {
	return Result{
		Status:      SUCCESS,
		Time:        time,
		Url:         url,
		HttpStatus:  httpStatus,
		Nanoseconds: nanoseconds,
	}
}

var statusStrings = [...]string{"S", "F"}

func (status Status) String() string {
	return statusStrings[status]
}
