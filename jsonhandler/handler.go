package jsonhandler

import (
	"encoding/json"
	"errors"
)

// PIRData represents the "pir" object in the JSON
type PIRData struct {
	Motion  bool `json:"motion"`
	Changed bool `json:"changed"`
}

// IAQData represents the "iaq" object in the JSON
type IAQData struct {
	CO2        int     `json:"co2"`
	TVOC       int     `json:"tvoc"`
	Resistance float64 `json:"resistance"`
	Status     int     `json:"status"`
}

// AGSData represents the "ags" object in the JSON
type AGSData struct {
	GasResistance float64 `json:"gasResistance"`
	TVOC          int     `json:"tvoc"`
}

// DHTData represents the "dht" object in the JSON
type DHTData struct {
	TemperatureC float64 `json:"temperatureC"`
	TemperatureF float64 `json:"temperatureF"`
	Humidity     float64 `json:"humidity"`
	HeatIndexC   float64 `json:"heatIndexC"`
	HeatIndexF   float64 `json:"heatIndexF"`
}

// PMSData represents the "pms" object in the JSON
type PMSData struct {
	PM1   int `json:"pm1"`
	PM2_5 int `json:"pm2_5"`
	PM10  int `json:"pm10"`
}

// SensorData represents the entire JSON structure
type SensorData struct {
	DeviceID   int     `json:"device_id"`
	MACAddress string  `json:"MAC_Address"`
	Timestamp  int64   `json:"timestamp"`
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
