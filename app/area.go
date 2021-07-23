package app

import (
	"fmt"
)

type IArea interface {
	GetSpace()
}



type Area struct {
	Space *Space
}

func NewArea() *Area {

	area := &Area{
		Space: &Space{MinX: 0, MaxX:10, MinY:0, MaxY:10},
	}

	return area
}

func (area *Area) GetSpace(){
	fmt.Printf("Space: %v ---------\n", area.Space)
}