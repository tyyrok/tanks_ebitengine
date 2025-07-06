package main

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawPlayer(p *Tank, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	baseOffsetX := float64(p.image.Bounds().Dx()) / 2
	baseOffsetY := float64(p.image.Bounds().Dy()) / 2
	op.GeoM.Translate(-baseOffsetX, -baseOffsetY)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(p.positionX, p.positionY)
	screen.DrawImage(p.image, op)
	msg := fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS())
	ebitenutil.DebugPrint(screen, msg)
}

func UpdatePlayer(g *Game) {
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
	if g.player.positionX <= g.player.width {
		g.player.positionX = g.player.width
	}
	if g.player.positionX >= ScreenWidth - g.player.width {
		g.player.positionX = ScreenWidth - g.player.width
	}
	if g.player.positionY <= g.player.height / 2 {
		g.player.positionY = g.player.height / 2
	}
	if g.player.positionY >= ScreenHeight - g.player.height / 2 {
		g.player.positionY = ScreenHeight - g.player.height / 2
	}
}