package PreProcessing

import (
	"image"

	"github.com/ernyoke/imger/imgio"
	"github.com/ernyoke/imger/threshold"
)

func Binarize(imgPath string) (*image.Image, error) {

	var binaryImg image.Image

	greyImg, err := imgio.ImreadGray(imgPath)
	if err != nil {
		return &binaryImg, err
	}
	greyImg, err = SharpenGray(greyImg)

	if err != nil {
		return &binaryImg, err
	}

	binaryImg, err = threshold.OtsuThreshold(greyImg, threshold.Method(threshold.ThreshBinary))

	return &binaryImg, err
}
