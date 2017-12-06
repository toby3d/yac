package yac

import "image/color"

// Analyze method grab your raw input image data and analyze them on key colors.
// This method do all basic steps: Open > Filter > Find and Sort. Returned an
// sorted Color array of 4 keys colors which can already be used as a basis for
// coloring.
func Analyze(data interface{}) ([]color.Color, error) {
	src, err := Open(data)
	if err != nil {
		return nil, err
	}

	pixels := Filter(src)
	pixels = Find(pixels)
	// pixels = pixels.Prune()
	pixels = Sort(pixels)
	pixels = Fix(pixels)

	return pixels, nil
}
