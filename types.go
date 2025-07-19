package main

import (
	//"log"

	//"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Rect interface{
	getCoordinates() (float64, float64, float64, float64, float64)
}

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
	explosionSpeed int // lower is faster

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
	containerImage *ebiten.Image
	playerHullImage *ebiten.Image
	playerTurretImage *ebiten.Image
	playerTracksImage *ebiten.Image
	projectileImage *ebiten.Image
	projectileExplImage *ebiten.Image
}

type Tank struct {
	width, height float64
	posX, posY float64
	prevPosX, prevPosY float64
	rotation, prevRotation, moveSpeed float64
	hullImage *ebiten.Image
	turretImage *ebiten.Image
	tracksImage *ebiten.Image
	reloadSpeed int
	lastShot int
	fireRollbackOffset int
	isMoving bool
}

func (t *Tank) getCoordinates() (float64, float64, float64, float64, float64) {
	return t.posX, t.posY, t.width, t.height, t.rotation
}

func (t *Tank) checkBlockCollision(b *Block) bool {
	tRotatedX, tRotatedY, tWidth, tHeight := getRotatedCoords(t)
	return checkRectCollision(tRotatedX, tRotatedY, tWidth, tHeight, b.posX, b.posY, b.width, b.height)
	//log.Printf("tX: %0f, tY: %0f, tW: %0f, tH: %0f, bX: %0f, bY: %0f", tRotatedX, tRotatedY, tWidth, tHeight, b.posX, b.posY)
}

func (p *Projectile) getCoordinates() (float64, float64, float64, float64, float64) {
	return p.posX, p.posY, p.width, p.height, p.rotation
}

func (p *Projectile) checkBlockCollision(b *Block) bool {
	return checkRectCollision(p.posX, p.posY, p.width, p.height, b.posX, b.posY, b.width, b.height)
}

func (p *Projectile) getExplositionOffset() (float64, float64) {
	switch math.Round(p.rotation) {
	case 0:
		return p.posX - p.explosion1SpriteWidth / 2, p.posY - p.explosion1SpriteHeight / 2
	case math.Round(math.Pi):
		return p.posX - p.explosion1SpriteWidth / 2, p.posY + p.height / 2 - p.explosion1SpriteHeight / 2
	case math.Round(3*math.Pi/2):
		return p.posX + p.width / 2 - p.explosion1SpriteWidth / 2, p.posY - p.explosion1SpriteHeight / 2
	default:
		return p.posX - p.explosion1SpriteWidth / 2, p.posY - p.explosion1SpriteHeight / 2
	}
}