package nmbrs

import (
	"net/http"

	"github.com/tim-online/go-nmbrs/auth"
	"github.com/tim-online/go-nmbrs/companies"
	"github.com/tim-online/go-nmbrs/costcenter"
	"github.com/tim-online/go-nmbrs/costcenters"
	"github.com/tim-online/go-nmbrs/days"
	"github.com/tim-online/go-nmbrs/employees"
	"github.com/tim-online/go-nmbrs/hourcodes"
	"github.com/tim-online/go-nmbrs/hours"
	"github.com/tim-online/go-nmbrs/schedules"
	"github.com/tim-online/go-nmbrs/soap"
	"github.com/tim-online/go-nmbrs/wages"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-nmbrs/" + libraryVersion
)

// Client manages communication with Nmbrs API
type Client struct {
	// SOAP client used to communicate with the API.
	client *soap.Client

	username string
	token    string

	// Services used for communicating with the API
	Companies   *companies.Service
	Employees   *employees.Service
	Days        *days.Service
	Hours       *hours.Service
	Wages       *wages.Service
	CostCenters *costcenters.Service
	HourCodes   *hourcodes.Service
	Schedules   *schedules.Service
	CostCenter  *costcenter.Service
}

// NewClient returns a new Nmbrs API client
func NewClient(httpClient *http.Client, username string, token string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		client: &soap.Client{
			Client:    httpClient,
			UserAgent: userAgent,
			Debug:     false,
		},
	}

	authHeader := auth.NewAuthHeader()
	authHeader.Username = username
	authHeader.Token = token

	// CompanyService
	c.Companies = companies.NewService(authHeader)
	c.Companies.Client = c.client
	c.CostCenters = costcenters.NewService(authHeader)
	c.CostCenters.Client = c.client

	// EmployeeService
	c.Employees = employees.NewService(authHeader)
	c.Employees.Client = c.client
	c.Days = days.NewService(authHeader)
	c.Days.Client = c.client
	c.Hours = hours.NewService(authHeader)
	c.Hours.Client = c.client
	c.HourCodes = hourcodes.NewService(authHeader)
	c.HourCodes.Client = c.client

	// DebtorService

	return c
}

// Client manages communication with Nmbrs API
type Client struct {
	// SOAP client used to communicate with the API.
	client *soap.Client

	username string
	token    string

	// Services used for communicating with the API
	Companies   *companies.Service
	Employees   *employees.Service
	Days        *days.Service
	Hours       *hours.Service
	CostCenters *costcenters.Service
	HourCodes   *hourcodes.Service
}
