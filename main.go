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
	startPosX = 134
	startPosY = 260
)


func (g *Game) Update() error {
	g.count++
	if !g.player.isShot {
		UpdateEnemies(g)
		UpdateProjectiles(g)
		UpdatePlayer(g)
		UpdateLevel(g)
	} else {
		UpdateEnemies(g)
		UpdateProjectiles(g)
		UpdateLevel(g)
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			resetGame(g)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !g.player.isActive {
		DrawLevel(g, screen)
		DrawEnemies(g, screen)
		DrawGameOverScreen(screen)
	} else {
		DrawLevel(g, screen)
		DrawEnemies(g, screen)
		DrawPlayer(g, screen)
		DrawProjectiles(g, screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func resetGame(g *Game) {
	g.player.isShot = false
	g.player.isActive = true
	g.player.posX = startPosX
	g.player.posY = startPosY
	g.player.prevPosX = startPosX
	g.player.prevPosY = startPosY
	g.player.rotation = 0
	g.player.explosionFrame = 0
	initLevel(g)
	g.projectiles = []Projectile{}
}

func main() {
	game := &Game{enemyKilledCount: 0}
	loadResources(game)
	initLevel(game)
	game.player = Tank{
		width: float64(game.resources.playerHullImage.Bounds().Dx()),
		height: float64(game.resources.playerHullImage.Bounds().Dy()),
		posX: startPosX, posY: startPosY,
		prevPosX: startPosX, prevPosY: startPosY,
		scale: 1,
		rotation: 0, moveSpeed: 1.2,
		reloadSpeed: 60, lastShot: 0,
		hullImage: game.resources.playerHullImage,
		turretImage: game.resources.playerTurretImage,
		tracksImage: game.resources.playerTracksImage,
		fireRollbackOffset: 2,
		isMoving: false, isShot: false, isActive: true,
		explosionNumSprites: 5,
		explosionFrame: 0,
		explosionSpeed: 3,
		explosionImage: game.resources.tankExplImage,}
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Hello, Gamer!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}