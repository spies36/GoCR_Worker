package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/otiai10/gosseract/v2"
	"github.com/spies36/GoCR_Worker/PreProcessing"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	croppedImage, err := PreProcessing.CropImage("../../Downloads/Bill of Lading.jpg")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	writeJpegToDisc(croppedImage, "../../Downloads/croppedImg.jpg")

	binarizedImg, err := PreProcessing.Binarize("../../Downloads/croppedImg.jpg")
	if err != nil {
		fmt.Println(err.Error())
	}
	writeJpegToDisc(binarizedImg, "../../Downloads/cropBinImg.jpg")

	imgToProcess, err := imageToBytes(binarizedImg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client.SetImageFromBytes(imgToProcess)
	fmt.Println(client.Languages)
	text, err := client.Text()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(text)
}

func imageToBytes(img image.Image) ([]byte, error) {
	var b bytes.Buffer
	err := jpeg.Encode(&b, img, nil)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func writeJpegToDisc(img image.Image, path string) {

	imgByteArr, err := imageToBytes(img)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	os.WriteFile(path, imgByteArr, 0777)
}
