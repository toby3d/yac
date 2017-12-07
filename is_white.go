package yac

import "image/color"

// isWhite helper check what the current pixel is white.
func isWhite(c color.Color) bool {
	r, g, b, _ := c.RGBA()
	R, G, B, _ := color.White.RGBA()

	return r == R && g == G && b == B
}
