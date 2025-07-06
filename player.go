package main

import (

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawPlayer(p *Tank, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	baseOffsetX := float64(p.image.Bounds().Dx()) / 2
	baseOffsetY := float64(p.image.Bounds().Dy()) / 2
	op.GeoM.Translate(-baseOffsetX, -baseOffsetY)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(p.positionX, p.positionY)
	screen.DrawImage(p.image, op)
}