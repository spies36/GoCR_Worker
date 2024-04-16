package main

import (
	"fmt"
	"time"

	"github.com/otiai10/gosseract/v2"
	"github.com/spies36/GoCR_Worker/FileHandling"
	"github.com/spies36/GoCR_Worker/PreProcessing"
	"gopkg.in/gographics/imagick.v2/imagick"
)

const imgPath string = "/home/spies/Downloads/17725436.pdf"

func main() {
	ocr := gosseract.NewClient()
	defer ocr.Close()
	imagick.Initialize()
	defer imagick.Terminate()
	preProcTimeStart := time.Now()
	img, err := FileHandling.GetFileAsImg(imgPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	croppedImage := PreProcessing.CropImage(*img)

	imgToProcess, err := FileHandling.PngImageToBytes(*croppedImage)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	preProcTimeEnd := time.Since(preProcTimeStart)
	ocrTimeStart := time.Now()
	ocr.SetImageFromBytes(imgToProcess)
	text, err := ocr.Text()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ocrTimeFinished := time.Since(ocrTimeStart)
	fmt.Printf("----- OCR Results ----- \n")
	fmt.Println(text)
	fmt.Printf("\n ----- Time results ----- \n")
	fmt.Printf("Preprocessing: %f \n", preProcTimeEnd.Seconds())
	fmt.Printf("OCR: %f \n", ocrTimeFinished.Seconds())
	fmt.Printf("Total time: %f \n", (ocrTimeFinished + preProcTimeEnd).Seconds())

}
