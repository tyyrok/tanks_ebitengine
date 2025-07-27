package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)
const (
	normalFontSize = 24
	offsetY1 = 50
	offsetY2 = 80
)

func DrawGameOverScreen(screen *ebiten.Image) {
	gameOverText := "Game Over"
	op := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign: text.AlignCenter}}
	op.GeoM.Translate(float64(screen.Bounds().Dx())/2, offsetY1)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, gameOverText, &text.GoTextFace{
		Source: MplusFaceSource,
		Size:   normalFontSize,
	}, op)

	gameOverText = "Press Enter to try again"
	op = &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign: text.AlignCenter}}
	op.GeoM.Translate(float64(screen.Bounds().Dx())/2, offsetY2)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, gameOverText, &text.GoTextFace{
		Source: MplusFaceSource,
		Size:   normalFontSize,
	}, op)
}

func DrawGameWonScreen(screen *ebiten.Image) {
	gameOverText := "You Won!"
	op := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign: text.AlignCenter}}
	op.GeoM.Translate(float64(screen.Bounds().Dx())/2, offsetY1)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, gameOverText, &text.GoTextFace{
		Source: MplusFaceSource,
		Size:   normalFontSize,
	}, op)

	gameOverText = "Press Enter to try again"
	op = &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign: text.AlignCenter}}
	op.GeoM.Translate(float64(screen.Bounds().Dx())/2, offsetY2)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, gameOverText, &text.GoTextFace{
		Source: MplusFaceSource,
		Size:   normalFontSize,
	}, op)
}