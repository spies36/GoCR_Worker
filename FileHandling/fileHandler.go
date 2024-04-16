package FileHandling

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// Open file and return image.Image
func GetFileAsImg(path string) (*image.Image, error) {
	var img *image.Image = nil
	file, err := os.Open(path)
	if err != nil {
		return img, err
	}
	defer file.Close()

	img, err = DecodeFile(file, filepath.Ext(path))
	return img, err
}

func JpgImageToBytes(img image.Image) ([]byte, error) {
	var b bytes.Buffer
	err := jpeg.Encode(&b, img, nil)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func PngImageToBytes(img image.Image) ([]byte, error) {
	var b bytes.Buffer
	err := png.Encode(&b, img)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
