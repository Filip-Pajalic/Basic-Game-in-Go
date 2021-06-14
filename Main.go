package main

import (
	//"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yuin/gopher-lua"
)

const (
	screenWidth  = 240
	screenHeight = 240
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
	tileSize = 16
	tileXNum = 25
)

var (
	tilesImage *ebiten.Image
	tilesImage2 *ebiten.Image
)

func newLevel(x int , y int){	
	//fmt.Println(x)
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

	//for i := range genMap{
		// for j := range genMap[i]{
		// 	if genMap[i][j] == int32(TileTypeEnum.TileSolid){
		// 		rect := sdl.Rect{int32(screenWidth*i), int32(screenHeight*j), screenHeight, screenWidth}
		// 		surface.FillRect(&rect, 0xffff0000)
		// 	} else{
		// 		rect := sdl.Rect{int32(screenWidth*i), int32(screenHeight*j), screenHeight, screenWidth}
		// 		surface.FillRect(&rect, 0x11ff88aa)
		// 	}
		// }
}

func setTile(L *lua.LState) int{
	y := L.ToInt(2)
	//print("Y:"+ fmt.Sprint(y))
	x := L.ToInt(1)
	//println(", X:"+ fmt.Sprint(x))	
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
	fmt.Println(genMap)

	var err error
	ship_img, _, err := ebitenutil.NewImageFromFile("img/ship1.png")
	ship_img2, _, err := ebitenutil.NewImageFromFile("img/ship2.png")
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = (ship_img)
	tilesImage2 = (ship_img2)
}

type Game struct {
	layers [][]int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	const xNum = screenWidth / tileSize
	//fmt.Printf(g.layers)
	//for _, l := range g.layers {
	//	for i, t := range l {
	//		op := &ebiten.DrawImageOptions{}
	//		op.GeoM.Translate(float64((i%xNum)*tileSize), float64((i/xNum)*tileSize))

	//		sx := (t % tileXNum) * tileSize
	//		sy := (t / tileXNum) * tileSize
	//		screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
	//	}
	//}
		
	for i := range g.layers{
		for j := range g.layers[i]{
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xNum)*tileSize), float64((i/xNum)*tileSize))
			//sx := (i % tileXNum) * tileSize
			//sy := (i / tileXNum) * tileSize
			if g.layers[i][j] == int(TileTypeEnum.TileSolid){				
				screen.DrawImage(tilesImage.SubImage(image.Rect(i*20, j*20, 20, 20)).(*ebiten.Image), op)
			} else{
				screen.DrawImage(tilesImage2.SubImage(image.Rect(i*20, j*20, 20, 20)).(*ebiten.Image), op)
			}
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	fmt.Println(genMap)
	g := &Game{ layers: genMap}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("2XD2")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
