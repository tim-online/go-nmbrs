// +build !integration

package schedules_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/tim-online/go-nmbrs/auth"
	"github.com/tim-online/go-nmbrs/schedules"
)

const (
	username = "info@tim-online.nl"
	password = "mysecret"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	// authHeader *auth.AuthHeaderWithDomain
	service *schedules.Service
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// build xml auth header
	authHeader := auth.NewAuthHeader()
	authHeader.Username = username
	authHeader.Token = password

	service = schedules.NewService(authHeader)

	// set custom http client
	service.Client.Client = &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	// set testing endpoint
	url, _ := url.Parse(server.URL)
	service.Endpoint = url

	// enable debugging
	// service.Client.Debug = true
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	got := r.Method
	if expected != got {
		t.Errorf("Request method = %v, expected %v", got, expected)
	}
}

func testHeader(t *testing.T, r *http.Request, key string, expected string) {
	got := r.Header.Get(key)
	if expected != got {
		t.Errorf("Request header %v = %v, expected %v", key, got, expected)
	}
}
