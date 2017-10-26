package yac

import "image/color"

// IsWhite check what the current pixel is white.
func isWhite(c color.Color) bool {
	r, g, b, a := c.RGBA()
	rW, gW, bW, aW := color.White.RGBA()

	if r == rW &&
		g == gW &&
		b == bW &&
		a == aW {
		return true
	}

	return false
}
