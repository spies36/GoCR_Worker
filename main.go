package main

import (
	"fmt"
	"time"

	"github.com/otiai10/gosseract/v2"
	"github.com/spies36/GoCR_Worker/AmqpController"
	"github.com/spies36/GoCR_Worker/FileHandling"
	"github.com/spies36/GoCR_Worker/PreProcessing"
	"gopkg.in/gographics/imagick.v2/imagick"
)

// const imgPath1 string = "/home/spies/Downloads/img.jpg"
const imgPath2 string = "/home/spies/Downloads/17730027.pdf"

// const imgPath3 string = "/home/spies/Downloads/17730033.pdf"

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Error loading env: " + err.Error())
	}
	amqpClient := AmqpController.AmqpClient{
		Address:                config.AMQP_Conn_String,
		Host:                   config.AMQP_Host,
		Username:               config.AMQP_Username,
		Password:               config.AMQP_Pass,
		RecQueueName:           config.RecQueueName,
		RecQueueDeadLetterName: config.RecQueueDeadLetterName,
		PubQueueName:           config.PubQueueName,
	}
	amqpClient.Connect()
	defer amqpClient.Destroy()

	imagick.Initialize()
	defer imagick.Terminate()

	startTime := time.Now()
	processImage(imgPath2)
	fmt.Println("All Done")
	fmt.Printf("\n Total Time: %f \n", time.Since(startTime).Seconds())
}

func processImage(imgPath string) (string, error) {
	preProcTimeStart := time.Now()
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
	preProcTimeEnd := time.Since(preProcTimeStart)
	ocrTimeStart := time.Now()
	ocr.SetImageFromBytes(imgToProcess)
	text, err := ocr.Text()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	ocrTimeFinished := time.Since(ocrTimeStart)
	fmt.Printf("\n ----- Time results ----- \n")
	fmt.Printf("Preprocessing: %f \n", preProcTimeEnd.Seconds())
	fmt.Printf("OCR: %f \n", ocrTimeFinished.Seconds())
	fmt.Printf("Total time: %f \n", (ocrTimeFinished + preProcTimeEnd).Seconds())
	return text, nil
}
