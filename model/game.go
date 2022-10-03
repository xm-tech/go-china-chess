package model

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	WindowWidth  int
	WindowHeight int
	ScreenWidth  int
	ScreenHeight int
	Titile       string
	ChessBoard   *Chessboard
}

func NewGame(windowWidth, windowHeight, screenWidth, screenHeight int, titile string) *Game {
	game := &Game{
		WindowWidth:  windowWidth,
		WindowHeight: windowHeight,
		Titile:       titile,
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}

	chessBoard := NewChessBoard()
	game.ChessBoard = chessBoard
	return game
}

func (g *Game) Update() error {
	log.Println("Update exec")
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.Titile)
	img, _, err := ebitenutil.NewImageFromFile("res/ChessBoard.png")
	if err != nil {
		log.Println(err)
		return
	}
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(img, op)
	log.Println("Draw exec")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ScreenWidth, g.ScreenHeight
}

func (g *Game) Start() {
	ebiten.SetWindowSize(g.WindowWidth, g.WindowHeight)
	ebiten.SetWindowTitle(g.Titile)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
