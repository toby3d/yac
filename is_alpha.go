package yac

import "image/color"

// isAlpha check what the current pixel is any-transparant.
// TODO: I'm not sure that I need to filter "any-transparent" pixels...
func isAlpha(c color.Color) bool {
	_, _, _, a := c.RGBA()
	if uint8(a>>8) < 255 {
		return true
	}

	return false
}
