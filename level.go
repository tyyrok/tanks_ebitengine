package main

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	LevelCellOffsetX = 30
	LevelCellOffsetY = 30
)


var levelObjects = map[int][]uint16{
	0:{0, 0, 0, 0, 0, 0, 0, 0, 4, 0},
	1:{2, 0, 0, 0, 0, 0, 1, 0, 0, 0},
	2:{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	3:{0, 4, 0, 1, 0, 0, 0, 0, 1, 0},
	4:{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	5:{0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
	6:{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
	7:{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	8:{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	9:{0, 0, 0, 3, 0, 3, 0, 0, 0, 0},
}

func DrawLevel(g *Game, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, Gamer!")
	tileWidth, tileHeight := g.resources.background.Bounds().Dx(), g.resources.background.Bounds().Dy()
	for x := 0; x < ScreenWidth; x += tileWidth {
		for y := 0; y < ScreenHeight; y += tileHeight {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(g.resources.background, op)
		}
	}
	for _, block := range g.blocks {
		op := &ebiten.DrawImageOptions{}
		baseOffsetX := float64(block.image.Bounds().Dx()) / 2
		baseOffsetY := float64(block.image.Bounds().Dy()) / 2
		op.GeoM.Translate(-baseOffsetX, -baseOffsetY)
		op.GeoM.Rotate(block.rotation)
		op.GeoM.Translate(block.posX+baseOffsetX, block.posY+baseOffsetY)
		screen.DrawImage(block.image, op)
	}
	msg := fmt.Sprintf("\nKilled: %d", g.enemyKilledCount)
	ebitenutil.DebugPrint(screen, msg)
}

func UpdateLevel(g *Game) {
	blocksToRemove := map[int]int{}
	for i, block := range g.blocks {
		if block.isDestroyable && block.isShot {
			blocksToRemove[i] = 1
		}
	}
	if len(blocksToRemove) > 0 {
		var newBlocks []Block
		for i := 0; i < len(g.blocks); i++ {
			if _, ok := blocksToRemove[i]; !ok {
				newBlocks = append(newBlocks, g.blocks[i])
			}
		}
		g.blocks = newBlocks
	}
}

func initLevel(g *Game) {
	for k, v := range levelObjects {
		for i, elem := range v {
			switch elem {
			case 1:
				g.blocks = append(g.blocks, Block{
					posX: float64(i*LevelCellOffsetX),
					posY: float64(k*LevelCellOffsetY),
					rotation: 0,
					width: float64(g.resources.blockImage.Bounds().Dx()),
					height: float64(g.resources.blockImage.Bounds().Dy()),
					image: g.resources.blockImage,
					isDestroyable: true,
					isShot: false,
				})
			case 2:
				g.blocks = append(g.blocks, Block{
					posX: float64(i*LevelCellOffsetX),
					posY: float64(k*LevelCellOffsetY),
					rotation: 0,
					width: float64(g.resources.containerImage.Bounds().Dx()),
					height: float64(g.resources.containerImage.Bounds().Dy()),
					image: g.resources.containerImage,
					isDestroyable: false,
					isShot: false,
				})
			case 3:
				g.blocks = append(g.blocks, Block{
					posX: float64(i*LevelCellOffsetX),
					posY: float64(k*LevelCellOffsetY),
					rotation: math.Pi /2,
					width: float64(g.resources.containerImage.Bounds().Dx()),
					height: float64(g.resources.containerImage.Bounds().Dy()),
					image: g.resources.containerImage,
				})
			case 4:
				g.tanks = append(g.tanks, Tank{
					width: float64(g.resources.playerHullImage.Bounds().Dx()),
					height: float64(g.resources.playerHullImage.Bounds().Dy()),
					posX: float64(i*LevelCellOffsetX), posY: float64(k*LevelCellOffsetY),
					prevPosX: float64(i*LevelCellOffsetX), prevPosY: float64(k*LevelCellOffsetY),
					rotation: math.Pi, moveSpeed: 1,
					reloadSpeed: 60, lastShot: 0,
					scale: 1,
					hullImage: g.resources.enemy1HullImage,
					turretImage: g.resources.enemy1TurretImage,
					tracksImage: g.resources.playerTracksImage,
					fireRollbackOffset: 2,
					isMoving: false,
					isMovable: true,
					isShot: false,
					isActive: true,
					explosionNumSprites: 5,
					explosionFrame: 0,
					explosionSpeed: 3,
					explosionImage: g.resources.tankExplImage,
				})
			}
		}
	}
}