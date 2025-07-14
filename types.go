package main

import (
	//"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Projectile struct {
	width float64
	height float64
	rotation float64
	moveSpeed float64
	posX float64
	posY float64
	image *ebiten.Image
	explosion1 *ebiten.Image
	explosion1SpriteWidth float64
	explosion1SpriteHeight float64
	explosionNumSprites int
	explosionFrame int

	isCollided bool
	isActive bool
}

type Game struct{
	count int
	player Tank
	tanks []Tank
	projectiles []Projectile
	blocks []Block
	resources Resource
}

type Block struct {
	posX, posY float64
	width, height float64
	image *ebiten.Image
}

type Resource struct {
	background *ebiten.Image
	blockImage *ebiten.Image
	playerImage *ebiten.Image
	projectileImage *ebiten.Image
	projectileExplImage *ebiten.Image
}

type Tank struct {
	width, height float64
	posX, posY float64
	prevPosX, prevPosY float64
	rotation, prevRotation, moveSpeed float64
	image *ebiten.Image
	reloadSpeed int
	lastShot int
}

func (t *Tank) checkBlockCollision(b *Block) bool {
	tRotatedX, tRotatedY, tWidth, tHeight := getRotatedCoords(t)
	
	//log.Printf("tX: %0f, tY: %0f, tW: %0f, tH: %0f, bX: %0f, bY: %0f", tRotatedX, tRotatedY, tWidth, tHeight, b.posX, b.posY)
	if b.posX >= tRotatedX && b.posX <= (tRotatedX + tWidth) {
		if b.posY >= tRotatedY && b.posY <= (tRotatedY + tHeight) {
			return true
		}
		if (b.posY + b.height) >= tRotatedY && (b.posY + b.height) <= (tRotatedY + tHeight) {
			return  true
		}
	}
	if (b.posX + b.width) >= tRotatedX && (b.posX + b.width) <= (tRotatedX + tWidth) {
		if b.posY >= tRotatedY && b.posY <= (tRotatedY + tHeight) {
			return true
		}
		if (b.posY + b.height) >= tRotatedY && (b.posY + b.height) <= (tRotatedY + tHeight) {
			return  true
		}
	}
	return false
}