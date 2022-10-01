package main

import (
	"github.com/xm-tech/go-china-chess/model"
)

func main() {
	game := model.NewGame(640, 480, 320, 260, "ebiten game demo")
	game.Start()
}
