package nmbrs

import (
	"net/http"
	"net/url"

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
	c.Wages = wages.NewService(authHeader)
	c.Wages.Client = c.client
	c.HourCodes = hourcodes.NewService(authHeader)
	c.HourCodes.Client = c.client
	c.Schedules = schedules.NewService(authHeader)
	c.Schedules.Client = c.client
	c.CostCenter = costcenter.NewService(authHeader)
	c.CostCenter.Client = c.client

	// DebtorService

	return c
}

func (c *Client) SetDebug(debug bool) {
	c.client.Debug = debug
}

func (c *Client) SetSandbox(sandbox bool) {
	if sandbox == true {
		u, _ := url.ParseRequestURI(companies.SandboxEndpoint)
		c.Companies.Endpoint = u

		u, _ = url.ParseRequestURI(costcenters.SandboxEndpoint)
		c.CostCenters.Endpoint = u

		u, _ = url.ParseRequestURI(employees.SandboxEndpoint)
		c.Employees.Endpoint = u

		u, _ = url.ParseRequestURI(days.SandboxEndpoint)
		c.Days.Endpoint = u

		u, _ = url.ParseRequestURI(hours.SandboxEndpoint)
		c.Hours.Endpoint = u

		u, _ = url.ParseRequestURI(wages.SandboxEndpoint)
		c.Wages.Endpoint = u

		u, _ = url.ParseRequestURI(hourcodes.SandboxEndpoint)
		c.HourCodes.Endpoint = u

		u, _ = url.ParseRequestURI(schedules.SandboxEndpoint)
		c.Schedules.Endpoint = u
	} else {
		u, _ := url.ParseRequestURI(companies.Endpoint)
		c.Companies.Endpoint = u

		u, _ = url.ParseRequestURI(costcenters.Endpoint)
		c.CostCenters.Endpoint = u

		u, _ = url.ParseRequestURI(employees.Endpoint)
		c.Employees.Endpoint = u

		u, _ = url.ParseRequestURI(days.Endpoint)
		c.Days.Endpoint = u

		u, _ = url.ParseRequestURI(hours.Endpoint)
		c.Hours.Endpoint = u

		u, _ = url.ParseRequestURI(wages.Endpoint)
		c.Wages.Endpoint = u

		u, _ = url.ParseRequestURI(hourcodes.Endpoint)
		c.HourCodes.Endpoint = u

		u, _ = url.ParseRequestURI(schedules.Endpoint)
		c.Schedules.Endpoint = u
	}
}
