// +build integration
// go test -tags=integration

package hours

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/tim-online/go-nmbrs/auth"
)

var (
	authHeader *auth.AuthHeaderWithDomain
	employeeID = 540897
	period     = 2
	year       = 2016
)

func setupIntegration() {
	authHeader = auth.NewAuthHeader()
	authHeader.Username = os.Getenv("NMBRS_USERNAME")
	authHeader.Token = os.Getenv("NMBRS_TOKEN")

	if authHeader.Username == "" && authHeader.Token == "" {
		panic("No username or token specified")
	}
}

func teardownIntegration() {
	// server.Close()
}

func TestIntegrationListFixed(t *testing.T) {
	setupIntegration()
	defer teardownIntegration()

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
		t.Errorf("Hours.ListFixed returned error: %v", err)
		return
	}

	if len(response.HourComponents) == 0 {
		t.Errorf("hours.ListFixed returned %d fixed hour components", len(response.HourComponents))
	}
}

func TestIntegrationListVar(t *testing.T) {
	setupIntegration()
	defer teardownIntegration()

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
		t.Errorf("Hours.ListVar returned error: %v", err)
		return
	}

	if len(response.HourComponents) == 0 {
		t.Errorf("hours.ListVar returned %d variable hour component", len(response.HourComponents))
	}
}

func TestIntegrationListFixedCurrent(t *testing.T) {
	setupIntegration()
	defer teardownIntegration()

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
		t.Errorf("Hours.ListFixedCurrent returned error: %v", err)
		return
	}

	if len(response.HourComponents) == 0 {
		t.Errorf("hours.ListFixedCurrent returned %d fixed hour component", len(response.HourComponents))
	}
}

func TestIntegrationListVarCurrent(t *testing.T) {
	setupIntegration()
	defer teardownIntegration()

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
		t.Errorf("Hours.ListVarCurrent returned error: %v", err)
		return
	}

	if len(response.HourComponents) == 0 {
		t.Errorf("hours.ListVarCurrent returned %d variable hour components", len(response.HourComponents))
	}
}
