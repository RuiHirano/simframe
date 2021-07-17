package scenario

import (
	"fmt"
)

type IArea interface {
	GetSpace()
}

type Space struct{
	MinX float64
	MaxX float64
	MinY float64
	MaxY float64
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