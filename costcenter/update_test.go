// +build !integration

package costcenter

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"

	"github.com/tim-online/go-nmbrs/soap"
)

func TestUpdate(t *testing.T) {
	setup()
	defer teardown()

	log.Println("aaah")

	updateRequest := NewUpdateRequest()
	updateRequest.EmployeeID = 1200
	updateRequest.CostCenters = []EmployeeCostCenter{
		EmployeeCostCenter{
			CostCenter: CostCenter{
				Code:        "12",
				Description: "test",
			},
			Percentage: 0.12,
			Default:    false,
		},
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Test meta data of incoming request
		testMethod(t, r, "POST")
		testHeader(t, r, "Content-Type", "text/xml; charset=utf-8")
		testHeader(t, r, "Soapaction", "https://api.nmbrs.nl/soap/v2.1/EmployeeService/CostCenter_Update")

		// convert incoming request back to struct
		req := soap.NewRequest()
		got := NewUpdateRequest()
		req.Envelope.Body.Data = got
		err := xml.NewDecoder(r.Body).Decode(req.Envelope)
		if err != nil {
			t.Errorf("service.Update sent invalid xml: %v", err)
		}

		// create the wanted request
		want := updateRequest

		// compare them
		if !reflect.DeepEqual(got, want) {
			t.Errorf("service.Update\n got=%#v\nwant=%#v", got, want)
		}

		// Send response as in API documentation
		w.Header().Add("Content-Type", "text/xml; charset=utf-8")
		fmt.Fprint(w, `<?xml version="1.0" encoding="utf-8"?>
			<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
				<soap:Body>
					<CostCenter_UpdateResponse xmlns="https://api.nmbrs.nl/soap/v2.1/EmployeeService" />
				</soap:Body>
		</soap:Envelope>`)
	})

	got, err := service.Update(updateRequest)
	if err != nil {
		t.Errorf("service.Update returned error: %v", err)
	}

	// create wanted response
	want := &UpdateResponse{}

	// compare them
	if !reflect.DeepEqual(got, want) {
		t.Errorf("service.Update\n got=%#v\nwant=%#v", got, want)
	}
}
