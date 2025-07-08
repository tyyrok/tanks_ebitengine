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
	minXCoordinate = 24
	minYCoordinate = 16
	maxXCoordinate = 295
	maxYCoordinate = 300
	ligthProjectileWidth = 12
	ligthProjectileHeight = 25
)


var background, projImg, explositionLight *ebiten.Image

func init() {
	var err error
	background, _, err = ebitenutil.NewImageFromFile("resources/back.png")
	if err != nil {
		log.Fatal(err)
	}
	projImg, _, err = ebitenutil.NewImageFromFile("resources/Light_Shell.png") 
	if err != nil {
		log.Fatal(err)
	}
	explositionLight, _, err = ebitenutil.NewImageFromFile("resources/explosion_1.png") 
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	g.count++
	UpdatePlayer(g)
	UpdateProjectiles(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, Gamer!")
	screen.DrawImage(background, nil)
	DrawPlayer(&g.player, screen)
	DrawProjectiles(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	game := &Game{}
	img, _, err := ebitenutil.NewImageFromFile("resources/tankAtiny2.png")
	if err != nil {
		log.Fatal(err)
	}
	game.player = Tank{
		width: 30, height: 57,
		positionX: 100, positionY: 100,
		rotation: 0, moveSpeed: 1.2,
		reloadSpeed: 60, lastShot: 0,
		image: img}
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Hello, Gamer!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}