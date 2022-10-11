package model

import "github.com/hajimehoshi/ebiten/v2"

// 棋子
type Chessman struct {
	Id int
	X  int
	Y  int
	// 阵营 0:红， 1:黑
	Camp  int
	Alive bool
	Name  string
	// 图片名
	Image string
	// 渲染图片
	Img *ebiten.Image
}

func NewChessman(id, x, y, camp int, alive bool, name string, image string) *Chessman {
	chess := &Chessman{
		Id:    id,
		X:     x,
		Y:     y,
		Camp:  camp,
		Alive: alive,
		Name:  name,
		Image: image,
	}
	return chess
}

func (self *Chessman) move(x, y int) {
	self.X = x
	self.Y = y
}
