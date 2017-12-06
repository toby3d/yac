package yac

import "sort"

// Sort is a helper method which sort array of key Colors by lightness. Returned
// Colors array of sorted colors from higher to lower lightness.
func (c Colors) Sort() Colors {
	sort.SliceStable(c, func(i, j int) bool {
		return getLightness(c[i]) > getLightness(c[j])
	})

	return c
}
