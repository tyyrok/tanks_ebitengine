package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth = 300
	ScreenHeight = 300
	WindowWidth = 1080
	WindowHeight = 720
	minXCoordinate = 20
	minYCoordinate = 20
	maxXCoordinate = 700
	maxYCoordinate = 390
)

type Tank struct {
	width, height float64
	positionX, positionY float64
	rotation, moveSpeed float64
	image *ebiten.Image
}

type Game struct{
	count int
	player Tank
	tanks []Tank
}

var background *ebiten.Image

func init() {
	var err error
	background, _, err = ebitenutil.NewImageFromFile("resources/back.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	g.count++
	UpdatePlayer(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, Gamer!")
	screen.DrawImage(background, nil)
	DrawPlayer(&g.player, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	game := &Game{}
	img, _, err := ebitenutil.NewImageFromFile("resources/tankAtiny.png")
	if err != nil {
		log.Fatal(err)
	}
	game.player = Tank{
		width: 30, height: 57,
		positionX: 100, positionY: 100,
		rotation: 0, moveSpeed: 1.2,
		image: img}
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Hello, Gamer!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}