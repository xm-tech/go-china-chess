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
	Circle       *Circle
}

func NewGame(windowWidth, windowHeight, screenWidth, screenHeight int, titile string, circle *Circle) *Game {
	game := &Game{
		WindowWidth:  windowWidth,
		WindowHeight: windowHeight,
		Titile:       titile,
		Circle:       circle,
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}
	return game
}

func (g *Game) Update() error {
	log.Println("Update exec")
	g.Circle.Run(300)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.Titile)
	ebitenutil.DrawCircle(screen, g.Circle.X, g.Circle.Y, g.Circle.R, g.Circle.Color)
	log.Println("Draw exec")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ScreenWidth, g.ScreenWidth
}
