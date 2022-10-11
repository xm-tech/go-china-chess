package model

import (
	"fmt"
	"log"
	"math/rand"
	"time"

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
	self.Chesses = make(map[int]*Chessman)

	// 递增的棋子ID
	var idx int
	// 初始化我方棋子
	rand.Seed(time.Now().Unix())
	camp := rand.Intn(2)

	for _, v := range ChessBothA {
		chess := NewChessman(idx, v[1], v[2], camp, "", fmt.Sprintf("res/%v-%v.png", camp, v[0]))
		fmt.Println("+++", idx, chess.Image, v, chess.Pos)
		self.Chesses[idx] = chess
		PosChesses[chess.Pos] = chess.Id

		idx++
		if idx == 16 {
			break
		}
	}

	// 初始化对方棋子
	if camp == 0 {
		camp = 1
	} else {
		camp = 0
	}
	for _, v := range ChessBothB {
		chess := NewChessman(idx, v[1], v[2], camp, "", fmt.Sprintf("res/%v-%v.png", camp, v[0]))
		fmt.Println("---", idx, chess.Image, v, chess.Pos)
		self.Chesses[idx] = chess
		PosChesses[chess.Pos] = chess.Id

		idx++
		if idx == 32 {
			break
		}
	}
}
