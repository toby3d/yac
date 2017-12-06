package yac

import "image"

// Filter method filter raw image pixels from white, black and any-transparent
// pixels. Returned Colors array of filtered pixels.
func Filter(src image.Image) Colors {
	bounds := src.Bounds()

	pixels := make(Colors, bounds.Max.X*bounds.Max.Y)
	count := 0
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := src.At(x, y)
			if isWhite(pixel) ||
				isBlack(pixel) ||
				isAlpha(pixel) {
				continue
			}

			pixels[count] = pixel
			count++
		}
	}

	pixels = pixels[:count-1]
	return pixels
}
