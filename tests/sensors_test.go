package tests

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type PIR struct {
	Motion  bool `json:"motion"`
	Changed bool `json:"changed"`
}

type IAQ struct {
	CO2        int `json:"co2"`
	TVOC       int `json:"tvoc"`
	Resistance int `json:"resistance"`
	Status     int `json:"status"`
}

type AGS struct {
	GasResistance float64 `json:"gasResistance"`
	TVOC          int     `json:"tvoc"`
}

type DHT struct {
	TemperatureC float64 `json:"temperatureC"`
	TemperatureF float64 `json:"temperatureF"`
	Humidity     float64 `json:"humidity"`
	HeatIndexC   float64 `json:"heatIndexC"`
	HeatIndexF   float64 `json:"heatIndexF"`
}

type PMS struct {
	PM1  int `json:"pm1"`
	PM25 int `json:"pm2_5"`
	PM10 int `json:"pm10"`
}

type SensorData struct {
	DeviceID   int    `json:"device_id"`
	MACAddress string `json:"MAC_Address"`
	Timestamp  int64  `json:"timestamp"`
	PIR        PIR    `json:"pir"`
	IAQ        IAQ    `json:"iaq"`
	AGS        AGS    `json:"ags"`
	DHT        DHT    `json:"dht"`
	PMS        PMS    `json:"pms"`
}

var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received message on topic: %s\n", msg.Topic())
	fmt.Printf("Message: %s\n", msg.Payload())
}

func TestSensors(t *testing.T) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883") // MQTT Broker (Mosquitto)
	opts.SetClientID("test_mqtt_client")
	opts.SetDefaultPublishHandler(messagePubHandler)

	// Connect to the MQTT broker
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	// Simulate sensor data
	for {
		data := SensorData{
			DeviceID:   13504477,
			MACAddress: "EC64C9CE0FDD",
			Timestamp:  time.Now().Unix(),
			PIR: PIR{
				Motion:  false,
				Changed: false,
			},
			IAQ: IAQ{
				CO2:        869,
				TVOC:       241,
				Resistance: 239012,
				Status:     0,
			},
			AGS: AGS{
				GasResistance: 18962.7,
				TVOC:          938,
			},
			DHT: DHT{
				TemperatureC: 27.9,
				TemperatureF: 82.22,
				Humidity:     50.1,
				HeatIndexC:   28.34619,
				HeatIndexF:   83.02364,
			},
			PMS: PMS{
				PM1:  2,
				PM25: 3,
				PM10: 3,
			},
		}
		// Serialize data to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error marshaling data:", err)
			return
		}

		// Publish to MQTT topic
		token := client.Publish("sensors/data", 0, false, jsonData)
		token.Wait()

		fmt.Printf("Sent data: %s\n", string(jsonData))

		time.Sleep(5 * time.Second) // Delay between messages
	}

	client.Disconnect(250)
}
