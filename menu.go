package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)
const (
	normalFontSize = 18
)

func DrawGameOverScreen(screen *ebiten.Image) {
	gameOverText := "Game Over"
	op := &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 6*normalFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = normalFontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, gameOverText, &text.GoTextFace{
		Source: MplusFaceSource,
		Size:   normalFontSize,
	}, op)

	gameOverText = "Press Enter to try again"
	op = &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 8*normalFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = normalFontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, gameOverText, &text.GoTextFace{
		Source: MplusFaceSource,
		Size:   normalFontSize/2,
	}, op)
}

func DrawGameWonScreen(screen *ebiten.Image) {
	gameOverText := "You Won!"
	op := &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 6*normalFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = normalFontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, gameOverText, &text.GoTextFace{
		Source: MplusFaceSource,
		Size:   normalFontSize,
	}, op)

	gameOverText = "Press Enter to try again"
	op = &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 8*normalFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = normalFontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, gameOverText, &text.GoTextFace{
		Source: MplusFaceSource,
		Size:   normalFontSize/2,
	}, op)
}

func DrawGameStartScreen(screen *ebiten.Image, g *Game) {
	screen.DrawImage(g.resources.menuImage, nil)
	op := &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, normalFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = normalFontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, "Tank Battle", &text.GoTextFace{
		Source: MplusFaceSource,
		Size:   normalFontSize,
	}, op)

	op = &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 8*normalFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = normalFontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, "Press Enter Key to Start", &text.GoTextFace{
		Source: MplusFaceSource,
		Size:   normalFontSize/1.5,
	}, op)

	op = &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 14*normalFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = normalFontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, "Controls", &text.GoTextFace{
		Source: MplusFaceSource,
		Size:   normalFontSize/2,
	}, op)

	op = &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 15*normalFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = normalFontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, "WASD - move, Space - fire", &text.GoTextFace{
		Source: MplusFaceSource,
		Size:   normalFontSize/2.5,
	}, op)
}