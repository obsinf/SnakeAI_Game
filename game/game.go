package game

import (
	"Snake/Err"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Game struct {
	Field
	CellSize int
	Width    int
	Height   int
}

func NewGame(ccize, width, height int) Game {
	f := make(Field, width)
	for i := 0; i < width; i++ {
		f[i] = make([]byte, height)
	}
	return Game{
		Field:    f,
		CellSize: ccize,
		Width:    width,
		Height:   height,
	}
}

var (
	ThisGame Game
)

func (g Game) Start() {
	ThisGame = g
}

func Run() {
	cfg := pixelgl.WindowConfig{
		Title: "Snake",
		Bounds: pixel.R(
			0, 0,
			float64(ThisGame.Width),
			float64(ThisGame.Height)),
		VSync: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	Err.Handle(err, "Error creating window")

	win.Clear(colornames.Aqua)

	block := GenerateRectangle(ThisGame.CellSize, ThisGame.CellSize, Color{
		R: 255,
		G: 0,
		B: 0,
	})
	sprite := pixel.NewSprite(block, block.Bounds())

	x, y := 0, 0

	for !win.Closed() {

		if win.Pressed(pixelgl.KeyS) {
			y -= 10
		}
		if win.Pressed(pixelgl.KeyW) {
			y += 10
		}

		if win.Pressed(pixelgl.KeyA) {
			x -= 10
		}
		if win.Pressed(pixelgl.KeyD) {
			x += 10
		}

		win.Clear(colornames.Aqua)

		sprite.Draw(win, pixel.IM.Moved(pixel.V(float64(x), float64(y))))
		win.Update()
	}
}
