package yac

import "math"

// Find method find key pixels in Colors array of filtered (or not) pixels.
// Returned an non-sorted Colors array of 4 key colors which can already be used
// as a basis for coloring.
func (c Colors) Find() Colors {
	colors := make(Colors, 4)
	count := 0

	// Find 4 square segments lenght of array
	segment := len(c) / 4
	// Find key pixel in center of first segment line
	key := math.Sqrt(float64(segment)) / 2

	for i := range c {
		if i == (segment*count+int(key))-1 {
			colors[count] = c[i]
			count++
			if count >= len(colors) {
				break
			}
		}
	}

	return colors
}
