package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/mochi-mqtt/server/v2/packets"

	"github.com/conxtech/jsonhandler"
)

// Clear the terminal screen
func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	// Create signals channel to run server until interrupted
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	// Create the new MQTT Server.
	server := mqtt.New(&mqtt.Options{
		InlineClient: true,
	})

	// Allow all connections.
	_ = server.AddHook(new(auth.AllowHook), nil)

	// Create a TCP listener on a standard port.
	tcp := listeners.NewTCP(listeners.Config{ID: "t1", Address: ":1883"})
	err := server.AddListener(tcp)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err := server.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()

	callbackFn := func(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
		jsonString := string(pk.Payload)

		clearTerminal()

		fmt.Println("Received JSON")

		// Use jsonhandler to parse the JSON
		sensorData, err := jsonhandler.ParseJSON(jsonString)
		if err != nil {
			server.Log.Error("Failed to parse JSON", "error", err)
			return
		}

		// Access the fields
		fmt.Printf("Device ID: %d\n", sensorData.DeviceID)
		fmt.Printf("MAC Address: %s\n", sensorData.MACAddress)
		fmt.Printf("Timestamp: %d\n", sensorData.Timestamp)

		// PIR data
		fmt.Printf("PIR Sensor | Motion: %v, Changed: %v\n", sensorData.PIR.Motion, sensorData.PIR.Changed)

		// IAQ data
		fmt.Printf("IAQ Sensor | CO2: %d, TVOC: %d, Resistance: %.2f, Status: %d\n",
			sensorData.IAQ.CO2, sensorData.IAQ.TVOC, sensorData.IAQ.Resistance, sensorData.IAQ.Status)

		// AGS data
		fmt.Printf("AGS Sensor | Gas Resistance: %.2f, TVOC: %d\n",
			sensorData.AGS.GasResistance, sensorData.AGS.TVOC)

		// DHT data
		fmt.Printf("DHT Sensor | Temperature (C): %.2f, Humidity: %.2f\n",
			sensorData.DHT.TemperatureC, sensorData.DHT.Humidity)

		// PMS data
		fmt.Printf("PMS Sensor | PM1: %d, PM2.5: %d, PM10: %d\n",
			sensorData.PMS.PM1, sensorData.PMS.PM2_5, sensorData.PMS.PM10)
	}

	server.Log.Info("inline client subscribing")
	_ = server.Subscribe("sensors/data", 1, callbackFn)

	// Run server until interrupted
	<-done

	// Cleanup
}
