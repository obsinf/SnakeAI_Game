package main

import (
	"Snake/game"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	game.NewGame(50, 1200, 840).Start()
	pixelgl.Run(game.Run)
}
