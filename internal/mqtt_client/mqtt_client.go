package mqtt_client

import (
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func GetMqttOpts() (*mqtt.ClientOptions, error) {
	opts := mqtt.NewClientOptions()
	broker := os.Getenv("MQTT_BROKER")
	port := os.Getenv("MQTT_PORT")
	brokerUrl := fmt.Sprintf("tcp://%v:%v", broker, port)
	fmt.Printf("Connecting to %v", brokerUrl)
	opts.AddBroker(brokerUrl)
	username := os.Getenv("MQTT_USERNAME")
	password := os.Getenv("MQTT_PASSWORD")
	if (username != "" && password == "") || (username == "" && password != "") {
		return opts, errors.New("Invalid credentials supplied.")
	}
	if username != "" && password != "" {
		opts.SetUsername(username)
		opts.SetPassword(password)
	}
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	return opts, nil
}

func GetClient(opts *mqtt.ClientOptions) mqtt.Client {
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

func Subscribe(client mqtt.Client, callback func(client mqtt.Client, message mqtt.Message)) {
	topic := os.Getenv("MQTT_TOPIC")
	if topic == "" {
		panic("No topic given")
	}

	token := client.Subscribe(topic, 1, callback)
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)
}
