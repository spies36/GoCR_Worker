package PreProcessing

import (
	"image"

	"github.com/ernyoke/imger/effects"
)

func SharpenGray(greyImg *image.Gray) (sharpGreyImg *image.Gray, err error) {
	sharpGreyImg, err = effects.SharpenGray(greyImg)
	return
}
