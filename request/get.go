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
	StartTime     time.Time
	Url           string
	HttpStatus    int
	Nanoseconds   int64
	EndTime       time.Time
	StatusMessage string
}

func Get(data Data, client *http.Client) Result {
	start := time.Now()
	url := data.Url()

	request := createGetRequest(data)
	response, err := client.Do(request)
	end := time.Now()
	requestDuration := end.Sub(start).Nanoseconds()

	if err != nil {
		return newFailureResult(err, start, end, url, requestDuration)
	}
	defer response.Body.Close()

	return newSuccessResult(start, end, url, response.StatusCode, requestDuration)
}

func createGetRequest(data Data) *http.Request {
	request, _ := http.NewRequest("GET", data.Url(), nil)
	request.SetBasicAuth(data.Username, data.Password)
	return request
}

func newFailureResult(err error, startTime time.Time, endTime time.Time, url string, nanoseconds int64) Result {
	return Result{
		Status:        FAILURE,
		StartTime:     startTime,
		Url:           url,
		StatusMessage: err.Error(),
		Nanoseconds:   nanoseconds,
		EndTime:       endTime,
	}
}

func newSuccessResult(startTime time.Time, endTime time.Time, url string, httpStatus int, nanoseconds int64) Result {
	return Result{
		Status:      SUCCESS,
		StartTime:   startTime,
		Url:         url,
		HttpStatus:  httpStatus,
		Nanoseconds: nanoseconds,
		EndTime:     endTime,
	}
}

var statusStrings = [...]string{"S", "F"}

func (status Status) String() string {
	return statusStrings[status]
}
