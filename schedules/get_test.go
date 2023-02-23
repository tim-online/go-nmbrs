// +build !integration

package schedules_test

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/tim-online/go-nmbrs/lib"
	"github.com/tim-online/go-nmbrs/schedules"
	"github.com/tim-online/go-nmbrs/soap"
)

// https://api.nmbrs.nl/soap/v3/EmployeeService.asmx?op=HourComponentVar_Get
func TestGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Test meta data of incoming request
		testMethod(t, r, "POST")
		testHeader(t, r, "Content-Type", "text/xml; charset=utf-8")
		testHeader(t, r, "Soapaction", "https://api.nmbrs.nl/soap/v3/EmployeeService/Schedule_Get")

		// convert incoming request back to struct
		req := soap.NewRequest()
		got := schedules.NewGetRequest()
		req.Envelope.Body.Data = got
		err := xml.NewDecoder(r.Body).Decode(req.Envelope)
		if err != nil {
			t.Errorf("service.Get sent invalid xml: %v", err)
		}

		// create the wanted request
		want := schedules.NewGetRequest()
		want.EmployeeID = 666
		want.Period = 4
		want.Year = 1983

		// compare them
		if !reflect.DeepEqual(got, want) {
			t.Errorf("service.Get\n got=%#v\nwant=%#v", got, want)
		}

		// Send response as in API documentation
		w.Header().Add("Content-Type", "text/xml; charset=utf-8")
		fmt.Fprint(w, `<?xml version="1.0" encoding="utf-8"?>
			<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
				<soap:Body>
					<Schedule_GetResponse xmlns="https://api.nmbrs.nl/soap/v3/EmployeeService">
						<Schedule_GetResult>
							<HoursMonday>8.6</HoursMonday>
							<HoursTuesday>8.6</HoursTuesday>
							<HoursWednesday>8.6</HoursWednesday>
							<HoursThursday>8.6</HoursThursday>
							<HoursFriday>3.6</HoursFriday>
							<HoursSaturday>0.0</HoursSaturday>
							<HoursSunday>0.0</HoursSunday>
							<HoursMonday2>0.0</HoursMonday2>
							<HoursTuesday2>0.0</HoursTuesday2>
							<HoursWednesday2>0.0</HoursWednesday2>
							<HoursThursday2>0.0</HoursThursday2>
							<HoursFriday2>0.0</HoursFriday2>
							<HoursSaturday2>0.0</HoursSaturday2>
							<HoursSunday2>0.0</HoursSunday2>
							<ParttimePercentage>1.0</ParttimePercentage>
							<StartDate>2017-01-01T00:00:00</StartDate>
						</Schedule_GetResult>
					</Schedule_GetResponse>
				</soap:Body>
		</soap:Envelope>`)
	})

	req := schedules.NewGetRequest()
	req.EmployeeID = 666
	req.Period = 4
	req.Year = 1983
	got, err := service.Get(req)
	if err != nil {
		t.Errorf("service.Get returned error: %v", err)
	}

	// create wanted response
	// want := NewGetResponse()
	startDate := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
	want := &schedules.GetResponse{
		HoursMonday:        8.6,
		HoursTuesday:       8.6,
		HoursWednesday:     8.6,
		HoursThursday:      8.6,
		HoursFriday:        3.6,
		HoursSaturday:      0.0,
		HoursSunday:        0.0,
		HoursMonday2:       0.0,
		HoursTuesday2:      0.0,
		HoursWednesday2:    0.0,
		HoursThursday2:     0.0,
		HoursFriday2:       0.0,
		HoursSaturday2:     0.0,
		HoursSunday2:       0.0,
		ParttimePercentage: 1.0,
		StartDate:          &lib.Time{Time: startDate},
	}

	// compare them
	if !reflect.DeepEqual(got, want) {
		t.Errorf("service.Get\n got=%#v\nwant=%#v", got, want)
	}
}
