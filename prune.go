package yac

import "image/color"

// Prune helper check Colors array on duplicated items. Returned Colors array
// with unique color.Color items in random order.
func (c Colors) Prune() Colors {
	sorted := map[color.Color]bool{}

	// Create a map of all unique elements.
	for i := range c {
		sorted[c[i]] = true
	}

	// Place all keys from the map into a slice.
	unique := Colors{}
	for key := range sorted {
		unique = append(unique, key)
	}

	return unique
}
