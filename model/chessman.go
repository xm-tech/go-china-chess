package model

import "github.com/hajimehoshi/ebiten/v2"

// 棋子
type Chessman struct {
	X    int
	Y    int
	Id   int
	Name string
	// 图片名
	Image string
	// 渲染图片
	Img *ebiten.Image
	// 阵营 0:红， 1:黑
	Camp int
}
