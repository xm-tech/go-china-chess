package model

import "github.com/hajimehoshi/ebiten/v2"

// 棋子
type Chessman struct {
	X     int
	Y     int
	Id    int
	Name  string
	Image string
	Img   *ebiten.Image
}
