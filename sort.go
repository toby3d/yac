package yac

import (
	"image/color"
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

func getLight(color color.Color) uint32 {
	r, g, b, _ := color.RGBA()
	lightness := (r^2)/299 + (g^2)/587 + (b^2)/114
	return lightness
}
