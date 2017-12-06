package yac

import "sort"

// Sort is a helper method which sort array of key Colors by lightness. Returned
// Colors array with sorted key colors: from the lightest to the darkest colors,
// if the current array is a light theme and vice versa, if it's dark.
func (c Colors) Sort() Colors {
	var light, dark float64
	for i := range c {
		shade := getLightness(c[i])
		if getLightness(c[i]) >= 50 {
			light += shade
		} else {
			dark += shade
		}
	}

	if light >= dark {
		sort.SliceStable(c, func(i, j int) bool {
			return getLightness(c[i]) > getLightness(c[j])
		})
	} else {
		sort.SliceStable(c, func(i, j int) bool {
			return getLightness(c[i]) < getLightness(c[j])
		})
	}

	return c
}
