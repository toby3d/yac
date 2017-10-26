package yac

import (
	"image"
	"image/color"
	"image/draw"
)

func filter(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	buffer := make([][]color.Color, bounds.Max.X)
	for i := range buffer {
		buffer[i] = make([]color.Color, bounds.Max.Y)
	}
	bX := 0
	bY := 0

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			// TODO: check alpha pixels
			if isWhite(src.At(x, y)) ||
				isBlack(src.At(x, y)) {
				continue
			}

			r, g, b, a := src.At(x, y).RGBA()
			c := color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			}

			buffer[bX][bY] = c
			bX++
			if bX >= bounds.Max.X {
				bX = 0
				bY++
			}
		}
	}

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := image.NewRGBA(image.Rect(x, y, x+1, y+1))

			c := buffer[x][y]
			if c == nil {
				draw.Draw(dst, pixel.Bounds(), &image.Uniform{color.RGBA{0, 0, 255, 255}}, image.ZP, draw.Src)
			} else {
				draw.Draw(dst, pixel.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
			}

		}
	}

	return dst
}
