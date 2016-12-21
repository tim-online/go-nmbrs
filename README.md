# Go Nmbrs API client

go-nmbrs is an API client library for accessing the Nmbrs API via SOAP 1.1.

API documentation can be found here:
https://support.nmbrs.nl/hc/nl/categories/200216836-Nmbrs-API-for-developers

## Usage

``` go
import "github.com/tim-online/go-nmbrs"
```

### Request companies

``` go
// get username & password
username := os.Getenv("NMBRS_USERNAME")
token := os.Getenv("NMBRS_TOKEN")

// build client
client := nmbrs.NewClient(nil, username, token)
client.client.Debug = true

// request all companies this token has access to
resp, err := client.Companies.List()
if err != nil {
	panic(err)
}

companies := resp.Companies
```

### Request all employees for a company

``` go
import "github.com/tim-online/go-nmbrs/employees"

// get id of company
companyID := companies[0].ID

// request all employees for this company ID
resp2, err := client.Employees.ListByCompany(companyID, employees.All)
if err != nil {
	panic(err)
}
```
