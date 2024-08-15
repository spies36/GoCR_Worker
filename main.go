package main

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spies36/GoCR_Worker/AmqpController"

	"github.com/spies36/GoCR_Worker/FileHandling"
	"github.com/spies36/GoCR_Worker/PreProcessing"
	"gopkg.in/gographics/imagick.v2/imagick"
)

func main() {
	//Parse config
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Error loading env: " + err.Error())
	}

	//Startup AMQP
	amqpClient := AmqpController.AmqpClient{
		Protocol:               config.AMQP_Protocol,
		Port:                   config.AMQP_Port,
		Host:                   config.AMQP_Host,
		Username:               config.AMQP_Username,
		Password:               config.AMQP_Pass,
		RecQueueName:           config.RecQueueName,
		RecQueueDeadLetterName: config.RecQueueDeadLetterName,
		PubQueueName:           config.PubQueueName,
	}
	amqpClient.Connect()
	defer amqpClient.Destroy()

	//Stand up the OCR engine
	imagick.Initialize()
	defer imagick.Terminate()

	startConsumer(amqpClient)
}

func startConsumer(amqpClient AmqpController.AmqpClient) {
	amqpClient.ConsumeFromQueue(handleMessage)
}

func handleMessage(deliveries <-chan amqp.Delivery) {

	for delivery := range deliveries {

		var tryCount int
		if deathCount, ok := delivery.Headers["x-death"].(int); ok {
			tryCount = deathCount
		}

		msg, err := parseMessage(delivery.Body)
		if err != nil { //Message failed to parse
			if tryCount > 2 {
				delivery.Ack(false)
			}
			delivery.Nack(false, false) //dead letter the message
		}

		txtFound, err := processImage(msg.ImgPath)
		if err != nil {
			if tryCount > 2 {
				delivery.Ack(false)
			}
			delivery.Nack(false, false) //dead letter the message
		}

		fmt.Println(txtFound)
		delivery.Ack(false)
	}

}

func processImage(imgPath string) (string, error) {
	ocr := gosseract.NewClient()
	defer ocr.Close()
	img, err := FileHandling.GetFileAsImg(imgPath)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	croppedImage := PreProcessing.CropImage(*img)

	imgToProcess, err := FileHandling.PngImageToBytes(*croppedImage)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	ocr.SetImageFromBytes(imgToProcess)
	text, err := ocr.Text()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return text, nil
}
