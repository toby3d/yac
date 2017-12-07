package yac

import (
	"image/color"
	"sort"

	"github.com/lucasb-eyer/go-colorful"
)

// Sort is a helper method which sort color.Color array of key colors by
// lightness. Returned color.Color array with sorted key colors: from the
// lightest to the darkest colors, if the current array is a light theme and
// vice versa, if it's dark.
func Sort(c []color.Color) []color.Color {
	var light, dark float64
	for i := range c {
		_, _, l := colorful.MakeColor(c[i]).Hsl()
		if l >= 0.5 {
			light++
		} else {
			dark++
		}
	}

	if light >= dark {
		sort.SliceStable(c, func(i, j int) bool {
			_, _, li := colorful.MakeColor(c[i]).Hsl()
			_, _, lj := colorful.MakeColor(c[j]).Hsl()
			return li > lj
		})
	} else {
		sort.SliceStable(c, func(i, j int) bool {
			_, _, li := colorful.MakeColor(c[i]).Hsl()
			_, _, lj := colorful.MakeColor(c[j]).Hsl()
			return li < lj
		})
	}

	return c
}
