package main

import (
	"bytes"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	MplusFaceSource *text.GoTextFaceSource
)

func loadResources(g *Game) {
	var err error
	r := Resource{}
	r.background, _, err = ebitenutil.NewImageFromFile("resources/Ground_Tile_02_A.png")
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
	r.containerImage, _, err = ebitenutil.NewImageFromFile("resources/Container_A.png")
	if err != nil {
		log.Fatal(err)
	}
	r.blockImage, _, err = ebitenutil.NewImageFromFile("resources/Block_C_02.png")
	if err != nil {
		log.Fatal(err)
	}
	r.playerHullImage, _, err = ebitenutil.NewImageFromFile("resources/hull_new_mini2.png")
	if err != nil {
		log.Fatal(err)
	}
	r.playerTracksImage, _, err = ebitenutil.NewImageFromFile("resources/tracks.png")
	if err != nil {
		log.Fatal(err)
	}
	r.playerTurretImage, _, err = ebitenutil.NewImageFromFile("resources/turret_new_mini.png")
	if err != nil {
		log.Fatal(err)
	}
	r.enemy1HullImage, _, err = ebitenutil.NewImageFromFile("resources/e1Hull_04.png")
	if err != nil {
		log.Fatal(err)
	}
	r.enemy1TracksImage, _, err = ebitenutil.NewImageFromFile("resources/tracks.png")
	if err != nil {
		log.Fatal(err)
	}
	r.enemy1TurretImage, _, err = ebitenutil.NewImageFromFile("resources/Gun_03.png")
	if err != nil {
		log.Fatal(err)
	}
	r.enemy2HullImage, _, err = ebitenutil.NewImageFromFile("resources/Hull_08.png")
	if err != nil {
		log.Fatal(err)
	}
	r.enemy2TracksImage, _, err = ebitenutil.NewImageFromFile("resources/tracks.png")
	if err != nil {
		log.Fatal(err)
	}
	r.enemy2TurretImage, _, err = ebitenutil.NewImageFromFile("resources/Gun_04.png")
	if err != nil {
		log.Fatal(err)
	}
	r.enemy3HullImage, _, err = ebitenutil.NewImageFromFile("resources/Hull_02.png")
	if err != nil {
		log.Fatal(err)
	}
	r.enemy3TracksImage, _, err = ebitenutil.NewImageFromFile("resources/tracks.png")
	if err != nil {
		log.Fatal(err)
	}
	r.enemy3TurretImage, _, err = ebitenutil.NewImageFromFile("resources/Gun_01.png")
	if err != nil {
		log.Fatal(err)
	}
	r.enemy4HullImage, _, err = ebitenutil.NewImageFromFile("resources/Hull_05.png")
	if err != nil {
		log.Fatal(err)
	}
	r.enemy4TracksImage, _, err = ebitenutil.NewImageFromFile("resources/tracks.png")
	if err != nil {
		log.Fatal(err)
	}
	r.enemy4TurretImage, _, err = ebitenutil.NewImageFromFile("resources/Gun_06.png")
	if err != nil {
		log.Fatal(err)
	}
	r.tankExplImage, _, err = ebitenutil.NewImageFromFile("resources/Explosion_A_01.png")
	if err != nil {
		log.Fatal(err)
	}
	r.menuImage, _, err = ebitenutil.NewImageFromFile("resources/menu.png")
	if err != nil {
		log.Fatal(err)
	}
	r.audioFile, err = os.ReadFile("resources/8bit-music-for-game-68698.mp3")
	if err != nil {
		log.Fatal(err)
	}

	MplusFaceSource, err = text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}

	g.resources = r
}

func getRotatedCoords(t Rect) (float64, float64, float64, float64) {
	var tWidth, tHeight, tRotatedX, tRotatedY float64
	x, y, width, height, rotation := t.getCoordinates()
	if rotation == 0 || (math.Round(rotation) == math.Round(math.Pi)) {
		tWidth = width
		tHeight = height
		tRotatedX = x
		tRotatedY = y
	} else {
		tWidth = height
		tHeight = width
		tRotatedX = (x + width / 2) - (height / 2)
		tRotatedY = (y + height / 2) - (width / 2)
	}
	return tRotatedX, tRotatedY, tWidth, tHeight
}


func getTurretOffset(t *Tank, is_rollback bool) (float64, float64) {
	var turretOffsetX, turretOffsetY float64
	if is_rollback {
		if t.rotation == 0 {
			turretOffsetX = (t.width / 2 - float64(t.turretImage.Bounds().Dx()) / 2)
			turretOffsetY = float64(t.fireRollbackOffset)
		} else if math.Round(t.rotation) == math.Round(math.Pi) {
			turretOffsetX = - (t.width / 2 - float64(t.turretImage.Bounds().Dx()) / 2)
			turretOffsetY = -float64(t.fireRollbackOffset)
		} else if math.Round(t.rotation) == math.Round(3*math.Pi/2) {
			turretOffsetX = float64(t.fireRollbackOffset)
			turretOffsetY = - (t.width / 2 - float64(t.turretImage.Bounds().Dx()) / 2)
		} else {
			turretOffsetX = -float64(t.fireRollbackOffset)
			turretOffsetY = (t.width / 2 - float64(t.turretImage.Bounds().Dx()) / 2)
		}
	} else {
		if t.rotation == 0 {
			turretOffsetX = (t.width / 2 - float64(t.turretImage.Bounds().Dx()) / 2)
			turretOffsetY = 0
		} else if math.Round(t.rotation) == math.Round(math.Pi) {
			turretOffsetX = - (t.width / 2 - float64(t.turretImage.Bounds().Dx()) / 2)
			turretOffsetY = 0
		} else if math.Round(t.rotation) == math.Round(3*math.Pi/2) {
			turretOffsetX = 0
			turretOffsetY = - (t.width / 2 - float64(t.turretImage.Bounds().Dx()) / 2)
		} else {
			turretOffsetX = 0
			turretOffsetY = (t.width / 2 - float64(t.turretImage.Bounds().Dx()) / 2)
		}
	}
	return turretOffsetX, turretOffsetY
}

func getTracksOffset(t *Tank, is_left bool) (float64, float64) {
	var offsetX, offsetY float64
	if is_left {
		offsetX, offsetY = 0, 0
	} else {
		if t.rotation == 0 {
			offsetX = t.width - float64(t.tracksImage.Bounds().Dx())
			offsetY = 0
		} else if math.Round(t.rotation) == math.Round(math.Pi) {
			offsetX = - t.width + float64(t.tracksImage.Bounds().Dx())
			offsetY = 0
		} else if math.Round(t.rotation) == math.Round(3*math.Pi/2) {
			offsetX = 0
			offsetY = - t.width + float64(t.tracksImage.Bounds().Dx())
		} else {
			offsetX = 0
			offsetY = t.width - float64(t.tracksImage.Bounds().Dx())
		}
	}

	return offsetX, offsetY
}

func checkRectCollision(aRotation, aX, aY, aWidth, aHeight, bX, bY, bWidth, bHeight float64) bool {
	if (aX >= bX) && (aX <= (bX + bWidth)) {
		if (aY >= bY) && (aY <= (bY + bHeight)) {
			if (math.Round(aRotation) == math.Round(math.Pi / 2)) || (math.Round(aRotation) == math.Round(math.Pi)) {
				return false
			} else {
				return true
			}
		} else if ((aY + aHeight) >= bY) && ((aY + aHeight) <= (bY + bHeight)) {
			if (aRotation == math.Round(math.Pi / 2)) || (math.Round(aRotation) == 0) {
				return false
			} else {
				return true
			}
		}
	}
	if ((aX + aWidth) >= bX) && ((aX + aWidth) <= (bX + bWidth)) {
		if (aY >= bY) && (aY <= (bY + bHeight)) {
			if (math.Round(aRotation) == math.Round(3*math.Pi / 2)) || (math.Round(aRotation) == math.Round(math.Pi)) {
				return false
			} else {
				return true
			}
		} else if ((aY + aHeight) >= bY) && ((aY + aHeight) <= (bY + bHeight)) {
			if (aRotation == math.Round(3* math.Pi / 2)) || (math.Round(aRotation) == 0) {
				return false
			} else {
				return true
			}
		}
	}
	return false
}

func checkAxisCollision(aRotation, aX, aY, aWidth, aHeight, bX, bY, bWidth, bHeight float64) bool {
	switch math.Round(aRotation) {
	case 0:
		if (aY >= bY) {
			if ((aX + aWidth) >= bX) && ((aX + aWidth) <= (bX + bWidth)) || (aX >= bX) && (aX <= (bX + bWidth)) {
				return  true
			}
		}
	case math.Round(3*math.Pi / 2):
		if (aX + aWidth) >= bX {
			if ((aY + aHeight) >= bY) && ((aY + aHeight) <= (bY + bHeight)) || (aY >= bY) && (aY <= (bY + bHeight)) {
				return  true
			}
		}
	case math.Round(math.Pi/2):
		if aX <= (bX + bWidth) {
			if ((aY + aHeight) >= bY) && ((aY + aHeight) <= (bY + bHeight)) || (aY >= bY) && (aY <= (bY + bHeight)) {
				return  true
			}
		}
	default:
		if (aY <= bY) {
			if ((aX + aWidth) >= bX) && ((aX + aWidth) <= (bX + bWidth)) || (aX >= bX) && (aX <= (bX + bWidth)) {
				return  true
			}
		}
	}
	return false
}

func UpdateCollisions(t *Tank, g *Game) {
	if t != &g.player {
		if t.checkBlockCollision(&g.player) {
			t.posX = t.prevPosX
			t.posY = t.prevPosY
			t.rotation = t.prevRotation
		}
	}
	for _, block := range g.blocks {
		if t.checkBlockCollision(&block) {
			t.posX = t.prevPosX
			t.posY = t.prevPosY
			t.rotation = t.prevRotation
		}
	}
	for i, tank := range g.tanks {
		if &g.tanks[i] == t {
			continue
		}
		if t.checkBlockCollision(&tank) {
			t.posX = t.prevPosX
			t.posY = t.prevPosY
			t.rotation = t.prevRotation
		}
	}
	tRotatedX, tRotatedY, tWidth, tHeight  := getRotatedCoords(t)
	//log.Printf("tX: %0f, tY: %0f, tW: %0f, tH: %0f", tRotatedX, tRotatedY, tWidth, tHeight)
	if tRotatedX <= minXCoordinate {
		t.posX = t.prevPosX
	}
	if tRotatedX >= maxXCoordinate - tWidth {
		t.posX = t.prevPosX
	}
	if tRotatedY <= minYCoordinate {
		t.posY = t.prevPosY
	}
	if tRotatedY >= maxYCoordinate - tHeight {
		t.posY = t.prevPosY
	}
}

func CheckCollisions(t *Tank, g *Game) bool {
	if t != &g.player {
		if t.checkBlockCollision(&g.player) {
			return true
		}
	}
	for _, block := range g.blocks {
		if t.checkBlockCollision(&block) {
			return true
		}
	}
	for i, tank := range g.tanks {
		if &g.tanks[i] == t {
			continue
		}
		if t.checkBlockCollision(&tank) {
			return  true
		}
	}
	return false
}