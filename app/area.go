package app

import (
	"fmt"

	"github.com/RuiHirano/simframe/api"
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
	*api.Area
}

func NewArea() *Area {

	area := &Area{
		&api.Area{
			Space: &Space{MinX: 0, MaxX:10, MinY:0, MaxY:10},
		},
	}

	return area
}

func (area *Area) GetSpace(){
	fmt.Printf("Space: %v ---------\n", area.Space)
}