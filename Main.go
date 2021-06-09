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

var size_w int = 0
var size_h int = 0
var map_m string = ""

func createLevel(L *lua.LState) int{
	size_w = L.ToInt(1)
	size_h = L.ToInt(2)
	return 1
}

func setTile(L *lua.LState) int{
	switch tile := L.ToInt(3); tile {
	case 1:

	case 2:

	}
	map_m = L.ToString(1)
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
	print(fmt.Sprint(size_w) + " " + fmt.Sprint(size_h))

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

	rect := sdl.Rect{0, 0, screenHeight, screenWidth}
	surface.FillRect(&rect, 0xffff0000)
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
