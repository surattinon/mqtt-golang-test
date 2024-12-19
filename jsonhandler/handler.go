package jsonhandler

import (
	"encoding/json"
	"errors"
)

// PIRData represents the "pir" object in the JSON
type PIRData struct {
	IsValid    bool `json:"isValid"`
	IsMotion   bool `json:"isMotion"`
	HasChanged bool `json:"hasChanged"`
	LastTime   int  `json:"lastTime"`
	Status     int  `json:"status"`
}

type IAQData struct {
	IsValid    int     `json:"isValid"`
	CO2        int     `json:"co2"`
	TVOC       int     `json:"tvoc"`
	Resistance float64 `json:"ohms"`
	Status     int     `json:"status"`
}

type AGSData struct {
	IsValid    int     `json:"isValid"`
	TVOC       int     `json:"tvoc"`
	Resistance float64 `json:"ohms"`
	Status     int     `json:"status"`
}

type DHTData struct {
	IsValid      int     `json:"isValid"`
	TemperatureC float64 `json:"tempC"`
	Humidity     float64 `json:"humid"`
	Status       int     `json:"status"`
}

type PMSData struct {
	IsValid int `json:"isValid"`
	PM1     int `json:"pm1"`
	PM2_5   int `json:"pm25"`
	PM10    int `json:"pm10"`
	PM1ATM  int `json:"pm1Atm"`
	PM25ATM int `json:"pm25Atm"`
	PM10ATM int `json:"pm10Atm"`
	Status  int `json:"status"`
}

// SensorData represents the entire JSON structure
type SensorData struct {
	DeviceID   int     `json:"devId"`
	MACAddress string  `json:"mac"`
	UpTime     int64   `json:"upTime"`
	PIR        PIRData `json:"pir"`
	IAQ        IAQData `json:"iaq"`
	AGS        AGSData `json:"ags"`
	DHT        DHTData `json:"dht"`
	PMS        PMSData `json:"pms"`
}

// ParseJSON parses the JSON string into a SensorData struct
func ParseJSON(jsonString string) (SensorData, error) {
	var data SensorData
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return SensorData{}, errors.New("failed to parse JSON: " + err.Error())
	}
	return data, nil
}
