package PreProcessing

import (
	"math/rand"
	"os"
	"strconv"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func Rasterize(pdf *os.File) (*os.File, error) {
	magickWand := imagick.NewMagickWand()
	defer magickWand.Destroy()

	err := magickWand.ReadImageFile(pdf)
	if err != nil {
		return nil, err
	}

	err = magickWand.ResizeImage(500, 500, imagick.FILTER_LAGRANGE, 1)
	if err != nil {
		return nil, err
	}

	magickWand.SetIteratorIndex(0)
	err = magickWand.SetImageFormat("jpg")
	if err != nil {
		return nil, err
	}

	randNum := rand.Int()
	tempFileName := "temp" + strconv.Itoa(randNum) + ".jpg"
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
