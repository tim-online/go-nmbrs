package nmbrs

import (
	"fmt"
	"os"
	"testing"

	"github.com/tim-online/go-nmbrs/employees"
)

func TestDing(t *testing.T) {
	// get username & password
	username := os.Getenv("NMBRS_USERNAME")
	token := os.Getenv("NMBRS_TOKEN")

	// build client
	client := NewClient(nil, username, token)
	client.client.Debug = true

	// request all companies this token has access to
	resp, err := client.Companies.List()
	if err != nil {
		panic(err)
	}

	companies := resp.Companies

	// ------------------------

	companyID := companies[0].ID
	resp2, err := client.Employees.ListByCompany(companyID, employees.All)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp2)
}
