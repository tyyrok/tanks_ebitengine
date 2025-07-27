package main

import (
	"math"
	//"log"

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
	scale float64
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
	enemyKilledCount int
	spawnPlaces [][]int
	nextEnemyType int
	isWon bool
}

type Block struct {
	posX, posY float64
	width, height float64
	rotation float64
	image *ebiten.Image
	isDestroyable bool
	isShot bool
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
	enemy1HullImage *ebiten.Image
	enemy1TracksImage *ebiten.Image
	enemy1TurretImage *ebiten.Image
	enemy2HullImage *ebiten.Image
	enemy2TracksImage *ebiten.Image
	enemy2TurretImage *ebiten.Image
	enemy3HullImage *ebiten.Image
	enemy3TracksImage *ebiten.Image
	enemy3TurretImage *ebiten.Image
	enemy4HullImage *ebiten.Image
	enemy4TracksImage *ebiten.Image
	enemy4TurretImage *ebiten.Image
	tankExplImage *ebiten.Image

}

type Tank struct {
	width, height float64
	posX, posY float64
	prevPosX, prevPosY float64
	rotation, prevRotation, moveSpeed float64
	scale float64
	hullImage *ebiten.Image
	turretImage *ebiten.Image
	tracksImage *ebiten.Image
	reloadSpeed int
	lastShot int
	fireRollbackOffset int
	isMoving bool
	isMovable bool
	isShot bool
	isDoubleFire bool
	isActive bool
	explosionImage *ebiten.Image
	explosionNumSprites int
	explosionFrame int
	explosionSpeed int // lower is faster
	enemyType int
}


func (b *Block) getCoordinates() (float64, float64, float64, float64, float64) {
	return b.posX, b.posY, b.width, b.height, b.rotation
}

func (t *Tank) getCoordinates() (float64, float64, float64, float64, float64) {
	return t.posX, t.posY, t.width, t.height, t.rotation
}

func (t *Tank) checkBlockCollision(b Rect) bool {
	tRotatedX, tRotatedY, tWidth, tHeight := getRotatedCoords(t)
	bRotatedX, bRotatedY, bWidth, bHeight := getRotatedCoords(b)
	//return checkRectCollision(tRotatedX, tRotatedY, tWidth, tHeight, bRotatedX, bRotatedY, bWidth, bHeight)
	return checkRectCollision(t.rotation, tRotatedX, tRotatedY, tWidth, tHeight, bRotatedX, bRotatedY, bWidth, bHeight)
	//log.Printf("tX: %0f, tY: %0f, tW: %0f, tH: %0f, bX: %0f, bY: %0f", tRotatedX, tRotatedY, tWidth, tHeight, b.posX, b.posY)
}

func (t *Tank) getExplositionOffset() (float64, float64) {
	tRotatedX, tRotatedY, tWidth, tHeight := getRotatedCoords(t)
	return tRotatedX + tWidth / 2 - float64(t.explosionImage.Bounds().Dy()) / 2, tRotatedY + tHeight / 2 - float64(t.explosionImage.Bounds().Dy()) / 2
}

func (p *Projectile) getCoordinates() (float64, float64, float64, float64, float64) {
	return p.posX, p.posY, p.width*p.scale, p.height*p.scale, p.rotation
}

func (p *Projectile) checkBlockCollision(b Rect) bool {
	rotatedX, rotatedY, width, height := getRotatedCoords(p)
	bRotatedX, bRotatedY, bWidth, bHeight := getRotatedCoords(b)
	//log.Printf("aX: %0f, aY: %0f, aW: %0f, aH: %0f, bX: %0f, bY: %0f, bW: %0f, bH: %0f", rotatedX, rotatedY, width, height, b.posX, b.posY, b.width, b.height)
	return checkRectCollision(p.rotation, rotatedX, rotatedY, width*p.scale, height*p.scale, bRotatedX, bRotatedY, bWidth, bHeight)
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