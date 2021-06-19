package main

import (
	"fmt"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

type vector struct {
	x, y float64
}

type component interface {
	onUpdate() error
	onDraw(*ebiten.Image) error
}

type element struct {
	position vector
	rotation float64
	active bool
	components []component
}

func (elem *element) addComponent(new component){
	for _, existing := range elem.components{
		if reflect.TypeOf(new) == reflect.TypeOf(existing){
			panic(fmt.Sprintf("attemnt to add new component without existing type %v", reflect.TypeOf(new)))
		}
	}
	elem.components = append(elem.components, new)
}

func (elem *element) getComponent(withType component) component {
	typ := reflect.TypeOf(withType)
	for _, comp := range elem.components{
		if reflect.TypeOf(comp) == typ{
			return comp
		}
	}
	panic(fmt.Sprintf("no component with type %v", reflect.TypeOf(withType)))
	return withType
}

var elements []*element