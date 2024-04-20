package PreProcessing

import (
	"image"
)

type subImager interface {
	SubImage(r image.Rectangle) image.Image
}

// Crop image as rectangle
func CropImage(img image.Image) *image.Image {

	var croppedImage image.Image

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	cropSize := image.Rect(0, 50, width/2, height/2)
	croppedImage = img.(subImager).SubImage(cropSize)

	return &croppedImage
}
