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
	domain := os.Getenv("NMBRS_DOMAIN")

	// build client
	client := NewClient(nil, username, token, domain)
	client.client.Debug = true

	resp, err := client.Companies.List()
	if err != nil {
		t.Error(err)
		return
	}

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

	// companyID := 65754
	// startPeriod := 1
	// endPeriod := 5
	// year := 2022
	// resp, err := client.Reports.JournalsReportByCompanyBackground(companyID, startPeriod, endPeriod, year)
	// if err != nil {
	// 	panic(err)
	// }

	// taskID := "ef5bc4d9-17bd-4c94-9e41-2cec055b1209"
	// resp, err := client.Reports.BackgroundTaskResult(taskID)
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

	// resp, err := client.Companies.ReportGetPayslipByRunCompanyV2(65756, 101453, 2020)
	// if err != nil {
	// 	panic(err)
	// }

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
