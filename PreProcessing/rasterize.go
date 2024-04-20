package PreProcessing

import (
	"math/rand"
	"os"
	"strconv"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func Rasterize(fileToRaster *os.File) (*os.File, error) {
	magickWand := imagick.NewMagickWand()
	defer magickWand.Destroy()

	magickWand.SetOption("density", "300")

	err := magickWand.ReadImageFile(fileToRaster)
	if err != nil {
		return nil, err
	}

	magickWand.SetIteratorIndex(0)

	magickWand.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_DEACTIVATE)

	magickWand.ResampleImage(180, 180, imagick.FILTER_MITCHELL, 1)

	err = magickWand.SetImageFormat("png")
	if err != nil {
		return nil, err
	}

	randNum := rand.Int()
	tempFileName := "temp" + strconv.Itoa(randNum) + ".png"
	file, err := os.Create(tempFileName)
	if err != nil {
		return nil, err
	}

	err = magickWand.WriteImageFile(file)
	if err != nil {
		return nil, err
	}
	return file, err
}
