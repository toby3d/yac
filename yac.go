









package yac

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

func GetColors(data interface{}) error {
	src, err := Open(data)
	if err != nil {
		return err
	}

	bounds := src.Bounds()
	size := bounds.Size()

	log.Println("MaxWidth:", size.X)
	log.Println("MawHeight:", size.Y)

	newImg := image.NewRGBA(bounds)

	pixelArray := make([][]color.Color, size.X)
	for i := range pixelArray {
		pixelArray[i] = make([]color.Color, size.Y)
	}
	pAx := 0
	pAy := 0

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			// TODO: check alpha pixels
			if isWhite(src.At(x, y)) ||
				isBlack(src.At(x, y)) {
				continue
			}

			r, g, b, a := src.At(x, y).RGBA()
			newColor := color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			}

			pixelArray[pAx][pAy] = newColor
			pAx++
			if pAx >= size.X {
				pAx = 0
				pAy++
			}

		}
	}
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {

			m := image.NewRGBA(image.Rect(x, y, x+1, y+1))

			newColor := pixelArray[x][y]
			if newColor == nil {
				draw.Draw(newImg, m.Bounds(), &image.Uniform{color.RGBA{0, 0, 255, 255}}, image.ZP, draw.Src)
			} else {
				draw.Draw(newImg, m.Bounds(), &image.Uniform{newColor}, image.ZP, draw.Src)
			}

		}
	}


	// Step 2: resize image to 2x2

	// Step 3: get colors, convert to HSL, sort by L

	// Step 4: done!
	// Light theme:
	// - Background: most lighter color;
	// - Button: cosest to light color;
	// - Text: most darker color;
	//
	// Dark theme:
	// - Background: most darker color;
	// - Button: cosest to dark color;
	// - Text: most lighter color;

	// Step 5: Correct some data.
	// If icon has only one color - make background more darker;
	// If icon have background - look up on pixels on borders: if them all equals - set same background color.

	// TODO: Don't write result to file - just read and sort pixel colors
	result, err := os.Create("yac.jpg")
	defer result.Close()
	if err != nil {
		return err
	}

	return jpeg.Encode(result, newImg, &jpeg.Options{Quality: 100})
}
