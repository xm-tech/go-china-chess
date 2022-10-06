package model

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	WindowWidth      int
	WindowHeight     int
	ScreenWidth      int
	ScreenHeight     int
	Titile           string
	ChessBoard       *Chessboard
	DrawImageOptions *ebiten.DrawImageOptions
}

func NewGame(windowWidth, windowHeight, screenWidth, screenHeight int, titile string) *Game {
	game := &Game{
		WindowWidth:      windowWidth,
		WindowHeight:     windowHeight,
		Titile:           titile,
		ScreenWidth:      screenWidth,
		ScreenHeight:     screenHeight,
		DrawImageOptions: &ebiten.DrawImageOptions{},
	}

	chessBoard := NewChessBoard()
	game.ChessBoard = chessBoard
	return game
}

func (g *Game) Update() error {
	// log.Println("Update exec")
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.Titile)
	// 画棋盘
	g.drawChessBoard(screen)
	// log.Println("Draw exec")
}

func (g *Game) drawChessBoard(screen *ebiten.Image) {
	// 画棋盘
	screen.DrawImage(g.ChessBoard.bg, g.DrawImageOptions)

	// 画棋子
	for _, chess := range g.ChessBoard.Chesses {
		log.Printf("drawChess,chessId:%v,X:%v,Y:%v", chess.Id, chess.X, chess.Y)
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(chess.X), float64(chess.Y))
		screen.DrawImage(chess.Img, options)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.WindowWidth, g.WindowHeight
}

func (g *Game) Start() {
	ebiten.SetWindowSize(g.WindowWidth, g.WindowHeight)
	ebiten.SetWindowTitle(g.Titile)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
