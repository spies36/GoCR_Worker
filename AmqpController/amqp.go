package AmqpController

import (
	"context"
	"fmt"

	"github.com/Azure/go-amqp"
)

type amqpClientSuper struct {
	connection *amqp.Conn
}

type AmqpClient struct {
	amqpClientSuper
	Address string
	Host    string
}

type Handler func(*amqp.Message, *amqp.Receiver)

func (client *AmqpClient) Connect() (err error) {
	connOptions := amqp.ConnOptions{
		HostName: client.Host,
	}
	connection, err := amqp.Dial(context.TODO(), client.Address, &connOptions)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	client.connection = connection
	return
}

func (client *AmqpClient) Destroy() (err error) {
	err = client.connection.Close()
	return
}

func (client *AmqpClient) AddConsumer(queueName string, handler Handler) (err error) {
	session, err := client.connection.NewSession(context.TODO(), nil)
	if err != nil {
		return err
	}

	receiverOpts := amqp.ReceiverOptions{Credit: 1}
	consumer, err := session.NewReceiver(context.TODO(), queueName, &receiverOpts)

	go beginConsumption(handler, consumer)

	return err
}

func beginConsumption(handler Handler, consumer *amqp.Receiver) {

}
