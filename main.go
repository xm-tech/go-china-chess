package main

import (
	"image/color"

	"github.com/xm-tech/go-china-chess/model"
)

func main() {
	circle := model.NewCircle(0, 100, 15, color.White)
	game := model.NewGame(640, 480, 320, 260, "ebiten game demo", circle)
	game.Start()
}
