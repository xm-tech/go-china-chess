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
		Y:     WindowHeight - GridSize + BoardEdgeWidth,
		Id:    RedJu0,
		Image: "res/RedJu.png",
		Camp:  0,
	}
	img, _, _ := ebitenutil.NewImageFromFile(self.Chesses[RedJu0].Image)
	self.Chesses[RedJu0].Img = img

	// 红马
	self.Chesses[RedMa0] = &Chessman{
		X:     self.Chesses[RedJu0].X + GridSize,
		Y:     WindowHeight - GridSize + BoardEdgeWidth,
		Id:    RedMa0,
		Image: "res/RedMa.png",
		Camp:  0,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedMa0].Image)
	self.Chesses[RedMa0].Img = img

	// 红相
	self.Chesses[RedXiang0] = &Chessman{
		X:     self.Chesses[RedMa0].X + GridSize,
		Y:     WindowHeight - GridSize + BoardEdgeWidth,
		Id:    RedXiang0,
		Image: "res/RedXiang.png",
		Camp:  0,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedXiang0].Image)
	self.Chesses[RedXiang0].Img = img

	// 红士
	self.Chesses[RedShi0] = &Chessman{
		X:     self.Chesses[RedXiang0].X + GridSize,
		Y:     WindowHeight - GridSize + BoardEdgeWidth,
		Id:    RedMa0,
		Image: "res/RedShi.png",
		Camp:  0,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedShi0].Image)
	self.Chesses[RedShi0].Img = img

	// 红帅
	self.Chesses[RedShuai] = &Chessman{
		X:     self.Chesses[RedShi0].X + GridSize,
		Y:     WindowHeight - GridSize + BoardEdgeWidth,
		Id:    RedShuai,
		Image: "res/RedShuai.png",
		Camp:  0,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedShuai].Image)
	self.Chesses[RedShuai].Img = img

	// 红士
	self.Chesses[RedShi1] = &Chessman{
		X:     self.Chesses[RedShuai].X + GridSize,
		Y:     WindowHeight - GridSize + BoardEdgeWidth,
		Id:    RedShi1,
		Image: "res/RedShi.png",
		Camp:  0,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedShi1].Image)
	self.Chesses[RedShi1].Img = img

	// 红相
	self.Chesses[RedXiang1] = &Chessman{
		X:     self.Chesses[RedShi1].X + GridSize,
		Y:     WindowHeight - GridSize + BoardEdgeWidth,
		Id:    RedXiang1,
		Image: "res/RedXiang.png",
		Camp:  0,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedXiang1].Image)
	self.Chesses[RedXiang1].Img = img

	// 红马
	self.Chesses[RedMa1] = &Chessman{
		X:     self.Chesses[RedXiang1].X + GridSize,
		Y:     WindowHeight - GridSize + BoardEdgeWidth,
		Id:    RedMa1,
		Image: "res/RedMa.png",
		Camp:  0,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedMa1].Image)
	self.Chesses[RedMa1].Img = img

	// 红车
	self.Chesses[RedJu1] = &Chessman{
		X:     self.Chesses[RedMa1].X + GridSize,
		Y:     WindowHeight - GridSize + BoardEdgeWidth,
		Id:    RedJu1,
		Image: "res/RedJu.png",
		Camp:  0,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedJu1].Image)
	self.Chesses[RedJu1].Img = img

	// 红炮0
	self.Chesses[RedPao0] = &Chessman{
		X:     self.Chesses[RedMa0].X,
		Y:     WindowHeight - 3*GridSize,
		Id:    RedPao0,
		Image: "res/RedPao.png",
		Camp:  0,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedPao0].Image)
	self.Chesses[RedPao0].Img = img

	// 红炮1
	self.Chesses[RedPao1] = &Chessman{
		X:     self.Chesses[RedMa1].X,
		Y:     WindowHeight - 3*GridSize,
		Id:    RedPao1,
		Image: "res/RedPao.png",
		Camp:  0,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedPao1].Image)
	self.Chesses[RedPao1].Img = img

	// 红兵
	self.Chesses[RedBing0] = &Chessman{
		X:     self.Chesses[RedJu0].X,
		Y:     WindowHeight - 4*GridSize,
		Id:    RedBing0,
		Image: "res/RedBing.png",
		Camp:  0,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedBing0].Image)
	self.Chesses[RedBing0].Img = img

	// 红兵
	self.Chesses[RedBing1] = &Chessman{
		X:     self.Chesses[RedXiang0].X,
		Y:     WindowHeight - 4*GridSize,
		Id:    RedBing1,
		Image: "res/RedBing.png",
		Camp:  0,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedBing1].Image)
	self.Chesses[RedBing1].Img = img

	// 红兵
	self.Chesses[RedBing2] = &Chessman{
		X:     self.Chesses[RedShuai].X,
		Y:     WindowHeight - 4*GridSize,
		Id:    RedBing2,
		Image: "res/RedBing.png",
		Camp:  0,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedBing2].Image)
	self.Chesses[RedBing2].Img = img

	// 红兵
	self.Chesses[RedBing3] = &Chessman{
		X:     self.Chesses[RedXiang1].X,
		Y:     WindowHeight - 4*GridSize,
		Id:    RedBing3,
		Image: "res/RedBing.png",
		Camp:  0,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedBing3].Image)
	self.Chesses[RedBing3].Img = img

	// 红兵
	self.Chesses[RedBing4] = &Chessman{
		X:     self.Chesses[RedJu1].X,
		Y:     WindowHeight - 4*GridSize,
		Id:    RedBing4,
		Image: "res/RedBing.png",
		Camp:  0,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[RedBing4].Image)
	self.Chesses[RedBing4].Img = img

	// 初始化黑棋

	// 黑ju
	self.Chesses[BlackJu0] = &Chessman{
		X:     0 + BoardEdgeWidth,
		Y:     BoardEdgeWidth,
		Id:    BlackJu0,
		Image: "res/BlackJu.png",
		Camp:  1,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[BlackJu0].Image)
	self.Chesses[BlackJu0].Img = img
	// 黑马
	self.Chesses[BlackMa0] = &Chessman{
		X:     self.Chesses[BlackJu0].X + GridSize,
		Y:     BoardEdgeWidth,
		Id:    BlackMa0,
		Image: "res/BlackMa.png",
		Camp:  1,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[BlackMa0].Image)
	self.Chesses[BlackMa0].Img = img
	// 黑相
	self.Chesses[BlackXiang0] = &Chessman{
		X:     self.Chesses[BlackMa0].X + GridSize,
		Y:     BoardEdgeWidth,
		Id:    BlackXiang0,
		Image: "res/BlackXiang.png",
		Camp:  1,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[BlackXiang0].Image)
	self.Chesses[BlackXiang0].Img = img
	// 黑士
	self.Chesses[BlackShi0] = &Chessman{
		X:     self.Chesses[BlackXiang0].X + GridSize,
		Y:     BoardEdgeWidth,
		Id:    BlackShi0,
		Image: "res/BlackShi.png",
		Camp:  1,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[BlackShi0].Image)
	self.Chesses[BlackShi0].Img = img
	// 黑将
	self.Chesses[BlackJiang] = &Chessman{
		X:     self.Chesses[BlackShi0].X + GridSize,
		Y:     BoardEdgeWidth,
		Id:    BlackJiang,
		Image: "res/BlackJiang.png",
		Camp:  1,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[BlackJiang].Image)
	self.Chesses[BlackJiang].Img = img
	// 黑士
	self.Chesses[BlackShi1] = &Chessman{
		X:     self.Chesses[BlackJiang].X + GridSize,
		Y:     BoardEdgeWidth,
		Id:    BlackShi1,
		Image: "res/BlackShi.png",
		Camp:  1,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[BlackShi1].Image)
	self.Chesses[BlackShi1].Img = img
	// 黑相
	self.Chesses[BlackXiang1] = &Chessman{
		X:     self.Chesses[BlackShi1].X + GridSize,
		Y:     BoardEdgeWidth,
		Id:    BlackXiang1,
		Image: "res/BlackXiang.png",
		Camp:  1,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[BlackXiang1].Image)
	self.Chesses[BlackXiang1].Img = img
	// 黑马
	self.Chesses[BlackMa1] = &Chessman{
		X:     self.Chesses[BlackXiang1].X + GridSize,
		Y:     BoardEdgeWidth,
		Id:    BlackMa1,
		Image: "res/BlackMa.png",
		Camp:  1,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[BlackMa1].Image)
	self.Chesses[BlackMa1].Img = img
	// 黑车
	self.Chesses[BlackJu1] = &Chessman{
		X:     self.Chesses[BlackMa1].X + GridSize,
		Y:     BoardEdgeWidth,
		Id:    BlackJu1,
		Image: "res/BlackJu.png",
		Camp:  1,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[BlackJu1].Image)
	self.Chesses[BlackJu1].Img = img
	// 黑炮
	self.Chesses[BlackPao0] = &Chessman{
		X:     self.Chesses[BlackJu0].X + GridSize,
		Y:     BoardEdgeWidth + GridSize*2,
		Id:    BlackPao0,
		Image: "res/BlackPao.png",
		Camp:  1,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[BlackPao0].Image)
	self.Chesses[BlackPao0].Img = img
	// 黑炮1
	self.Chesses[BlackPao1] = &Chessman{
		X:     self.Chesses[BlackMa1].X,
		Y:     BoardEdgeWidth + GridSize*2,
		Id:    BlackPao1,
		Image: "res/BlackPao.png",
		Camp:  1,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[BlackPao1].Image)
	self.Chesses[BlackPao1].Img = img
	// 黑兵1
	self.Chesses[BlackZu0] = &Chessman{
		X:     self.Chesses[BlackJu0].X,
		Y:     BoardEdgeWidth + GridSize*3,
		Id:    BlackZu0,
		Image: "res/BlackBing.png",
		Camp:  1,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[BlackZu0].Image)
	self.Chesses[BlackZu0].Img = img
	// 黑兵2
	self.Chesses[BlackZu1] = &Chessman{
		X:     self.Chesses[BlackXiang0].X,
		Y:     BoardEdgeWidth + GridSize*3,
		Id:    BlackZu1,
		Image: "res/BlackBing.png",
		Camp:  1,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[BlackZu1].Image)
	self.Chesses[BlackZu1].Img = img
	// 黑兵3
	self.Chesses[BlackZu2] = &Chessman{
		X:     self.Chesses[BlackJiang].X,
		Y:     BoardEdgeWidth + GridSize*3,
		Id:    BlackZu2,
		Image: "res/BlackBing.png",
		Camp:  1,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[BlackZu2].Image)
	self.Chesses[BlackZu2].Img = img

	// 黑兵4
	self.Chesses[BlackZu3] = &Chessman{
		X:     self.Chesses[BlackXiang1].X,
		Y:     BoardEdgeWidth + GridSize*3,
		Id:    BlackZu3,
		Image: "res/BlackBing.png",
		Camp:  1,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[BlackZu3].Image)
	self.Chesses[BlackZu3].Img = img
	// 黑兵5
	self.Chesses[BlackZu4] = &Chessman{
		X:     self.Chesses[BlackJu1].X,
		Y:     BoardEdgeWidth + GridSize*3,
		Id:    BlackZu4,
		Image: "res/BlackBing.png",
		Camp:  1,
	}
	img, _, _ = ebitenutil.NewImageFromFile(self.Chesses[BlackZu4].Image)
	self.Chesses[BlackZu4].Img = img

}
