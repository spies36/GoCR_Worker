package AmqpController

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type amqpClientSuper struct {
	connection   *amqp.Connection
	recQueue     *amqp.Queue
	recQueueDead *amqp.Queue
	recChannel   *amqp.Channel
	pubChannel   *amqp.Channel
	pubQueue     *amqp.Queue
}

type AmqpClient struct {
	amqpClientSuper
	Address                string
	Host                   string
	Username               string
	Password               string
	RecQueueName           string
	RecQueueDeadLetterName string
	PubQueueName           string
}

/*Connect to AMQP*/
func (client *AmqpClient) Connect() (err error) {
	amqpConfig := amqp.Config{
		Properties: amqp.NewConnectionProperties(),
		SASL: []amqp.Authentication{
			&amqp.PlainAuth{Username: client.Username, Password: client.Password},
		},
		Vhost: client.Host,
	}
	connection, err := amqp.DialConfig(client.Address, amqpConfig)
	if err != nil {
		return err
	}

	client.connection = connection
	if client.RecQueueDeadLetterName != "" {
		err = client.setupRecQueue()
		if err != nil {
			return err
		}
	} else {
		err = client.setupRecQueueWithDeadLetter()
		if err != nil {
			return err
		}
	}

	err = client.setupPubQueue()
	if err != nil {
		return err
	}

	return nil
}

/*Destroy connection*/
func (client *AmqpClient) Destroy() (err error) {
	client.recChannel.Close()
	err = client.connection.Close()
	return
}

func (client *AmqpClient) setupRecQueue() (err error) {
	channel, err := client.connection.Channel()
	if err != nil {
		return err
	}
	client.recChannel = channel

	queue, err := channel.QueueDeclare(client.RecQueueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	client.recQueue = &queue

	return nil
}

func (client *AmqpClient) setupRecQueueWithDeadLetter() (err error) {
	channel, err := client.connection.Channel()
	if err != nil {
		return err
	}
	queueArgs := amqp.Table{
		"x-dead-letter-exchange": "amq.direct",
		"deadLetterRoutingKey":   client.RecQueueDeadLetterName,
	}
	queue, err := channel.QueueDeclare(client.RecQueueName, true, false, false, false, queueArgs)
	if err != nil {
		return err
	}
	client.recQueue = &queue

	deadLetterQueueArgs := amqp.Table{
		"x-dead-letter-exchange":    "amq.direct",
		"x-dead-letter-routing-key": client.RecQueueName,
		"x-time-to-live":            120000,
	}
	deadLetterQueue, err := channel.QueueDeclare(client.RecQueueDeadLetterName, true, false, false, false, deadLetterQueueArgs)
	if err != nil {
		return err
	}
	client.recQueueDead = &deadLetterQueue

	return nil
}

func (client *AmqpClient) setupPubQueue() (err error) {
	channel, err := client.connection.Channel()
	if err != nil {
		return err
	}
	client.pubChannel = channel

	queue, err := channel.QueueDeclare(client.PubQueueName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	client.pubQueue = &queue
	return nil
}
