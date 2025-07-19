package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth = 300
	ScreenHeight = 300
	WindowWidth = 1080
	WindowHeight = 720
	minXCoordinate = 0
	minYCoordinate = 0
	maxXCoordinate = 300
	maxYCoordinate = 300
	ligthProjectileWidth = 12
	ligthProjectileHeight = 25
)


func (g *Game) Update() error {
	g.count++
	UpdatePlayer(g)
	UpdateProjectiles(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawLevel(g, screen)
	DrawPlayer(&g.player, screen, g.count)
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
		width: 24, height: 38,
		posX: 50, posY: 100,
		prevPosX: 50, prevPosY: 100,
		rotation: 0, moveSpeed: 1.2,
		reloadSpeed: 60, lastShot: 0,
		hullImage: game.resources.playerHullImage,
		turretImage: game.resources.playerTurretImage,
		tracksImage: game.resources.playerTracksImage,
		fireRollbackOffset: 2,
		isMoving: false,}
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Hello, Gamer!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}