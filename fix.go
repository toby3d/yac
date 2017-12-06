package yac

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

func Fix(c []color.Color) []color.Color {
	bgH, bgS, bgL := colorful.MakeColor(c[0]).Hsl()
	btnH, btnS, btnL := colorful.MakeColor(c[1]).Hsl()
	txtH, txtS, txtL := colorful.MakeColor(c[3]).Hsl()

	bgVSbtn := (btnL*100 - bgL*100)
	bgVStxt := (txtL*100 - bgL*100)

	if bgVSbtn < 20 {
		if bgL*100 > 50 {
			if bgL+((20-bgVSbtn)/2/100) < 1 {
				c[0] = colorful.MakeColor(colorful.Hsl(bgH, bgS, bgL+((20-bgVSbtn)/2/100)))
			} else {
				c[0] = colorful.MakeColor(colorful.Hsl(bgH, bgS, 1))
			}

			if btnL-((20-bgVSbtn)/2/100) > 0 {
				c[1] = colorful.MakeColor(colorful.Hsl(btnH, btnS, btnL-((20-bgVSbtn)/2/100)))
			} else {
				c[0] = colorful.MakeColor(colorful.Hsl(btnH, btnS, 0))
			}
		} else {
			if bgL-((20-bgVSbtn)/2/100) > 0 {
				c[0] = colorful.MakeColor(colorful.Hsl(bgH, bgS, bgL-((20-bgVSbtn)/2/100)))
			} else {
				c[0] = colorful.MakeColor(colorful.Hsl(bgH, bgS, 0))
			}

			if btnL+((20-bgVSbtn)/2/100) < 1 {
				c[1] = colorful.MakeColor(colorful.Hsl(btnH, btnS, btnL+((20-bgVSbtn)/2/100)))
			} else {
				c[0] = colorful.MakeColor(colorful.Hsl(btnH, btnS, 1))
			}
		}
	}

	if bgVStxt < 60 {
		if bgL*100 > 50 {
			if bgL+((60-bgVStxt)/2/100) < 1 {
				c[0] = colorful.MakeColor(colorful.Hsl(bgH, bgS, bgL+((60-bgVStxt)/2/100)))
			} else {
				c[0] = colorful.MakeColor(colorful.Hsl(bgH, bgS, 1))
			}

			if txtL-((60-bgVStxt)/2/100) > 0 {
				c[1] = colorful.MakeColor(colorful.Hsl(txtH, txtS, txtL-((60-bgVStxt)/2/100)))
			} else {
				c[0] = colorful.MakeColor(colorful.Hsl(txtH, txtS, 0))
			}
		} else {
			if bgL-((60-bgVStxt)/2/100) > 0 {
				c[0] = colorful.MakeColor(colorful.Hsl(bgH, bgS, bgL-((60-bgVStxt)/2/100)))
			} else {
				c[0] = colorful.MakeColor(colorful.Hsl(bgH, bgS, 0))
			}

			if txtL+((60-bgVStxt)/2/100) < 1 {
				c[1] = colorful.MakeColor(colorful.Hsl(txtH, txtS, txtL+((60-bgVStxt)/2/100)))
			} else {
				c[0] = colorful.MakeColor(colorful.Hsl(txtH, txtS, 1))
			}
		}
	}

	return c
}
