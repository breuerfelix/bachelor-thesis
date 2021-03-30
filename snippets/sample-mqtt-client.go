package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
)

func checkHealth(host string, port int) (ret bool) {
	// return false if error happens
	defer func() {
		if r := recover(); r != nil {
			ret = false
		}
	}()

	topic := "healthcheck"
	client_id := "control-plane-healthcheck"
	timeout := 10 * time.Second

	start := time.Now().Unix()
	var remainingTime = func() time.Duration {
		elapsed := time.Now().Unix() - start
		return timeout - time.Duration(elapsed)
	}

	queue := make(chan string)
	error := false

	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		if msg.Topic() != topic {
			error = true
			return
		}

		queue <- string(msg.Payload())
	}

	opts := mqtt.NewClientOptions().
		AddBroker(fmt.Sprintf("tcp://%v:%v", host, port)).
		SetClientID(client_id)

	opts.SetDefaultPublishHandler(f)
	opts.SetKeepAlive(timeout)
	opts.SetPingTimeout(timeout)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.WaitTimeout(remainingTime()) &&
		token.Error() != nil {
		fmt.Println("Error connect:", token.Error())
		return false
	}

	defer c.Disconnect(0)

	if token := c.Subscribe(topic, 2, nil); token.WaitTimeout(remainingTime()) &&
		token.Error() != nil {
		fmt.Println("Error subscribe:", token.Error())
		return false
	}

	defer c.Unsubscribe(topic)

	text := fmt.Sprintf("hello from control plane")
	if token := c.Publish(topic, 2, false, text); token.WaitTimeout(remainingTime()) &&
		token.Error() != nil {
		fmt.Println("Error publish:", token.Error())
		return false
	}

	timeout_channel := make(chan bool, 1)

	go func() {
		time.Sleep(remainingTime())
		timeout_channel <- true
	}()

	select {
	case message := <-queue:
		if error {
			return false
		}

		return message == text
	case <-timeout_channel:
		return false
	}
}
