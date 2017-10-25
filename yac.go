package yac

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)


// Check what current pixel is totally black
func isBlack(c color.Color) bool {
	r, g, b, a := c.RGBA()
	blackR, blackG, blackB, blackA := color.Black.RGBA()

	if r == blackR &&
		g == blackG &&
		b == blackB &&
		a == blackA {
		return true
	}

	return false
}

// Check what current pixel is totally white
func isWhite(c color.Color) bool {
	r, g, b, a := c.RGBA()
	whiteR, whiteG, whiteB, whiteA := color.White.RGBA()

	if r == whiteR &&
		g == whiteG &&
		b == whiteB &&
		a == whiteA {
		return true
	}

	return false
}

// TODO: Check what current pixel is totally transparant

func GetColors(fileName string) error {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return err
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	bounds := img.Bounds()
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
			c := img.At(x, y)

			// Step 1: remove all white and black
			if isWhite(c) || isBlack(c) {
				continue
			}

			// ...and alpha pixels.
			// TODO: check alpha pixels

			r, g, b, a := c.RGBA()
			newColor := color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			}

			pixelArray[pAx][pAy] = newColor;
			pAx++
			if pAx >= size.X {
				pAx = 0
				pAy++
			}

			//log.Printf("Save color on: [%d:%d]", x, y)
			//log.Printf("Save color: %d %d %d", uint8(r >> 8), uint8(g >> 8), uint8(b >> 8))

		}
	}
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {

			m := image.NewRGBA(image.Rect(x, y, x + 1, y + 1))

			newColor := pixelArray[x][y];
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

	// Write image to separate file for visual debug
	result, err := os.Create("yac.jpg")
	defer result.Close()
	if err != nil {
		return err
	}

	return jpeg.Encode(result, newImg, &jpeg.Options{Quality: 100})
}
