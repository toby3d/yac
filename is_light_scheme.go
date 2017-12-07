package yac

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

// IsLightScheme helper check first color.Color array element of sorted key
// colors on lightness. Returned true if current array contains colors for light
// scheme.
func IsLightScheme(c []color.Color) bool {
	_, _, L := colorful.MakeColor(c[0]).Hsl()
	return L >= 0.5
}
