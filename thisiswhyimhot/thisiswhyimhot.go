// Package thisiswhyimhot provides an API to thisiswhyimhot.herokuapp.com
package thisiswhyimhot

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// {"id":424892,"version":1,"createdAt":"2019-11-14T15:24:12.584Z","temp1":"20.45","rh1":"21.55","temp2":"20.81","rh2":"0.00","temp3":"19.68","rh3":"27.92","temp4":null,"rh4":null,"pressure1":"1015.08","pressure2":"0.00","pressure3":"0.00"}

// MinTemperatureReport as generated by thisiswhyimhot.
type MinTemperatureReport struct {
	// Time at which the report was created.
	Time time.Time `json:"createdAt"`

	// Temperatur in Celcius. The temp3 sensor is the most accurate, according
	// to Min.
	Temperature float64 `json:"temp3,string"`

	// Success is true if the fetch from the API went well.
	Success bool
}

const tiwihLatestAPI = "https://thisiswhyimhot.herokuapp.com/api/temperature/latest"

// Latest MinTemperatureReport.
func Latest() MinTemperatureReport {
	report := MinTemperatureReport{
		Time: time.Now(), // Default to the time here, in case the fetch goes poorly
	}
	data, err := fetchPayload()
	if err != nil {
		log.WithError(err).Warnf("Failed to fetch the payload from %q", tiwihLatestAPI)
		return report
	}
	if err := json.Unmarshal(data, &report); err != nil {
		log.WithError(err).Warnf("Failed to unmarshal the payload from %q", tiwihLatestAPI)
		return report
	}
	report.Success = true
	log.Printf("Successfully fetched: %+v", report)
	return report
}

func fetchPayload() ([]byte, error) {
	r, err := http.Get(tiwihLatestAPI)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return ioutil.ReadAll(r.Body)
}
