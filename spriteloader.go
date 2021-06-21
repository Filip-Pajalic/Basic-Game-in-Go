package main

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type Spriteloader struct {
	spriteData *goaseprite.File
	spriteTexture *ebiten.Image
}

func newSprite(spritename string) *Spriteloader{
	spriteloader := Spriteloader{}
	spriteloader.spriteData = goaseprite.ReadFile("img/"+spritename+".json")
	var err error
	spriteloader.spriteTexture, _, err = ebitenutil.NewImageFromFile("img/"+spritename+"-sheet.png")
	if err != nil {
		log.Fatal(err)
	}
	return &spriteloader
}