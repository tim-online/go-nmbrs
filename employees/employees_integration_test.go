// +build integration
// go test -tags=integration

package employees

import (
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/tim-online/go-nmbrs/auth"
)

var (
	authHeader *auth.AuthHeaderWithDomain
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

func TestListByCompany(t *testing.T) {
	setup()
	defer teardown()

	service := NewService(authHeader)

	// set custom http client
	service.Client.Client = &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	// set sandbox endpoint
	u, _ := url.ParseRequestURI(sandboxEndpoint)
	service.Endpoint = u

	// service.Client.Debug = true

	response, err := service.ListByCompany(63749, All)
	if err != nil {
		t.Errorf("Employees.List returned error: %v", err)
		return
	}

	if len(response.Employees) == 0 {
		t.Errorf("employees.List returned %d employees", len(response.Employees))
	}
}
