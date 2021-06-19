package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yuin/gopher-lua"
)

const (
	screenWidth  = 512
	screenHeight = 512
	level = 1
)

type TileTypeAlias int
type TileType struct{
	TileEmpty TileTypeAlias
	TileSolid TileTypeAlias
}

var TileTypeEnum = &TileType{
	TileEmpty: 0,
	TileSolid: 1, 
}

var TileSolid = 1
var TileEmpty = 0
var size_w int = 0
var size_h int = 0
var genMap [][]int

const (
	tileSize = 32
	tileXNum = 16
)

var (
	tilesImage *ebiten.Image
	tilesImage2 *ebiten.Image
)

func newLevel(x int , y int){	
	for i := 0; i < y+2; i++ {
		m := []int{}
		for j := 0; j < x+3; j++ {
			m = append(m, int(TileTypeEnum.TileEmpty))
		}
		genMap = append(genMap, m)
	}
	
}

func createLevel(L *lua.LState) int{
	size_w = L.ToInt(1)
	size_h = L.ToInt(2)

	newLevel(size_w,size_h)
	return 1

	//parse tilemap?
}

func setTile(L *lua.LState) int{
	y := L.ToInt(2)
	x := L.ToInt(1)	
	switch tile := L.ToInt(3); tile {
	case 0:
		genMap[x][y] = int(TileTypeEnum.TileEmpty)
	case 1:
		genMap[x][y] = int(TileTypeEnum.TileSolid)
	}
	return 1
}


func init() {

	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile("hello.lua"); err != nil {
		panic(err)
	}
	L.SetGlobal("_CreateLevel", L.NewFunction(createLevel))
	L.SetGlobal("_SetTile", L.NewFunction(setTile))
	if err := L.CallByParam(lua.P{
		Fn: L.GetGlobal("LoadLevel"),
		NRet: 1,
		Protect: true,
		}, lua.LNumber(level)); err!= nil {
		panic(err)					
	}
	//fmt.Println(genMap)
	//load images
	var err error
	ship_img, _, err := ebitenutil.NewImageFromFile("img/dirt_light.png")
	ship_img2, _, err2 := ebitenutil.NewImageFromFile("img/dirt_brown.png")
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err2)
	}
	tilesImage = (ship_img)
	tilesImage2 = (ship_img2)
}

type Game struct {
	layers [][]int
	tilemap *tile_map
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
		
	for i, tile := range g.tilemap.gameMap.Layers[0].Tiles {
		print((i))
		print(tile.ID)
		println()
	}

	for i := range g.layers{
		for j := range g.layers[i]{
			op := &ebiten.DrawImageOptions{}
	
			op.GeoM.Translate(float64(i*tileSize), float64(j*tileSize))
			if g.layers[i][j] == int(TileTypeEnum.TileSolid){				
				screen.DrawImage(tilesImage.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image), op)
			} else{
				screen.DrawImage(tilesImage2.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image), op)
			}
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	gameMap := getMap()
	//gameMap.append(1)
	//fmt.Print(gameMap.gameMap.Layers[0].Tiles)
	g := &Game{ tilemap: gameMap}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("2XD2")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
