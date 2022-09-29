package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	circle *Circle
}

type Circle struct {
	x float64
	y float64
}

func (circle *Circle) Run() {
	circle.x += 1
}

func (g *Game) Update() error {
	log.Println("Update exec")
	g.circle.Run()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello Ebiten !")
	ebitenutil.DrawCircle(screen, g.circle.x, g.circle.y, 30, color.White)
	log.Println("Draw exec")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("hello ebiten")
	game := &Game{&Circle{x: 0, y: 100}}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
