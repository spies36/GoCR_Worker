package main

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
	"github.com/spies36/GoCR_Worker/PreProcessing"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	croppedImage, err := PreProcessing.CropImage("../../Downloads/Bill of Lading.jpg")

	if err != nil {
		fmt.Println(err.Error())
	}

	client.SetImageFromBytes(croppedImage)
	fmt.Println(client.Languages)
	text, err := client.Text()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("HERE")
	fmt.Println(text)
}
