package yac

import (
	"image/color"
	"math"
)

// getLightness convert RGBA color to HSL color and returned value of L.
func getLightness(color color.Color) float64 {
	r, g, b, _ := color.RGBA()

	Cmax := math.Max(math.Max(float64(r), float64(g)), float64(b))
	Cmin := math.Min(math.Min(float64(r), float64(g)), float64(b))
	L := (Cmax + Cmin) / 2

	return L
}
