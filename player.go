package main

import (
	"fmt"
	"image"
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
	op.GeoM.Translate(p.positionX+baseOffsetX, p.positionY+baseOffsetY)
	screen.DrawImage(p.image, op)
	msg := fmt.Sprintf("FPS: %0.2f, TPS: %0.2f, X: %0.2f, Y: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS(), p.positionX, p.positionY)
	ebitenutil.DebugPrint(screen, msg)
	msg = "\nW - go up, S - go down, A - go left, D - go right"
	ebitenutil.DebugPrint(screen, msg)
}

func DrawProjectiles(g *Game, screen *ebiten.Image) {
	for i, shot := range g.projectiles{
		if !shot.isActive {
			continue
		}
		op := &ebiten.DrawImageOptions{}
		if shot.isCollided {
			op.GeoM.Translate(shot.positionX - shot.explosion1SpriteWidth / 2, shot.positionY - shot.explosion1SpriteHeight / 2)
			sx, sy := shot.explosionFrame * int(shot.explosion1SpriteWidth), 0
			screen.DrawImage(shot.explosion1.SubImage(image.Rect(sx, sy, sx+int(shot.explosion1SpriteWidth), sy+int(shot.explosion1SpriteHeight))).(*ebiten.Image), op)
			if shot.explosionFrame == shot.explosionNumSprites - 1 {
				g.projectiles[i].isActive = false
				continue
			}
			if g.count % 20 == 0 {
				g.projectiles[i].explosionFrame += 1
			}
		} else {
			baseOffsetX := float64(shot.image.Bounds().Dx()) / 2
			baseOffsetY := float64(shot.image.Bounds().Dy()) / 2
			op.GeoM.Translate(-baseOffsetX, -baseOffsetY)
			op.GeoM.Rotate(shot.rotation)
			op.GeoM.Scale(0.4, 0.4)
			op.GeoM.Translate(shot.positionX, shot.positionY)
			screen.DrawImage(shot.image, op)
		}
		msg := fmt.Sprintf("\n\n\n\n\nprojectiles: %d", len(g.projectiles))
		ebitenutil.DebugPrint(screen, msg)
	}
}

func UpdatePlayer(g *Game) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if g.count - g.player.lastShot  >= g.player.reloadSpeed {
			addProjectile(g)
		}
	}
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
	if g.player.positionX <= minXCoordinate {
		g.player.positionX = minXCoordinate
	}
	if g.player.positionX >= maxXCoordinate - g.player.width - (g.player.height - g.player.width) / 2 {
		g.player.positionX = maxXCoordinate - g.player.width - (g.player.height - g.player.width) / 2
	}
	if g.player.positionY <= minYCoordinate {
		g.player.positionY = minYCoordinate
	}
	if g.player.positionY >= maxYCoordinate - g.player.height {
		g.player.positionY = maxYCoordinate - g.player.height
	}
}

func UpdateProjectiles(g *Game) {
	updatedProjectiles := []Projectile{}
	for i, shot := range g.projectiles{
		if !shot.isActive {
			continue
		}
		if !shot.isCollided {
			deltaY := math.Cos(shot.rotation) * shot.moveSpeed
			deltaX := -math.Sin(shot.rotation) * shot.moveSpeed
			g.projectiles[i].positionX -= deltaX
			g.projectiles[i].positionY -= deltaY
			if checkProjectileCollision(&shot) {
				g.projectiles[i].isCollided = true
			}
		}
		updatedProjectiles = append(updatedProjectiles, g.projectiles[i])
	}
	g.projectiles = updatedProjectiles
}

func addProjectile(g *Game) {
	// Calculate offset for spawning a projectile
	posOffsetX := float64(g.player.image.Bounds().Dx()) / 2
	posOffsetY := float64(g.player.image.Bounds().Dy()) / 2
	deltaX := posOffsetX * math.Abs(math.Cos(g.player.rotation)) + posOffsetX * math.Abs(math.Sin(g.player.rotation))
	deltaY := posOffsetY * math.Abs(math.Cos(g.player.rotation)) + posOffsetY * math.Abs(math.Sin(g.player.rotation))
	if g.player.rotation == 0 || g.player.rotation == math.Pi {
		deltaY -= posOffsetY * (1 / math.Cos(g.player.rotation))
	} else {
		deltaX += 2 * posOffsetX * math.Sin(g.player.rotation)
	}

	g.projectiles = append(g.projectiles, Projectile{
		width: ligthProjectileWidth,
		height: ligthProjectileHeight,
		rotation: g.player.rotation,
		moveSpeed: g.player.moveSpeed + 1,
		positionX: g.player.positionX+deltaX,
		positionY: g.player.positionY+deltaY,
		explosion1SpriteWidth: 50, explosion1SpriteHeight: 50,
		explosionNumSprites: 8,
		explosionFrame: 0,
		isCollided: false,
		isActive: true,
		image: g.resources.projectileImage,
		explosion1: g.resources.projectileExplImage,
	})
	g.player.lastShot = g.count
}

func checkProjectileCollision(p *Projectile) bool {
	if p.positionX <= minXCoordinate || p.positionY <= minYCoordinate || p.positionX >= maxXCoordinate - 25 || p.positionY >= maxYCoordinate - 25 {
		return  true
	}
	return false
}