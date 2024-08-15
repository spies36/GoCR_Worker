package AmqpController

import "github.com/rabbitmq/amqp091-go"

func (client *AmqpClient) ConsumeFromQueue(handler func(<-chan amqp091.Delivery)) (err error) {

	deliveries, err := client.recChannel.Consume(
		client.recQueue.Name,
		"GoCR_Worker",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	go handler(deliveries)

	return nil
}
