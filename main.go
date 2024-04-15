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
	"gopkg.in/gographics/imagick.v2/imagick"
)

const imgPath string = "/home/spies/Downloads/17730027.pdf"

func main() {
	ocr := gosseract.NewClient()
	defer ocr.Close()
	imagick.Initialize()
	defer imagick.Terminate()

	img, err := FileHandling.GetFileAsImg(imgPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	writeJpegToDisc(img, "../../Downloads/croppedImg.jpg")

	croppedImage := PreProcessing.CropImage(*img)

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
