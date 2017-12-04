package yac

import "image/color"

func isAlpha(c color.Color) bool {
	_, _, _, a := c.RGBA()
	if uint16(a) == color.Transparent.A {
		return true
	}

	return false
}
