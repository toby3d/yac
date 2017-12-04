package yac

import (
	"image"
	"image/color"
)

func filter(src image.Image) []color.Color {
	bounds := src.Bounds()

	filtered := make([]color.Color, bounds.Max.X*bounds.Max.Y)
	count := 0
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := src.At(x, y)
			if isWhite(pixel) ||
				isBlack(pixel) ||
				isAlpha(pixel) {
				continue
			}

			filtered[count] = pixel
			count++
		}
	}

	filtered = filtered[:count-1]
	return filtered
}
