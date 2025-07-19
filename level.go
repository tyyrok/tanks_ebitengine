package main

import (

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

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
	//screen.DrawImage(g.resources.background, nil)
	for _, block := range g.blocks {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(block.posX, block.posY)
		screen.DrawImage(block.image, op)
	}
}

func initLevel(g *Game) {
	for i := range 2 {
		g.blocks = append(g.blocks, Block{
			posX: float64(100*i),
			posY: float64(100*i),
			width: 51,
			height: 26,
			image: g.resources.blockImage,
		})
	}
}