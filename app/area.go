package app

import (
)

type IArea interface {
	GetSpace()*Space
}

type Area struct {
	Space *Space
}

func NewArea(space *Space) IArea {

	area := &Area{
		Space: space,
	}

	return area
}

func (area *Area) GetSpace()*Space{
	return area.Space
}