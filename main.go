package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/otiai10/gosseract/v2"
	"github.com/spies36/GoCR_Worker/FileHandling"
	"github.com/spies36/GoCR_Worker/PreProcessing"
)

const imgPath string = "../../Downloads/img.jpg"

func main() {
	ocr := gosseract.NewClient()
	defer ocr.Close()

	img, err := FileHandling.GetFileAsImg(imgPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	croppedImage := PreProcessing.CropImage(*img)
	writeJpegToDisc(croppedImage, "../../Downloads/croppedImg.jpg")

	imgToProcess, err := imageToBytes(*croppedImage)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ocr.SetImageFromBytes(imgToProcess)
	text, err := ocr.Text()
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

func writeJpegToDisc(img *image.Image, path string) {

	imgByteArr, err := imageToBytes(*img)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	os.WriteFile(path, imgByteArr, 0777)
}
