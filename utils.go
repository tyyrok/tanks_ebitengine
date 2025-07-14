package main

import (
	"log"
	"math"

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

func getRotatedCoords(t *Tank) (float64, float64, float64, float64) {
	var tWidth, tHeight, tRotatedX, tRotatedY float64
	if t.rotation == 0 || math.Round(t.rotation) == math.Round(math.Pi) {
		tWidth = t.width
		tHeight = t.height
		tRotatedX = t.posX
		tRotatedY = t.posY
	} else {
		tWidth = t.height
		tHeight = t.width
		tRotatedX = (t.posX + t.width / 2) - (t.height / 2)
		tRotatedY = (t.posY + t.height / 2) - (t.width / 2)
	}
	return tRotatedX, tRotatedY, tWidth, tHeight
}