package yac

import (
	"image/color"
	gosort "sort"
)

func sort(finded []color.Color) []color.Color {
	gosort.SliceStable(finded, func(i, j int) bool {
		return getLightness(finded[i]) > getLightness(finded[j])
	})

	return finded
}
