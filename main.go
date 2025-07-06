package main

import (
	//"image"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth = 736
	ScreenHeight = 414
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
	rotation, moveSpeed, rotationSpeed float64
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

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.player.rotation = 0
		g.player.positionY -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.rotation = math.Pi
		g.player.positionY += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.rotation = 3 * math.Pi / 2
		g.player.positionX -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.rotation = math.Pi / 2
		g.player.positionX += 1
	}

	//if ebiten.IsKeyPressed(ebiten.KeyW) {
		//g.player.positionY -= 1
	//	deltaY := math.Cos(g.player.rotation) * g.player.moveSpeed
	//	deltaX := -math.Sin(g.player.rotation) * g.player.moveSpeed
	//	g.player.positionY -= deltaY
	//	g.player.positionX -= deltaX
	//}
	//if ebiten.IsKeyPressed(ebiten.KeyS) {
	//	deltaY := math.Cos(g.player.rotation) * g.player.moveSpeed
	//	deltaX := -math.Sin(g.player.rotation) * g.player.moveSpeed
	//	g.player.positionY += deltaY
	//	g.player.positionX += deltaX
	//}
	//if ebiten.IsKeyPressed(ebiten.KeyA) {
	//	g.player.rotation -= g.player.rotationSpeed
	//}
	//if ebiten.IsKeyPressed(ebiten.KeyD) {
	//	g.player.rotation += g.player.rotationSpeed
	//}
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
		rotationSpeed: 0.1, image: img}
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Hello, Gamer!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}