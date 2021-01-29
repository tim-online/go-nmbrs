package nmbrs

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDing(t *testing.T) {
	// get username & password
	username := os.Getenv("NMBRS_USERNAME")
	token := os.Getenv("NMBRS_TOKEN")

	// build client
	client := NewClient(nil, username, token)
	client.client.Debug = true

	// request all companies this token has access to
	// companyID := 65757
	// startYear := 2017
	// startPeriod := 12
	// endPeriod := 1
	// endYear := 2018
	// year := 2018

	// runID := 11048
	// resp, err := client.Companies.WageCodesByRunCompanyV2(companyID, runID, year)
	// if err != nil {
	// 	panic(err)
	// }

	// name := "test.txt"
	// subFolder := ""
	// body := base64.StdEncoding.EncodeToString([]byte("Dit is een test van Leon"))
	// log.Println(body)
	// resp, err := client.FileExplorer.UploadFile(companyID, name, subFolder, []byte(body))
	// if err != nil {
	// 	panic(err)
	// }

	// resp, err := client.Debtors.Department_GetList(12643970)
	// if err != nil {
	// 	panic(err)
	// }

	// resp, err := client.Employees.Department_GetAll_AllEmployeesByCompany(65755)
	// if err != nil {
	// 	panic(err)
	// }

	// resp, err := client.Employees.Department_GetCurrent(579512)
	// if err != nil {
	// 	panic(err)
	// }

	// resp, err := client.Reports.BusinessCompanyEmployeeWageComponentsPerRun(companyID, year)
	// if err != nil {
	// 	panic(err)
	// }

	// resp, err := client.Companies.RunGetList(companyID, year)
	// if err != nil {
	// 	panic(err)
	// }

	// resp, err := client.Companies.CostUnit_GetList(65757)
	// if err != nil {
	// 	panic(err)
	// }

	resp, err := client.Companies.ReportGetPayslipByRunCompanyV2(65756, 101453, 2020)
	if err != nil {
		panic(err)
	}

	// employeeID := 603463
	// name := "test.txt"
	// body := base64.StdEncoding.EncodeToString([]byte("Dit is een test van Leon"))
	// documentType := "c21d471a-2d8c-47a7-b8c4-9bde8d188b13"
	// log.Println(body)
	// resp, err := client.Employees.UploadDocument(employeeID, name, []byte(body), documentType)
	// if err != nil {
	// 	panic(err)
	// }

	b, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println(string(b))

	// for _, s := range resp.BusinessEmployeeHoursSalary.EmployeeHoursSalaryItems {
	// 	log.Printf("%+v", s)
	// }
}
