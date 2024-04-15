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
		img, err = decodePdf(file)
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

func decodePdf(file *os.File) (*image.Image, error) {
	//Convert pdf to jpg
	jpgFile, err := PreProcessing.Rasterize(file)
	if err != nil {
		return nil, err
	}
	defer jpgFile.Close()

	img, err := jpeg.Decode(jpgFile)
	if err != nil {
		return &img, err
	}
	err = os.Remove(jpgFile.Name())
	if err != nil {
		return &img, err
	}

	return &img, err
}
