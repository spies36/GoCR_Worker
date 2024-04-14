package FileHandling

import (
	"image"
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
