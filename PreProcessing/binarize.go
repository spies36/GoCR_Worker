package PreProcessing

import (
	"image"

	"github.com/ernyoke/imger/imgio"
	"github.com/ernyoke/imger/threshold"
)

func Binarize(imgPath string) (binaryImg image.Image, err error) {
	greyImg, err := imgio.ImreadGray(imgPath)
	if err != nil {
		return
	}

	binaryImg, err = threshold.OtsuThreshold(greyImg, threshold.Method(threshold.ThreshBinary))

	return
}
