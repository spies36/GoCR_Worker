package PreProcessing

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

type subImager interface {
	SubImage(r image.Rectangle) image.Image
}

func CropImage(imgPath string) (croppedImage image.Image, err error) {

	imgType, err := getImgType(imgPath)
	if err != nil {
		return
	}

	img, err := os.Open(imgPath)

	if err != nil {
		return
	}
	defer img.Close()

	var decodedImage image.Image

	switch imgType {
	case "jpg":
		decodedImage, err = decodeJpg(img)
	case "png":
		decodedImage, err = decodePng(img)
	}

	if err != nil {
		return
	}

	bounds := decodedImage.Bounds()
	width := bounds.Dx()

	cropSize := image.Rect(0, 0, width/2, width/2)
	croppedImage = decodedImage.(subImager).SubImage(cropSize)
	return
}

func getImgType(imgPath string) (extension string, err error) {
	lastIndex := strings.LastIndex(imgPath, ".")
	if lastIndex < 0 {
		err = errors.New("file has no type")
	} else {
		extension = imgPath[lastIndex+1:]
	}

	return
}

func decodeJpg(img *os.File) (origImg image.Image, err error) {
	origImg, err = jpeg.Decode(img)
	return
}

func decodePng(img *os.File) (origImg image.Image, err error) {
	origImg, err = png.Decode(img)
	return
}
