package yac

import (
	"image/color"
	"sort"
)

func Sort(finded []color.Color) []color.Color {
	sort.SliceStable(finded, func(i, j int) bool {
		return getLightness(finded[i]) > getLightness(finded[j])
	})

	return finded
}
