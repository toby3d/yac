package yac

import "image/color"

func isAlpha(c color.Color) bool {
	_, _, _, a := c.RGBA()
	if uint8(a>>8) < 255 {
		return true
	}

	return false
}
