package PreProcessing

import (
	"image"

	"github.com/ernyoke/imger/effects"
)

// Sharpen gray image
func SharpenGray(greyImg *image.Gray) (sharpGreyImg *image.Gray, err error) {
	sharpGreyImg, err = effects.SharpenGray(greyImg)
	return
}
