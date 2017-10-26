package yac

import "image/color"

// IsBlack check what the current pixel is black.
func isBlack(c color.Color) bool {
	r, g, b, a := c.RGBA()
	rB, gB, bB, aB := color.Black.RGBA()

	if r == rB &&
		g == gB &&
		b == bB &&
		a == aB {
		return true
	}

	return false
}
