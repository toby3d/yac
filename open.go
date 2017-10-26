package yac

import (
	"bytes"
	"errors"
	"image"
	"io"
	"os"
)

var UnsupportedType = errors.New("use string, byte array or io.Reader only")

// Open function analyze input data and try decode that as image.Image.
func Open(data interface{}) (image.Image, error) {
	switch src := data.(type) {
	case string:
		file, err := os.Open(src)
		defer file.Close()
		if err != nil {
			return nil, err
		}

		img, _, err := image.Decode(file)
		return img, err
	case []byte:
		img, _, err := image.Decode(bytes.NewReader(src))
		return img, err
	case io.Reader:
		img, _, err := image.Decode(src)
		return img, err
	default:
		return nil, UnsupportedType
	}
}
