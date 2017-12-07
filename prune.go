package yac

import "image/color"

// Prune helper check color.Color array on duplicated items. Returned
// color.Colors array with unique color.Color items in random order.
func Prune(c []color.Color) []color.Color {
	sorted := map[color.Color]bool{}

	// Create a map of all unique elements.
	for i := range c {
		sorted[c[i]] = true
	}

	// Place all keys from the map into a slice.
	var unique []color.Color
	for key := range sorted {
		unique = append(unique, key)
	}

	return unique
}
