package main

import (
	"fmt"
	"github.com/brewneaux/volume-control-mqtt/internal/mqtt_client"
	"github.com/brewneaux/volume-control-mqtt/internal/volume_control"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var increment = 1

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
	vol, _ := volume_control.GetVolume()
	fmt.Printf(" current volume %v\n", vol)
	if string(message.Payload()) == "up" {
		currentVolume, err := volume_control.TurnUpVolume(increment)
		if err != nil {
			_ = fmt.Errorf("Error setting volume: %v", err)
		}
		fmt.Printf("Volume now %v\n", currentVolume)
	}
	if string(message.Payload()) == "down" {
		currentVolume, err := volume_control.TurnDownVolume(increment)
		if err != nil {
			_ = fmt.Errorf("Error setting volume: %v", err)
		}
		fmt.Printf("Volume now %v\n", currentVolume)
	}
}

func main() {
	incrementEnv := os.Getenv("VOLUME_INCREMENT")
	if incrementEnv == "" {
		increment = 1
	}
	increment, err := strconv.Atoi(incrementEnv)
	if err != nil {
		fmt.Printf("Error converting %v to int: %v", increment, err)
		increment = 1
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	opts, err := mqtt_client.GetMqttOpts()
	if err != nil {
		panic(err)
	}
	client := mqtt_client.GetClient(opts)
	mqtt_client.Subscribe(client, onMessageReceived)
	<-c

}
