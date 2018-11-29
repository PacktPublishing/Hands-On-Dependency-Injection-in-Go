package wasted_effort

import (
	"image"
	"image/color"
	"math"
)

func d(r, v float64, i *image.RGBA, c color.Color) {
	for a := float64(0); a < 360; a++ {
		ra := math.Pi * 2 * a / 360
		x := r*math.Sin(ra) + v
		y := r*math.Cos(ra) + v
		i.Set(int(x), int(y), c)
	}
}
