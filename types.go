package main

import (

	"github.com/hajimehoshi/ebiten/v2"
)

type Tank struct {
	width, height float64
	positionX, positionY float64
	rotation, moveSpeed float64
	image *ebiten.Image
	reloadSpeed int
	lastShot int
}

type Projectile struct {
	width float64
	height float64
	rotation float64
	moveSpeed float64
	positionX float64
	positionY float64
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
	posX, posY int
	image *ebiten.Image
}

type Resource struct {
	background *ebiten.Image
	blockImage *ebiten.Image
	playerImage *ebiten.Image
	projectileImage *ebiten.Image
	projectileExplImage *ebiten.Image
}