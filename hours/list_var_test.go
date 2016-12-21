// +build !integration

package hours

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/tim-online/go-nmbrs/soap"
)

// https://api.nmbrs.nl/soap/v2.1/EmployeeService.asmx?op=HourComponentVar_Get
func TestListVar(t *testing.T) {
	setup()
	defer teardown()

	employeeID := 1200
	period := 2
	year := 2016

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Test meta data of incoming request
		testMethod(t, r, "POST")
		testHeader(t, r, "Content-Type", "text/xml; charset=utf-8")
		testHeader(t, r, "Soapaction", "https://api.nmbrs.nl/soap/v2.1/EmployeeService/HourComponentVar_Get")

		// convert incoming request back to struct
		req := soap.NewRequest()
		got := newListVarRequest(0, 0, 0)
		req.Envelope.Body.Data = got
		err := xml.NewDecoder(r.Body).Decode(req.Envelope)
		if err != nil {
			t.Errorf("service.ListVar sent invalid xml: %v", err)
		}

		// create the wanted request
		want := newListVarRequest(employeeID, period, year)

		// compare them
		if !reflect.DeepEqual(got, want) {
			t.Errorf("service.ListVar\n got=%#v\nwant=%#v", got, want)
		}

		// Send response as in API documentation
		w.Header().Add("Content-Type", "text/xml; charset=utf-8")
		fmt.Fprint(w, ` <?xml version="1.0" encoding="utf-8"?>
			<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
				<soap:Body>
					<HourComponentVar_GetResponse xmlns="https://api.nmbrs.nl/soap/v2.1/EmployeeService">
						<HourComponentVar_GetResult>
							<HourComponent>
								<Id>12</Id>
								<HourCode>13</HourCode>
								<Hours>14.1</Hours>
							</HourComponent>
							<HourComponent>
								<Id>15</Id>
								<HourCode>16</HourCode>
								<Hours>17.2</Hours>
							</HourComponent>
						</HourComponentVar_GetResult>
					</HourComponentVar_GetResponse>
				</soap:Body>
			</soap:Envelope>`)
	})

	got, err := service.ListVar(employeeID, period, year)
	if err != nil {
		t.Errorf("service.ListVar returned error: %v", err)
	}

	// create wanted response
	want := &listVarResponse{
		HourComponents: []HourComponent{
			HourComponent{
				ID:       12,
				HourCode: 13,
				Hours:    14.1,
			},
			HourComponent{
				ID:       15,
				HourCode: 16,
				Hours:    17.2,
			},
		},
	}

	// compare them
	if !reflect.DeepEqual(got, want) {
		t.Errorf("service.ListVar\n got=%#v\nwant=%#v", got, want)
	}
}
