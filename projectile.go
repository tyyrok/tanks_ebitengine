package main

import (
	"fmt"
	"image"

	//"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawProjectiles(g *Game, screen *ebiten.Image) {
	for i, shot := range g.projectiles{
		if !shot.isActive {
			continue
		}
		op := &ebiten.DrawImageOptions{}
		if shot.isCollided {
			offsetX, offsetY := shot.getExplositionOffset()
			op.GeoM.Translate(offsetX, offsetY)
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
			op.GeoM.Scale(shot.scale, shot.scale)
			op.GeoM.Translate(shot.posX, shot.posY)
			screen.DrawImage(shot.image, op)
		}
		msg := fmt.Sprintf("\n\n\n\n\nprojectiles: %d", len(g.projectiles))
		ebitenutil.DebugPrint(screen, msg)
	}
}

func UpdateProjectiles(g *Game) {
	updatedProjectiles := []Projectile{}
	for i, shot := range g.projectiles{
		if !shot.isActive {
			continue
		}
		if !shot.isCollided {
			//log.Println(shot.posX, shot.posY)
			deltaY := math.Cos(shot.rotation) * shot.moveSpeed
			deltaX := -math.Sin(shot.rotation) * shot.moveSpeed
			g.projectiles[i].posX -= deltaX
			g.projectiles[i].posY -= deltaY
			if checkProjectileCollision(&shot, g) {
				g.projectiles[i].isCollided = true
			}
		}
		updatedProjectiles = append(updatedProjectiles, g.projectiles[i])
	}
	g.projectiles = updatedProjectiles
}

func addProjectile(t *Tank, g *Game) {
	// Calculate offset for spawning a projectile
	posOffsetX := float64(t.hullImage.Bounds().Dx()) / 2
	posOffsetY := float64(t.hullImage.Bounds().Dy()) / 2
	deltaX := posOffsetX * math.Abs(math.Cos(t.rotation)) + posOffsetX * math.Abs(math.Sin(t.rotation))
	deltaY := posOffsetY * math.Abs(math.Cos(t.rotation)) + posOffsetY * math.Abs(math.Sin(t.rotation))
	if t.rotation == 0 || t.rotation == math.Pi {
		deltaY -= posOffsetY * (1 / math.Cos(t.rotation)) + 2
	} else {
		deltaX += 2 * posOffsetX * math.Sin(t.rotation)
	}
	g.projectiles = append(g.projectiles, Projectile{
		width: float64(g.resources.projectileImage.Bounds().Dx()),
		height: float64(g.resources.projectileImage.Bounds().Dy()),
		rotation: t.rotation,
		moveSpeed: t.moveSpeed + 1,
		posX: t.posX + deltaX,
		posY: t.posY + deltaY,
		explosion1SpriteWidth: 50, explosion1SpriteHeight: 50,
		explosionNumSprites: 8,
		explosionFrame: 0,
		scale: 0.3,
		isCollided: false,
		isActive: true,
		image: g.resources.projectileImage,
		explosion1: g.resources.projectileExplImage,
		explosionSpeed: 3,
	})
	t.lastShot = g.count
}

func checkProjectileCollision(p *Projectile, g *Game) bool {
	for i, block := range g.blocks {
		if p.checkBlockCollision(&block) {
			if block.isDestroyable {
				g.blocks[i].isShot = true
			}
			return true
		}
	}
	for i, e := range g.tanks {
		if p.checkBlockCollision(&e) {
			g.tanks[i].isShot = true
			return  true
		}
	}
	if p.checkBlockCollision(&g.player) {
		g.player.isShot = true
		return  true
	}
	if p.posX <= minXCoordinate || p.posY <= minYCoordinate || p.posX >= maxXCoordinate || p.posY >= maxYCoordinate {
		return  true
	}
	return false
}