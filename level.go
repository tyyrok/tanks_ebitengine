package main

import (

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawLevel(g *Game, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, Gamer!")
	screen.DrawImage(g.resources.background, nil)

}