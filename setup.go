package main

import (
	"encoding/json"
	"os"
)

/*Format of config.json*/
type Config struct {
	//URL to amqp
	AMQP_Conn_String string
	//Host or VHost
	AMQP_Host     string
	AMQP_Username string
	AMQP_Pass     string
	//Queue name for receiving messages/consumption
	RecQueueName string
	//Optional dead letter queue
	RecQueueDeadLetterName string
	//Queue for publishing results
	PubQueueName string
}

func loadConfig() (*Config, error) {
	configFile, err := os.Open("./config.json")
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil

}
