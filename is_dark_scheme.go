package yac

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

// IsDarkScheme helper check first color.Color array element of sorted key
// colors on lightness. Returned true if current array contains colors for dark
// scheme.
func IsDarkScheme(c []color.Color) bool {
	_, _, L := colorful.MakeColor(c[0]).Hsl()
	return L < 0.5
}
