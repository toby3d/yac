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

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			c := bounds.At(x, y)

			// Step 1: remove all white...
			if isWhite(c) {
				continue
			}

			// ...black...
			if isBlack(c) {
				continue
			}

			// ...and alpha pixels.
			// TODO: check alpha pixels

			r, g, b, a := c.RGBA()
			newColor := color.RGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				A: uint8(a),
			}

			log.Printf("Save color on: [%d:%d]", x, y)
			draw.Draw(newImg, img.Bounds(), &image.Uniform{newColor}, image.ZP, draw.Src)

			// draw.Image.Set(x, y, c)
			// log.Printf("[%d,%d]", x, y)

			// log.Println("R:", r, "G:", g, "B:", b, "A:", a)
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
