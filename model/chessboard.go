package model

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// 棋盘
type Chessboard struct {
	// 棋盘背景图
	bg      *ebiten.Image
	Chesses map[int]*Chessman
}

func NewChessBoard() *Chessboard {
	chessBoard := &Chessboard{}
	img, _, err := ebitenutil.NewImageFromFile("res/ChessBoard.png")
	if err != nil {
		log.Println("NewChessBoard Err,err=", err)
	}
	chessBoard.bg = img
	chessBoard.InitChess()
	return chessBoard
}

// 初始化棋盘上棋子
func (self *Chessboard) InitChess() {
	log.Println("InitChess")
	// 初始化红旗
	self.Chesses = make(map[int]*Chessman)

	// 红ju
	self.Chesses[RedJu0] = &Chessman{
		X:     0 + BoardEdgeWidth,
		Y:     WindowHeight - GridSize,
		Id:    RedJu0,
		Image: "res/RedJu.png",
	}
	img, _, _ := ebitenutil.NewImageFromFile(self.Chesses[RedJu0].Image)
	self.Chesses[RedJu0].Img = img

	// 红马
	self.Chesses[RedMa0] = &Chessman{
		X:     self.Chesses[RedJu0].X + GridSize,
		Y:     WindowHeight - GridSize,
		Id:    RedMa0,
		Image: "res/RedMa.png",
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedMa0].Image)
	self.Chesses[RedMa0].Img = img

	// 红相
	self.Chesses[RedXiang0] = &Chessman{
		X:     self.Chesses[RedMa0].X + GridSize,
		Y:     WindowHeight - GridSize,
		Id:    RedXiang0,
		Image: "res/RedXiang.png",
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedXiang0].Image)
	self.Chesses[RedXiang0].Img = img

	// 红士
	self.Chesses[RedShi0] = &Chessman{
		X:     self.Chesses[RedXiang0].X + GridSize,
		Y:     WindowHeight - GridSize,
		Id:    RedMa0,
		Image: "res/RedShi.png",
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedShi0].Image)
	self.Chesses[RedShi0].Img = img

	// 红帅
	self.Chesses[RedShuai] = &Chessman{
		X:     self.Chesses[RedShi0].X + GridSize,
		Y:     WindowHeight - GridSize,
		Id:    RedShuai,
		Image: "res/RedShuai.png",
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedShuai].Image)
	self.Chesses[RedShuai].Img = img

	// 红士
	self.Chesses[RedShi1] = &Chessman{
		X:     self.Chesses[RedShuai].X + GridSize,
		Y:     WindowHeight - GridSize,
		Id:    RedShi1,
		Image: "res/RedShi.png",
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedShi1].Image)
	self.Chesses[RedShi1].Img = img

	// 红相
	self.Chesses[RedXiang1] = &Chessman{
		X:     self.Chesses[RedShi1].X + GridSize,
		Y:     WindowHeight - GridSize,
		Id:    RedXiang1,
		Image: "res/RedXiang.png",
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedXiang1].Image)
	self.Chesses[RedXiang1].Img = img

	// 红马
	self.Chesses[RedMa1] = &Chessman{
		X:     self.Chesses[RedXiang1].X + GridSize,
		Y:     WindowHeight - GridSize,
		Id:    RedMa1,
		Image: "res/RedMa.png",
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedMa1].Image)
	self.Chesses[RedMa1].Img = img

	// 红车
	self.Chesses[RedJu1] = &Chessman{
		X:     self.Chesses[RedMa1].X + GridSize,
		Y:     WindowHeight - GridSize,
		Id:    RedJu1,
		Image: "res/RedJu.png",
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedJu1].Image)
	self.Chesses[RedJu1].Img = img

	// 红炮0
	self.Chesses[RedPao0] = &Chessman{
		X:     self.Chesses[RedMa0].X,
		Y:     WindowHeight - 3*GridSize,
		Id:    RedPao0,
		Image: "res/RedPao.png",
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedPao0].Image)
	self.Chesses[RedPao0].Img = img

	// 红炮1
	self.Chesses[RedPao1] = &Chessman{
		X:     self.Chesses[RedMa1].X,
		Y:     WindowHeight - 3*GridSize,
		Id:    RedPao1,
		Image: "res/RedPao.png",
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedPao1].Image)
	self.Chesses[RedPao1].Img = img

	// 红兵
	self.Chesses[RedBing0] = &Chessman{
		X:     self.Chesses[RedJu0].X,
		Y:     WindowHeight - 4*GridSize,
		Id:    RedBing0,
		Image: "res/RedBing.png",
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedBing0].Image)
	self.Chesses[RedBing0].Img = img

	// 红兵
	self.Chesses[RedBing1] = &Chessman{
		X:     self.Chesses[RedXiang0].X,
		Y:     WindowHeight - 4*GridSize,
		Id:    RedBing1,
		Image: "res/RedBing.png",
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedBing1].Image)
	self.Chesses[RedBing1].Img = img

	// 红兵
	self.Chesses[RedBing2] = &Chessman{
		X:     self.Chesses[RedShuai].X,
		Y:     WindowHeight - 4*GridSize,
		Id:    RedBing2,
		Image: "res/RedBing.png",
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedBing2].Image)
	self.Chesses[RedBing2].Img = img

	// 红兵
	self.Chesses[RedBing3] = &Chessman{
		X:     self.Chesses[RedXiang1].X,
		Y:     WindowHeight - 4*GridSize,
		Id:    RedBing3,
		Image: "res/RedBing.png",
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedBing3].Image)
	self.Chesses[RedBing3].Img = img

	// 红兵
	self.Chesses[RedBing4] = &Chessman{
		X:     self.Chesses[RedJu1].X,
		Y:     WindowHeight - 4*GridSize,
		Id:    RedBing4,
		Image: "res/RedBing.png",
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedBing4].Image)
	self.Chesses[RedBing4].Img = img

	// 初始化黑棋
}
