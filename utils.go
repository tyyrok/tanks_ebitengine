package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func loadResources(g *Game) {
	var err error
	r := Resource{}
	r.background, _, err = ebitenutil.NewImageFromFile("resources/back.png")
	if err != nil {
		log.Fatal(err)
	}
	r.projectileImage, _, err = ebitenutil.NewImageFromFile("resources/Light_Shell.png") 
	if err != nil {
		log.Fatal(err)
	}
	r.projectileExplImage, _, err = ebitenutil.NewImageFromFile("resources/explosion_1.png") 
	if err != nil {
		log.Fatal(err)
	}
	r.blockImage, _, err = ebitenutil.NewImageFromFile("resources/tile_0018.png")
	if err != nil {
		log.Fatal(err)
	}
	r.playerImage, _, err = ebitenutil.NewImageFromFile("resources/tankAtiny2.png")
	if err != nil {
		log.Fatal(err)
	}

	g.resources = r
}