package main

import (

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawEnemies(g *Game, screen *ebiten.Image) {
	for _, e := range g.tanks {
		DrawTank(&e, screen, g.count)
	}
}

func UpdateEnemies(g *Game, screen *ebiten.Image) {
	
}