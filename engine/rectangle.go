package engine

import "math/rand"

type point struct {
	x int
	y int
}

type Size struct {
	w int
	h int
}

type Rectangle struct {
	x     int //координата центра
	y     int //координата центра
	wedth int
	hight int
}

func (rect Rectangle) LeftBottom() point {
	return point{rect.x - rect.wedth/2, rect.y - rect.hight/2}
}
func (rect Rectangle) RightBottom() point {
	return point{rect.x + rect.wedth/2, rect.y - rect.hight/2}
}
func (rect Rectangle) RightTop() point {
	return point{rect.x + rect.wedth/2, rect.y + rect.hight/2}
}
func (rect Rectangle) LeftTop() point {
	return point{rect.x - rect.wedth/2, rect.y + rect.hight/2}
}

//Mutate: сдвигает значение по x  и  y на -2:+2
func (rect *Rectangle) Mutate() {
	rect.x += 2 - rand.Intn(5)
	rect.y += 2 - rand.Intn(5)
}
