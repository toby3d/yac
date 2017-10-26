package yac

import (
	"image"
	"image/jpeg"
	"os"
)

func save(src image.Image) error {
	dst, err := os.Create("yac.jpg")
	defer dst.Close()
	if err != nil {
		return err
	}

	return jpeg.Encode(dst, src, &jpeg.Options{Quality: 100})
}
