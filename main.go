package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth = 300
	ScreenHeight = 300
	WindowWidth = 480
	WindowHeight = 480
	minXCoordinate = 0
	minYCoordinate = 0
	maxXCoordinate = 300
	maxYCoordinate = 300
)


func (g *Game) Update() error {
	g.count++
	UpdateEnemies(g)
	UpdatePlayer(g)
	UpdateProjectiles(g)
	UpdateLevel(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawLevel(g, screen)
	DrawEnemies(g, screen)
	DrawPlayer(g, screen)
	DrawProjectiles(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	game := &Game{}
	loadResources(game)
	initLevel(game)
	game.player = Tank{
		width: float64(game.resources.playerHullImage.Bounds().Dx()),
		height: float64(game.resources.playerHullImage.Bounds().Dy()),
		posX: 134, posY: 260,
		prevPosX: 134, prevPosY: 260,
		scale: 1,
		rotation: 0, moveSpeed: 1.2,
		reloadSpeed: 60, lastShot: 0,
		hullImage: game.resources.playerHullImage,
		turretImage: game.resources.playerTurretImage,
		tracksImage: game.resources.playerTracksImage,
		fireRollbackOffset: 2,
		isMoving: false, isShot: false,}
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Hello, Gamer!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}