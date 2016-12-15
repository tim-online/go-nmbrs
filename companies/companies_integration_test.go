// +build integration
// go test -tags=integration

package companies

import (
	"net/url"
	"testing"
)

func setup() {
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

	service := NewService()
	u, _ := url.ParseRequestURI(sandboxEndpoint)
	service.Endpoint = u

	_, err := service.List()
	if err != nil {
		t.Errorf("companies.List returned error: %v", err)
	}
}
