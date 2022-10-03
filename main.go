package main

import (
	"github.com/xm-tech/go-china-chess/model"
)

func main() {
	game := model.NewGame(520, 576, 520, 576, "ebiten game demo")
	game.Start()
}
