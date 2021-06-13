package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/yuin/gopher-lua"
)

const (
	screenWidth  = 20
	screenHeight = 20	
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
var genMap [][]int32

//perhaps make struct for 2d array for tiles?

func newLevel(x int , y int){	
	fmt.Println(x)
	for i := 0; i < y+2; i++ {
		m := []int32{}
		for j := 0; j < x+3; j++ {
			m = append(m, int32(TileTypeEnum.TileEmpty))
		}
		genMap = append(genMap, m)
	}
	
}

func createLevel(L *lua.LState) int{
	size_w = L.ToInt(1)
	size_h = L.ToInt(2)

	newLevel(size_w,size_h)
	return 1
}

func setTile(L *lua.LState) int{
	y := L.ToInt(2)
	print("Y:"+ fmt.Sprint(y))
	x := L.ToInt(1)
	println(", X:"+ fmt.Sprint(x))
	
	
	switch tile := L.ToInt(3); tile {
	case 0:
		genMap[x][y] = int32(TileTypeEnum.TileEmpty)
	case 1:
		genMap[x][y] = int32(TileTypeEnum.TileSolid)
	}
	return 1
}


func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

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

	window, err := sdl.CreateWindow("GameWindow", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenHeight*int32(size_h), screenWidth*int32(size_w), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	for i := range genMap{
		for j := range genMap[i]{
			if genMap[i][j] == int32(TileTypeEnum.TileSolid){
				rect := sdl.Rect{int32(screenWidth*i), int32(screenHeight*j), screenHeight, screenWidth}
				surface.FillRect(&rect, 0xffff0000)
			} else{
				rect := sdl.Rect{int32(screenWidth*i), int32(screenHeight*j), screenHeight, screenWidth}
				surface.FillRect(&rect, 0x11ff88aa)
			}
		}

	}
	window.UpdateSurface()
	
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
			}
		}
	}
}
