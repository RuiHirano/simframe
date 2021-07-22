package api

func NewArea() *Area {

	area := &Area{
		Space: &Space{MinX: 0, MaxX:10, MinY:0, MaxY:10},
	}

	return area
}