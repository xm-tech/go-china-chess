package main

import (
	"github.com/xm-tech/go-china-chess/model"
)

func main() {
	game := model.NewGame(model.WindowWidth, model.WindowHeight, model.ScreenWidth, model.ScreenHeight, "china chess")
	game.Start()
}
