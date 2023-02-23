// +build integration
// go test -tags=integration

package companies

import (
	"net/http"
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

func TestList(t *testing.T) {
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

	response, err := service.List()
	if err != nil {
		t.Errorf("companies.List returned error: %v", err)
		return
	}

	if len(response.Companies) == 0 {
		t.Errorf("companies.List returned %d companies", len(response.Companies))
	}
}
