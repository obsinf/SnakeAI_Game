package game

import (
	"github.com/faiface/pixel"
	"image"
	"image/color"
)

type Color struct {
	R, G, B uint8
}

func GenerateRectangle(width, height int, col Color) *pixel.PictureData {
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{width, height},
	})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{R: col.R, G: col.G, B: col.B, A: 0xff})
		}
	}

	return pixel.PictureDataFromImage(img)
}
