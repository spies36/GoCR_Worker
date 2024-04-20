package FileHandling

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/spies36/GoCR_Worker/PreProcessing"
)

// Convert file to image.Image and return a pointer
func DecodeFile(file *os.File, ext string) (*image.Image, error) {
	var img *image.Image
	var err error

	switch ext {
	case ".jpg":
		img, err = decodeJpg(file)
	case ".png":
		img, err = decodePng(file)
	case ".pdf":
		img, err = decodePdfOrTif(file)
	case ".tif":
		img, err = decodePdfOrTif(file)
	default:
		err = errors.New("Unsupported file type: " + ext)
		return img, err
	}

	return img, err
}

// Decode jpg to image.Image
func decodeJpg(file *os.File) (*image.Image, error) {
	img, err := jpeg.Decode(file)
	return &img, err
}

// Decode pnh to image.Image
func decodePng(file *os.File) (*image.Image, error) {
	img, err := png.Decode(file)
	return &img, err
}

func decodePdfOrTif(file *os.File) (*image.Image, error) {
	//Convert pdf to jpg
	pngFile, err := PreProcessing.Rasterize(file)
	if err != nil {
		return nil, err
	}
	defer pngFile.Close()

	img, err := png.Decode(pngFile)
	if err != nil {
		return &img, err
	}
	err = os.Remove(pngFile.Name())
	if err != nil {
		return &img, err
	}

	return &img, err
}
