package engine

import "math/rand"

type point struct {
	x int
	y int
}

type rectangle struct {
	x     int //координата центра
	y     int //координата центра
	wedth int
	hight int
}

func (rect rectangle) LeftBottom() point {
	return point{rect.x - rect.wedth/2, rect.y - rect.hight/2}
}
func (rect rectangle) RightBottom() point {
	return point{rect.x + rect.wedth/2, rect.y - rect.hight/2}
}
func (rect rectangle) RightTop() point {
	return point{rect.x + rect.wedth/2, rect.y + rect.hight/2}
}
func (rect rectangle) LeftTop() point {
	return point{rect.x - rect.wedth/2, rect.y + rect.hight/2}
}

//мутирует: сдвигает значение по x  b  y на -2:+2
func (rect *rectangle) mutate() {
	rect.x += 2 - rand.Intn(5)
	rect.y += 2 - rand.Intn(5)
}
