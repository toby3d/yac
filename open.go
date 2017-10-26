package yac

import (
	"bytes"
	"errors"
	"image"
	"io"
	"os"
)

var ErrUnsupportedType = errors.New("use string, byte array or io.Reader only")

// Open function analyze input data and try decode that as image.Image.
func Open(data interface{}) (image.Image, error) {
	switch src := data.(type) {
	case string: // File path
		file, err := os.Open(src)
		defer file.Close()
		if err != nil {
			return nil, err
		}

		img, _, err := image.Decode(file)
		return img, err
	case []byte: // Raw bytes
		img, _, err := image.Decode(bytes.NewReader(src))
		return img, err
	case io.Reader: // Reader
		img, _, err := image.Decode(src)
		return img, err
	default: // Unsupported
		return nil, UnsupportedType
	}
}
