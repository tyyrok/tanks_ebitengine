package main

import (
	//"log"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawEnemies(g *Game, screen *ebiten.Image) {
	for i := 0; i < len(g.tanks); i++ {
		DrawTank(&g.tanks[i], screen, g.count)
	}
}

func UpdateEnemies(g *Game) {
	updatedTanks := []Tank{}
	for i, e := range g.tanks {
		if !e.isActive {
			g.enemyKilledCount += 1
			changeEnemyType(g)
			continue
		}
		if !e.isShot {
			updateEnemyByType(&g.tanks[i], g)
		}
		updatedTanks = append(updatedTanks, g.tanks[i])
	}
	g.tanks = updatedTanks
	if len(g.tanks) < 2 && g.enemyKilledCount <= GameKillsTreshhold - 1 {
		spawnEnemyOnRandomSpawnPlace(g)
	}
	if g.enemyKilledCount == GameKillsTreshhold+1 {
		g.isWon = true
	}
}

func updateEnemyByType(e *Tank, g *Game) {
	e.isMoving = false
	checkPlayer(e, g)
	if e.isMovable {
		switch e.enemyType {
		case 0, 3:
			switch rand.IntN(4) {
			case 0, 1:
				moveEnemy(e, g)
			case 2:
				rotateEnemy(e, g)
			case 3:
			}
		case 1:
			switch rand.IntN(3) {
			case 0, 1:
				moveEnemy(e, g)
			case 2:
				rotateEnemy(e, g)
			}
		case 2:
			switch rand.IntN(4) {
			case 0:
				moveEnemy(e, g)
			case 1, 2:
				rotateEnemy(e, g)
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

func changeEnemyType(g *Game) {
	check := g.enemyKilledCount / EnemyChangeTreshhold
	switch check {
	case 3:
		g.nextEnemyType = 3
	case 2:
		g.nextEnemyType = 2
	case 1:
		g.nextEnemyType = 1
	}
}

func spawnEnemy(x, y int, g *Game) {
	newEnemy := Tank{
		width: float64(g.resources.playerHullImage.Bounds().Dx()),
		height: float64(g.resources.playerHullImage.Bounds().Dy()),
		posX: float64(x*LevelCellOffsetX), posY: float64(y*LevelCellOffsetY),
		prevPosX: float64(x*LevelCellOffsetX), prevPosY: float64(y*LevelCellOffsetY),
		rotation: math.Pi,
		lastShot: 0,
		scale: 1,
		fireRollbackOffset: 2,
		isMoving: false,
		isMovable: true,
		isShot: false,
		isActive: true,
		explosionNumSprites: 5,
		explosionFrame: 0,
		explosionSpeed: 3,
		explosionImage: g.resources.tankExplImage,
		isDoubleFire: false,
	}
	switch g.nextEnemyType {
	case 0:
		newEnemy.moveSpeed = 1
		newEnemy.reloadSpeed = enemyReloadSpeed1
		newEnemy.hullImage = g.resources.enemy1HullImage
		newEnemy.turretImage = g.resources.enemy1TurretImage
		newEnemy.tracksImage = g.resources.enemy1TracksImage
		newEnemy.enemyType = 0
	case 1:
		newEnemy.moveSpeed = 1.2
		newEnemy.reloadSpeed = enemyReloadSpeed2
		newEnemy.hullImage = g.resources.enemy2HullImage
		newEnemy.turretImage = g.resources.enemy2TurretImage
		newEnemy.tracksImage = g.resources.enemy2TracksImage
		newEnemy.enemyType = 1
	case 2:
		newEnemy.moveSpeed = 1
		newEnemy.reloadSpeed = enemyReloadSpeed2
		newEnemy.hullImage = g.resources.enemy3HullImage
		newEnemy.turretImage = g.resources.enemy3TurretImage
		newEnemy.tracksImage = g.resources.enemy3TracksImage
		newEnemy.enemyType = 2
	case 3:
		newEnemy.moveSpeed = 0.8
		newEnemy.reloadSpeed = enemyReloadSpeed2
		newEnemy.hullImage = g.resources.enemy4HullImage
		newEnemy.turretImage = g.resources.enemy4TurretImage
		newEnemy.tracksImage = g.resources.enemy4TracksImage
		newEnemy.enemyType = 3
		newEnemy.isDoubleFire = true
	}
	if !CheckCollisions(&newEnemy, g) {
		g.tanks = append(g.tanks, newEnemy)
	}
}

func spawnEnemyOnRandomSpawnPlace(g *Game) {
	randomPlace := g.spawnPlaces[rand.IntN(len(g.spawnPlaces))]
	spawnEnemy(randomPlace[0], randomPlace[1], g)
}