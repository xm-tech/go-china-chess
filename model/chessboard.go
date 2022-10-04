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
	chesses []*Chessman
}

func NewChessBoard() *Chessboard {
	chessBoard := &Chessboard{}
	img, _, err := ebitenutil.NewImageFromFile("res/ChessBoard.png")
	if err != nil {
		log.Println("NewChessBoard Err,err=", err)
	}
	chessBoard.bg = img
	return chessBoard
}

// 初始化棋盘上棋子
func (self *Chessboard) InitChess() {

}