// +build integration
// go test -tags=integration

package days

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/tim-online/go-nmbrs/auth"
)

var (
	authHeader *auth.AuthHeader
	employeeID = 540897
	period     = 2
	year       = 2016
)

func setup() {
	authHeader = auth.NewAuthHeader()
	authHeader.Username = os.Getenv("NMBRS_USERNAME")
	authHeader.Token = os.Getenv("NMBRS_TOKEN")

	if authHeader.Username == "" && authHeader.Token == "" {
		panic("No username or token specified")
	}
}

func teardown() {
	// server.Close()
}

func TestListFixed(t *testing.T) {
	setup()
	defer teardown()

	service := NewService(authHeader)

	// set custom http client
	service.Client.Client = &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	// set sandbox endpoint
	// u, _ := url.ParseRequestURI(sandboxEndpoint)
	// service.Endpoint = u

	// service.Client.Debug = true

	response, err := service.ListFixed(employeeID, period, year)
	if err != nil {
		t.Errorf("Days.ListFixed returned error: %v", err)
		return
	}

	if response.Days == 0 {
		t.Errorf("days.ListFixed returned %d fixed days", response.Days)
	}
}

func TestListVar(t *testing.T) {
	setup()
	defer teardown()

	service := NewService(authHeader)

	// set custom http client
	service.Client.Client = &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	// set sandbox endpoint
	// u, _ := url.ParseRequestURI(sandboxEndpoint)
	// service.Endpoint = u

	// service.Client.Debug = true

	response, err := service.ListVar(employeeID, period, year)
	if err != nil {
		t.Errorf("days.ListVar returned error: %v", err)
		return
	}

	if response.Days == 0 {
		t.Errorf("days.ListVar returned %d variable days", response.Days)
	}
}

func TestListFixedCurrent(t *testing.T) {
	setup()
	defer teardown()

	service := NewService(authHeader)

	// set custom http client
	service.Client.Client = &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	// set sandbox endpoint
	// u, _ := url.ParseRequestURI(sandboxEndpoint)
	// service.Endpoint = u

	// service.Client.Debug = true

	response, err := service.ListFixed(employeeID, period, year)
	if err != nil {
		t.Errorf("Days.ListFixed returned error: %v", err)
		return
	}

	if response.Days == 0 {
		t.Errorf("days.ListFixed returned %d fixed days", response.Days)
	}
}

func TestFixedCurrent(t *testing.T) {
	setup()
	defer teardown()

	service := NewService(authHeader)

	// set custom http client
	service.Client.Client = &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	// set sandbox endpoint
	// u, _ := url.ParseRequestURI(sandboxEndpoint)
	// service.Endpoint = u

	// service.Client.Debug = true

	response, err := service.ListFixedCurrent(employeeID)
	if err != nil {
		t.Errorf("days.ListFixedCurrent returned error: %v", err)
		return
	}

	if response.Days == 0 {
		t.Errorf("days.ListFixedCurrent returned %d fixed days", response.Days)
	}
}

func TestVarCurrent(t *testing.T) {
	setup()
	defer teardown()

	service := NewService(authHeader)

	// set custom http client
	service.Client.Client = &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	// set sandbox endpoint
	// u, _ := url.ParseRequestURI(sandboxEndpoint)
	// service.Endpoint = u

	// service.Client.Debug = true

	response, err := service.ListVarCurrent(employeeID)
	if err != nil {
		t.Errorf("days.ListVarCurrent returned error: %v", err)
		return
	}

	fmt.Println(response)

	if response.Days == 0 {
		t.Errorf("days.ListVarCurrent returned %d variable days", response.Days)
	}
}

func TestSetVarCurrent(t *testing.T) {
	setup()
	defer teardown()

	service := NewService(authHeader)

	// set custom http client
	service.Client.Client = &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	// set sandbox endpoint
	// u, _ := url.ParseRequestURI(sandboxEndpoint)
	// service.Endpoint = u

	// service.Client.Debug = true

	_, err := service.SetVarCurrent(employeeID, 16)
	if err != nil {
		t.Errorf("days.SetVarCurrent returned error: %v", err)
		return
	}
}
