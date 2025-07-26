package main

import (
	//"log"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawEnemies(g *Game, screen *ebiten.Image) {
	for _, e := range g.tanks {
		DrawTank(&e, screen, g.count)
	}
}

func UpdateEnemies(g *Game) {
	for i, e := range g.tanks {
		g.tanks[i].isMoving = false
		checkPlayer(&g.tanks[i], g)
		if e.isMovable {
			switch rand.IntN(4) {
			case 0, 1:
				moveEnemy(&g.tanks[i], g)
			case 2:
				rotateEnemy(&g.tanks[i], g)
			case 3:
				continue
			}
		}
	}
}

func moveEnemy(e *Tank, g *Game) {
	e.isMoving = true
	e.prevPosX = e.posX
	e.prevPosY = e.posY
	e.prevRotation = e.rotation
	switch math.Round(e.rotation) {
	case 0:
		e.posY -= e.moveSpeed
		UpdateCollisions(e, g)
	case math.Round(math.Pi):
		e.posY += e.moveSpeed
		UpdateCollisions(e, g)
	case math.Round(3*math.Pi/2):
		e.posX -= e.moveSpeed
		UpdateCollisions(e, g)
	default:
		e.posX += e.moveSpeed
		UpdateCollisions(e, g)
	}
}

func rotateEnemy(e *Tank, g *Game) {
	e.prevPosX = e.posX
	e.prevPosY = e.posY
	e.prevRotation = e.rotation
	if g.count % 20 != 0 {
		return
	}
	switch rand.IntN(2) {
	case 0:
		e.rotation += math.Pi / 2
	case 1:
		e.rotation -= math.Pi / 2
	}
	if e.rotation >= 2 * math.Pi {
		e.rotation -= 2 * math.Pi
	}
	if e.rotation < 0 {
		e.rotation += 2 * math.Pi
	}
	UpdateCollisions(e, g)
}

func checkPlayer(e *Tank, g *Game) {
	tRotatedX, tRotatedY, tWidth, tHeight := getRotatedCoords(e)
	bRotatedX, bRotatedY, bWidth, bHeight := getRotatedCoords(&g.player)
	if checkAxisCollision(e.rotation, tRotatedX, tRotatedY, tWidth, tHeight, bRotatedX, bRotatedY, bWidth, bHeight) {
		fireEnemy(e, g)
	}
}

func fireEnemy(e *Tank, g *Game) {
	if g.count - e.lastShot >= e.reloadSpeed {
		addProjectile(e, g)
	}
}