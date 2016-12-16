// +build integration
// go test -tags=integration

package companies

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/tim-online/go-nmbrs/auth"
)

var (
	authHeader *auth.AuthHeader
)

func setup() {
	authHeader = auth.NewAuthHeader()
	authHeader.Username = os.Getenv("NMBRS_USERNAME")
	authHeader.Token = os.Getenv("NMBRS_TOKEn")

	// mux = http.NewServeMux()
	// server = httptest.NewServer(mux)

	// // build oauth http client
	// accessToken = "mysecrettoken"
	// token := &oauth2.Token{AccessToken: accessToken}
	// ts := oauth2.StaticTokenSource(token)
	// oauthClient := oauth2.NewClient(oauth2.NoContext, ts)

	// client = NewClient(oauthClient)
	// url, _ := url.Parse(server.URL)
	// client.BaseURL = url
}

func teardown() {
	// server.Close()
}

func TestList(t *testing.T) {
	setup()
	defer teardown()

	service := NewService(authHeader)
	u, _ := url.ParseRequestURI(sandboxEndpoint)
	service.Endpoint = u

	response, err := service.List()
	if err != nil {
		t.Errorf("companies.List returned error: %v", err)
		return
	}

	if len(response.Companies) == 0 {
		t.Errorf("companies.List returned %d companies", len(response.Companies))
	}

	fmt.Printf("%+v\n", response)
}
