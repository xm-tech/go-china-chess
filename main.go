package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xm-tech/go-china-chess/model"
)

func main() {
	aCircle := model.NewCircle(0, 100, 15, color.White)
	game := model.NewGame(640, 480, 320, 260, "ebiten game demo", aCircle)
	ebiten.SetWindowSize(game.WindowWidth, game.WindowHeight)
	ebiten.SetWindowTitle(game.Titile)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
