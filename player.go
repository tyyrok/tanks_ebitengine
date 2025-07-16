package main

import (
	"fmt"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


func DrawPlayer(p *Tank, screen *ebiten.Image, count int) {
	op := &ebiten.DrawImageOptions{}
	var turretOffsetX, turretOffsetY float64
	// Draw player hull
	baseOffsetX := float64(p.hullImage.Bounds().Dx()) / 2
	baseOffsetY := float64(p.hullImage.Bounds().Dy()) / 2
	op.GeoM.Translate(-baseOffsetX, -baseOffsetY)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(p.posX+baseOffsetX, p.posY+baseOffsetY)
	screen.DrawImage(p.hullImage, op)
	// Draw player turret
	if count - p.lastShot < 5 {
		turretOffsetX, turretOffsetY = getTurretOffset(p, true)
	} else {
		turretOffsetX, turretOffsetY = getTurretOffset(p, false)
	}

	op.GeoM.Translate(turretOffsetX, turretOffsetY)
	screen.DrawImage(p.turretImage, op)

	msg := fmt.Sprintf("FPS: %0.2f, TPS: %0.2f, X: %0.2f, Y: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS(), p.posX, p.posY)
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
			op.GeoM.Translate(shot.posX - shot.explosion1SpriteWidth / 2, shot.posY - shot.explosion1SpriteHeight / 2)
			sx, sy := shot.explosionFrame * int(shot.explosion1SpriteWidth), 0
			screen.DrawImage(shot.explosion1.SubImage(image.Rect(sx, sy, sx+int(shot.explosion1SpriteWidth), sy+int(shot.explosion1SpriteHeight))).(*ebiten.Image), op)
			if shot.explosionFrame == shot.explosionNumSprites - 1 {
				g.projectiles[i].isActive = false
				continue
			}
			if g.count % shot.explosionSpeed == 0 {
				g.projectiles[i].explosionFrame += 1
			}
		} else {
			baseOffsetX := float64(shot.image.Bounds().Dx()) / 2
			baseOffsetY := float64(shot.image.Bounds().Dy()) / 2
			op.GeoM.Translate(-baseOffsetX, -baseOffsetY)
			op.GeoM.Rotate(shot.rotation)
			op.GeoM.Scale(0.2, 0.2)
			op.GeoM.Translate(shot.posX, shot.posY)
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
		g.player.prevPosX = g.player.posX
		g.player.prevPosY = g.player.posY
		g.player.prevRotation = g.player.rotation
		g.player.rotation = 0
		g.player.posY -= 1
		UpdateCollisions(g)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.prevPosX = g.player.posX
		g.player.prevPosY = g.player.posY
		g.player.prevRotation = g.player.rotation
		g.player.rotation = math.Pi
		g.player.posY += 1
		UpdateCollisions(g)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.prevPosX = g.player.posX
		g.player.prevPosY = g.player.posY
		g.player.prevRotation = g.player.rotation
		g.player.rotation = 3 * math.Pi / 2
		g.player.posX -= 1
		UpdateCollisions(g)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.prevPosX = g.player.posX
		g.player.prevPosY = g.player.posY
		g.player.prevRotation = g.player.rotation
		g.player.rotation = math.Pi / 2
		g.player.posX += 1
		UpdateCollisions(g)
	}
	tRotatedX, tRotatedY, tWidth, tHeight  := getRotatedCoords(&g.player)
	//log.Printf("tX: %0f, tY: %0f, tW: %0f, tH: %0f", tRotatedX, tRotatedY, tWidth, tHeight)
	if tRotatedX <= minXCoordinate {
		g.player.posX = g.player.prevPosX
	}
	if tRotatedX >= maxXCoordinate - tWidth {
		g.player.posX = g.player.prevPosX
	}
	if tRotatedY <= minYCoordinate {
		g.player.posY = g.player.prevPosY
	}
	if tRotatedY >= maxYCoordinate - tHeight {
		g.player.posY = g.player.prevPosY
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
			g.projectiles[i].posX -= deltaX
			g.projectiles[i].posY -= deltaY
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
	posOffsetX := float64(g.player.hullImage.Bounds().Dx()) / 2
	posOffsetY := float64(g.player.hullImage.Bounds().Dy()) / 2
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
		posX: g.player.posX + deltaX,
		posY: g.player.posY + deltaY,
		explosion1SpriteWidth: 50, explosion1SpriteHeight: 50,
		explosionNumSprites: 8,
		explosionFrame: 0,
		isCollided: false,
		isActive: true,
		image: g.resources.projectileImage,
		explosion1: g.resources.projectileExplImage,
		explosionSpeed: 3,
	})
	g.player.lastShot = g.count
}

func checkProjectileCollision(p *Projectile) bool {
	if p.posX <= minXCoordinate || p.posY <= minYCoordinate || p.posX >= maxXCoordinate - 25 || p.posY >= maxYCoordinate - 25 {
		return  true
	}
	return false
}

func UpdateCollisions(g *Game) {
	for _, block := range g.blocks {
		if g.player.checkBlockCollision(&block) {
			g.player.posX = g.player.prevPosX
			g.player.posY = g.player.prevPosY
			g.player.rotation = g.player.prevRotation
		}
	}
}