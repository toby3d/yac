package yac

import "image/color"

// isBlack check what the current pixel is black.
func isBlack(c color.Color) bool {
	r, g, b, _ := c.RGBA()
	R, G, B, _ := color.Black.RGBA()

	return r == R && g == G && b == B
}
