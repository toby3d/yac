package yac

// Analyze method grab your raw input image data and analyze them on key colors.
// This method do all basic steps: Open > Filter > Find and Sort. Returned an
// sorted Color array of 4 keys colors which can already be used as a basis for
// coloring.
func Analyze(data interface{}) (Colors, error) {
	src, err := Open(data)
	if err != nil {
		return nil, err
	}

	pixels := Filter(src)
	pixels = pixels.Find()

	return pixels.Sort(), nil
}
