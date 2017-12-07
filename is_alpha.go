package yac

import "image/color"

// isAlpha check what the current pixel is any-transparant.
// TODO: I'm not sure that I need to filter "any-transparent" pixels...
func isAlpha(c color.Color) bool {
	_, _, _, a := c.RGBA()
	return a < 65535
}
