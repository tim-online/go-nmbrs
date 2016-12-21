// +build !integration

package companies

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/tim-online/go-nmbrs/soap"
)

// https://api.nmbrs.nl/soap/v2.1/EmployeeService.asmx?op=HourComponentVar_Get
func TestList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Test meta data of incoming request
		testMethod(t, r, "POST")
		testHeader(t, r, "Content-Type", "text/xml; charset=utf-8")
		testHeader(t, r, "Soapaction", "https://api.nmbrs.nl/soap/v2.1/CompanyService/List_GetAll")

		// convert incoming request back to struct
		req := soap.NewRequest()
		got := newListRequest()
		req.Envelope.Body.Data = got
		err := xml.NewDecoder(r.Body).Decode(req.Envelope)
		if err != nil {
			t.Errorf("service.List sent invalid xml: %v", err)
		}

		// create the wanted request
		want := newListRequest()

		// compare them
		if !reflect.DeepEqual(got, want) {
			t.Errorf("service.List\n got=%#v\nwant=%#v", got, want)
		}

		// Send response as in API documentation
		w.Header().Add("Content-Type", "text/xml; charset=utf-8")
		fmt.Fprint(w, `<?xml version="1.0" encoding="utf-8"?>
			<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
				<soap:Body>
					<List_GetAllResponse xmlns="https://api.nmbrs.nl/soap/v2.1/CompanyService">
						<List_GetAllResult>
							<Company>
								<ID>123</ID>
								<Number>456</Number>
								<Name>Test 1</Name>
								<PhoneNumber>555-555-555</PhoneNumber>
								<FaxNumber>555-555-556</FaxNumber>
								<Email>test@example.com</Email>
								<Website>http://example.com</Website>
								<LoonaangifteTijdvak>Month</LoonaangifteTijdvak>
								<KvkNr>20153354</KvkNr>
							</Company>
							<Company>
								<ID>789</ID>
								<Number>101112</Number>
								<Name>Test 2</Name>
								<PhoneNumber>555-555-666</PhoneNumber>
								<FaxNumber>555-555-667</FaxNumber>
								<Email>test@example.com</Email>
								<Website>http://www.example2.com</Website>
								<LoonaangifteTijdvak>Year</LoonaangifteTijdvak>
								<KvkNr>20153354</KvkNr>
							</Company>
						</List_GetAllResult>
					</List_GetAllResponse>
				</soap:Body>
			</soap:Envelope>`)
	})

	got, err := service.List()
	if err != nil {
		t.Errorf("service.List returned error: %v", err)
	}

	// create wanted response
	want := &listResponse{
		Companies: []Company{
			Company{
				ID:                  "123",
				Number:              456,
				PhoneNumber:         "555-555-555",
				FaxNumber:           "555-555-556",
				Email:               "test@example.com",
				Website:             "http://example.com",
				LoonaangifteTijdvak: "Month",
				KvkNr:               "20153354",
			},
			Company{
				ID:                  "789",
				Number:              101112,
				PhoneNumber:         "555-555-666",
				FaxNumber:           "555-555-667",
				Email:               "test@example.com",
				Website:             "http://www.example2.com",
				LoonaangifteTijdvak: "Year",
				KvkNr:               "20153354",
			},
		},
	}

	// compare them
	if !reflect.DeepEqual(got, want) {
		t.Errorf("service.List\n got=%#v\nwant=%#v", got, want)
	}
}
