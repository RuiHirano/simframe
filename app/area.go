package app

import (
)

type IArea interface {
	GetSpace()*Space
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

func (area *Area) GetSpace()*Space{
	return area.Space
}