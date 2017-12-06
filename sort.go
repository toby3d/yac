package yac

import (
	"image/color"
	"math"
	gosort "sort"
)

func sort(finded []color.Color) []color.Color {
	gosort.SliceStable(finded, func(i, j int) bool {
		one := getLight(finded[i])
		two := getLight(finded[j])
		return one > two
	})

	return finded
}

func getLight(color color.Color) float64 {
	r, g, b, _ := color.RGBA()

	Cmax := math.Max(math.Max(float64(r), float64(g)), float64(b))
	Cmin := math.Min(math.Min(float64(r), float64(g)), float64(b))
	L := (Cmax + Cmin) / 2

	return L
}
