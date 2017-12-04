package yac

import "image/color"

func find(filtered []color.Color) []color.Color {
	pixels := make([]color.Color, 4)
	count := 0

	segment := len(filtered) / 8
	for i := range filtered {
		// I'm bad in maths :(
		if i == (segment-1) ||
			i == (segment*3-1) ||
			i == (segment*5-1) ||
			i == (segment*7-1) {
			pixels[count] = filtered[i]
			count++
			if count >= len(pixels) {
				break
			}
		}

	}

	return pixels
}
