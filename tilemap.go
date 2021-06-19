package main

import (
	"fmt"
	"os"
	"github.com/lafriks/go-tiled"
)

const mapPath = "maps/level1.tmx"

type tile_map struct {
	tileSize int
	gameMap *tiled.Map
}

func getMap() *tile_map{
	tileMap := tile_map{20,loadMap()}
	tileMap.tileSize = 20
	return &tileMap
}


func loadMap() *tiled.Map{
	gameMap, err := tiled.LoadFromFile(mapPath)
	//tileMap.gameMap := gameMap
	if err != nil {
		fmt.Printf("error parsing map: %s", error.Error)
		os.Exit(2)
	}
	//fmt.Println(gameMap)
	return gameMap
}