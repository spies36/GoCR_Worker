package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
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
	writePngToDisc(img, "../../Downloads/croppedImg.png")

	croppedImage := PreProcessing.CropImage(*img)

	imgToProcess, err := pngImageToBytes(*croppedImage)
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

// func jpgImageToBytes(img image.Image) ([]byte, error) {
// 	var b bytes.Buffer
// 	err := jpeg.Encode(&b, img, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return b.Bytes(), nil
// }

// func writeJpegToDisc(img *image.Image, path string) {

// 	imgByteArr, err := jpgImageToBytes(*img)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	os.WriteFile(path, imgByteArr, 0777)
// }

func pngImageToBytes(img image.Image) ([]byte, error) {
	var b bytes.Buffer
	err := png.Encode(&b, img)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func writePngToDisc(img *image.Image, path string) {

	imgByteArr, err := pngImageToBytes(*img)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	os.WriteFile(path, imgByteArr, 0777)
}
