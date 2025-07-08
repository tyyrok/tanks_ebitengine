package main

import (

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawLevel(g *Game, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, Gamer!")
	screen.DrawImage(g.resources.background, nil)
	for _, block := range g.blocks {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(block.posX, block.posY)
		screen.DrawImage(block.image, op)
	}
}

func initLevel(g *Game) {
	for i := range 5 {
		g.blocks = append(g.blocks, Block{
			posX: float64(60*i),
			posY: float64(50*i),
			width: 18,
			height: 18,
			image: g.resources.blockImage,
		})
	}
}