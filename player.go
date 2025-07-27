package main

import (
	"fmt"
	"image"
	//"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


func DrawPlayer(g *Game, screen *ebiten.Image) {
	DrawTank(&g.player, screen, g.count)
	msg := fmt.Sprintf("FPS: %0.2f, TPS: %0.2f, X: %0.2f, Y: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS(), g.player.posX, g.player.posY)
	ebitenutil.DebugPrint(screen, msg)
}

func DrawTank(p *Tank, screen *ebiten.Image, count int) {
	if !p.isActive {
		return
	}
	if p.isShot {
		op := &ebiten.DrawImageOptions{}
		offsetX, offsetY := p.getExplositionOffset()
		op.GeoM.Translate(offsetX, offsetY)
		sx, sy := p.explosionFrame * int(p.explosionImage.Bounds().Dy()), 0
		screen.DrawImage(p.explosionImage.SubImage(image.Rect(sx, sy, sx+int(p.explosionImage.Bounds().Dy()), sy+int(p.explosionImage.Bounds().Dy()))).(*ebiten.Image), op)
		if p.explosionFrame == p.explosionNumSprites - 1 {
			p.isActive = false
			return
		}
		if count % p.explosionSpeed == 0 {
			p.explosionFrame += 1
		}
		return
	}
	// Calculate base transformation
	baseOffsetX := float64(p.hullImage.Bounds().Dx()) / 2
	baseOffsetY := float64(p.hullImage.Bounds().Dy()) / 2
	finalOffsetX := p.posX+baseOffsetX
	finalOffsetY := p.posY+baseOffsetY

	// Draw tracks
	tracksOp := &ebiten.DrawImageOptions{}
	tracksOp.GeoM.Translate(-baseOffsetX, -baseOffsetY)
	tracksOp.GeoM.Rotate(p.rotation)
	tracksOp.GeoM.Translate(finalOffsetX, finalOffsetY)
	trackOffsetX, trackOffsetY := getTracksOffset(p, true)
	tracksOp.GeoM.Translate(trackOffsetX, trackOffsetY)
	sx, sy := 0, 0
	fx, fy := p.tracksImage.Bounds().Dx(), p.tracksImage.Bounds().Dy() / 2
	if p.isMoving {
		i := count % 2
		screen.DrawImage(
			p.tracksImage.SubImage(
				image.Rect(sx, sy + i*fy, fx, fy + i*fy)).(*ebiten.Image), tracksOp)
	} else {
		screen.DrawImage(p.tracksImage.SubImage(image.Rect(sx, sy, fx, fy)).(*ebiten.Image), tracksOp)
	}
	trackOffsetX, trackOffsetY = getTracksOffset(p, false)
	tracksOp.GeoM.Translate(trackOffsetX, trackOffsetY)
	if p.isMoving {
		i := count % 2
		screen.DrawImage(
			p.tracksImage.SubImage(
				image.Rect(sx, sy + i*fy, fx, fy + i*fy)).(*ebiten.Image), tracksOp)
	} else {
		screen.DrawImage(p.tracksImage.SubImage(image.Rect(sx, sy, fx, fy)).(*ebiten.Image), tracksOp)
	}

	// Draw hull
	hullOp := &ebiten.DrawImageOptions{}
	hullOp.GeoM.Translate(-baseOffsetX, -baseOffsetY)
	hullOp.GeoM.Rotate(p.rotation)
	hullOp.GeoM.Translate(finalOffsetX, finalOffsetY)
	screen.DrawImage(p.hullImage, hullOp)

	// Draw turret
	turretOp := &ebiten.DrawImageOptions{}
	var turretOffsetX, turretOffsetY float64
	turretOp.GeoM.Translate(-baseOffsetX, -baseOffsetY)
	turretOp.GeoM.Rotate(p.rotation)
	turretOp.GeoM.Translate(finalOffsetX, finalOffsetY)
	if count - p.lastShot < 5 {
		turretOffsetX, turretOffsetY = getTurretOffset(p, true)
	} else {
		turretOffsetX, turretOffsetY = getTurretOffset(p, false)
	}
	turretOp.GeoM.Translate(turretOffsetX, turretOffsetY)
	screen.DrawImage(p.turretImage, turretOp)


	msg := fmt.Sprintf("FPS: %0.2f, TPS: %0.2f, X: %0.2f, Y: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS(), p.posX, p.posY)
	ebitenutil.DebugPrint(screen, msg)
}


func UpdatePlayer(g *Game) {
	g.player.isMoving = false
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if g.count - g.player.lastShot  >= g.player.reloadSpeed {
			addProjectile(&g.player, g)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.player.prevPosX = g.player.posX
		g.player.prevPosY = g.player.posY
		g.player.prevRotation = g.player.rotation
		g.player.rotation = 0
		g.player.posY -= 1
		g.player.isMoving = true
		UpdateCollisions(&g.player, g)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.prevPosX = g.player.posX
		g.player.prevPosY = g.player.posY
		g.player.prevRotation = g.player.rotation
		g.player.rotation = math.Pi
		g.player.posY += 1
		g.player.isMoving = true
		UpdateCollisions(&g.player, g)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.prevPosX = g.player.posX
		g.player.prevRotation = g.player.rotation
		g.player.rotation = 3 * math.Pi / 2
		g.player.posX -= 1
		g.player.isMoving = true
		UpdateCollisions(&g.player, g)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.prevPosX = g.player.posX
		g.player.prevRotation = g.player.rotation
		g.player.rotation = math.Pi / 2
		g.player.posX += 1
		g.player.isMoving = true
		UpdateCollisions(&g.player, g)
	}
}