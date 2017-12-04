package yac

import (
	"bytes"
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
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

		img, format, err := image.Decode(file)
		log.Println("FORMAT:", format)
		return img, err
	case []byte: // Raw bytes
		img, format, err := image.Decode(bytes.NewReader(src))
		log.Println("FORMAT:", format)
		return img, err
	case io.Reader: // Reader
		img, format, err := image.Decode(src)
		log.Println("FORMAT:", format)
		return img, err
	default: // Unsupported
		return nil, ErrUnsupportedType
	}
}
