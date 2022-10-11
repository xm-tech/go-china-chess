package model

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// 棋子
type Chessman struct {
	Id int
	X  int
	Y  int
	// 阵营 0:红， 1:黑
	Camp int
	// 棋子是否活着
	Alive bool
	Name  string
	// 图片名
	Image string
	// 渲染图片
	Img *ebiten.Image
}

func NewChessman(id, x, y, camp int, name string, image string) *Chessman {
	chess := &Chessman{
		Id:    id,
		X:     x,
		Y:     y,
		Camp:  camp,
		Alive: true,
		Name:  name,
		Image: image,
	}

	img, _, _ := ebitenutil.NewImageFromFile(chess.Image)
	chess.Img = img
	return chess
}

func (self *Chessman) dead() {
	self.Alive = false
}

func (self *Chessman) move(x, y int) {
	self.X = x
	self.Y = y
}
