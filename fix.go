package yac

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

func Fix(c []color.Color) []color.Color {
	bgH, bgS, bgL := colorful.MakeColor(c[0]).Hsl()
	btnH, btnS, btnL := colorful.MakeColor(c[1]).Hsl()
	txtH, txtS, txtL := colorful.MakeColor(c[3]).Hsl()

	var bgVSbtn, bgVStxt float64
	if bgL >= 0.5 {
		bgVSbtn = bgL - btnL
		bgVStxt = bgL - txtL
	} else {
		bgVSbtn = btnL - bgL
		bgVStxt = txtL - bgL
	}

	if bgVSbtn < 0.2 {
		if bgL >= 0.5 {
			bgL += (0.2 - bgVSbtn) / 2
			btnL -= (0.2 - bgVSbtn) / 2
		} else {
			bgL -= (0.2 - bgVSbtn) / 2
			btnL += (0.2 - bgVSbtn) / 2
		}

		if bgL <= 0 {
			bgL = 0
		} else if bgL >= 1 {
			bgL = 1
		}

		if btnL <= 0 {
			btnL = 0
		} else if btnL >= 1 {
			btnL = 1
		}

		bgNewR, bgNewG, bgNewB := colorful.Hsl(bgH, bgS, bgL).RGB255()
		btnNewR, btnNewG, btnNewB := colorful.Hsl(btnH, btnS, btnL).RGB255()
		c[0] = color.RGBA{R: bgNewR, G: bgNewG, B: bgNewB}
		c[1] = color.RGBA{R: btnNewR, G: btnNewG, B: btnNewB}
	}

	if bgVStxt < 0.6 {
		if bgL >= 0.5 {
			bgL += (0.6 - bgVStxt) / 2
			txtL -= (0.6 - bgVStxt) / 2
		} else {
			bgL -= (0.6 - bgVStxt) / 2
			txtL += (0.6 - bgVStxt) / 2
		}

		if bgL <= 0 {
			bgL = 0
		} else if bgL >= 1 {
			bgL = 1
		}

		if txtL <= 0 {
			txtL = 0
		} else if txtL >= 1 {
			txtL = 1
		}

		bgNewR, bgNewG, bgNewB := colorful.Hsl(bgH, bgS, bgL).RGB255()
		txtNewR, txtNewG, txtNewB := colorful.Hsl(txtH, txtS, txtL).RGB255()

		c[0] = color.RGBA{R: bgNewR, G: bgNewG, B: bgNewB}
		c[3] = color.RGBA{R: txtNewR, G: txtNewG, B: txtNewB}
	}

	return c
}
