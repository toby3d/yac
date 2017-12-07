package yac

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

// Fix helper check first, second and last colors and fixed his contrast by
// lightness for better visual result. Returned fixed array of color.Color.
func Fix(c []color.Color) []color.Color {
	isLightScheme := IsLightScheme(c)

	bgH, bgS, bgL := colorful.MakeColor(c[0]).Hsl()
	btnH, btnS, btnL := colorful.MakeColor(c[1]).Hsl()
	txtH, txtS, txtL := colorful.MakeColor(c[3]).Hsl()

	var bgVSbtn, bgVStxt float64
	if isLightScheme {
		bgVSbtn = bgL - btnL
		bgVStxt = bgL - txtL
	} else {
		bgVSbtn = btnL - bgL
		bgVStxt = txtL - bgL
	}

	if bgVSbtn < 0.2 {
		if isLightScheme {
			bgL += (0.2 - bgVSbtn) / 2
			btnL -= (0.2 - bgVSbtn) / 2
		} else {
			bgL -= (0.2 - bgVSbtn) / 2
			btnL += (0.2 - bgVSbtn) / 2
		}

		bgL = fixRange(bgL)
		btnL = fixRange(btnL)

		bgNewR, bgNewG, bgNewB := colorful.Hsl(bgH, bgS, bgL).RGB255()
		btnNewR, btnNewG, btnNewB := colorful.Hsl(btnH, btnS, btnL).RGB255()
		c[0] = color.RGBA{R: bgNewR, G: bgNewG, B: bgNewB}
		c[1] = color.RGBA{R: btnNewR, G: btnNewG, B: btnNewB}
	}

	if bgVStxt < 0.6 {
		if isLightScheme {
			bgL += (0.6 - bgVStxt) / 2
			txtL -= (0.6 - bgVStxt) / 2
		} else {
			bgL -= (0.6 - bgVStxt) / 2
			txtL += (0.6 - bgVStxt) / 2
		}

		bgL = fixRange(bgL)
		txtL = fixRange(txtL)

		bgNewR, bgNewG, bgNewB := colorful.Hsl(bgH, bgS, bgL).RGB255()
		txtNewR, txtNewG, txtNewB := colorful.Hsl(txtH, txtS, txtL).RGB255()

		c[0] = color.RGBA{R: bgNewR, G: bgNewG, B: bgNewB}
		c[3] = color.RGBA{R: txtNewR, G: txtNewG, B: txtNewB}
	}

	return c
}

func fixRange(i float64) float64 {
	if i <= 0 {
		i = 0
	}
	if i >= 1 {
		i = 1
	}
	return i
}
