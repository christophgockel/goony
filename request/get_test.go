package request_test

import (
	"github.com/christophgockel/goony/request"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"
)

var server *httptest.Server
var client *http.Client

var _ = Describe("Request Get", func() {
	AfterEach(func() {
		server.Close()
	})

	It("returns the status of a request", func() {
		startServerWithResponseCode(200)

		result := request.Get("/route", "http://host", client)

		Expect(result.Status).To(Equal(request.SUCCESS))
	})

	It("tracks result data of a request", func() {
		startServerWithResponseCode(200)

		result := request.Get("/route", "http://host", client)

		Expect(result.Status).To(Equal(request.SUCCESS))
		Expect(result.StartTime.Second()).To(Equal(time.Now().Second()))
		Expect(result.HttpStatus).To(Equal(200))
		Expect(result.Url).To(Equal("http://host/route"))
		Expect(result.Nanoseconds).To(BeNumerically(">", 0))
		Expect(result.EndTime.Second()).To(BeNumerically(">", 0))
		Expect(result.StatusMessage).To(Equal(""))
	})

	It("returns failure results when the request times out", func() {
		startTimingOutServer()

		result := request.Get("/route", "http://host", client)

		Expect(result.Status).To(Equal(request.FAILURE))
		Expect(result.StartTime.Second()).To(Equal(time.Now().Second()))
		Expect(result.Url).To(Equal("http://host/route"))
		Expect(result.Nanoseconds).To(BeNumerically(">", 0))
		Expect(result.EndTime.Second()).To(BeNumerically(">", 0))
		Expect(strings.ToLower(result.StatusMessage)).To(ContainSubstring("timeout"))
	})

	Context("Status Type", func() {
		It("has string representations", func() {
			Expect(request.SUCCESS.String()).To(Equal("S"))
			Expect(request.FAILURE.String()).To(Equal("F"))
		})
	})
})

func startServerWithResponseCode(code int) {
	startServerWith(func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(code)
	})
}

func startTimingOutServer() {
	startServerWith(func(response http.ResponseWriter, request *http.Request) {
		time.Sleep(5)
	})

	client.Timeout = 1
}

func startServerWith(handler http.HandlerFunc) {
	server = httptest.NewServer(handler)

	transport := &http.Transport{
		Proxy: func(request *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	client = &http.Client{Transport: transport}
}
