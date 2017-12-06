package yac

import (
	"image/color"
	"math"
)

func find(filtered []color.Color) []color.Color {
	pixels := make([]color.Color, 4)
	count := 0

	// Find 4 square segments lenght of array
	segment := len(filtered) / 4
	// Find key pixel in center of first segment line
	key := math.Sqrt(float64(segment)) / 2

	for i := range filtered {
		if i == (segment*count+int(key))-1 {
			pixels[count] = filtered[i]
			count++
			if count >= len(pixels) {
				break
			}
		}
	}

	return pixels
}
