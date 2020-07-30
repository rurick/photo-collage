package engine

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

//MaxInt32 - максимальное целое 32 бит
const MaxInt32 = 1<<31 - 1

//canvas - холст с расположеными на нём прямоугольниками
type canvas struct {
	rectangles []rectangle
	Square     int
	minX       int
	minY       int
	maxX       int
	maxY       int
}

//CalcSize - рассчитать верхний и нихний угол канваса
func (canv *canvas) CalcSize() {
	canv.minX, canv.minY, canv.maxX, canv.maxY = MaxInt32, MaxInt32, 0, 0
	for _, rec := range canv.rectangles {
		lb := rec.LeftBottom()
		canv.minX = min(canv.minX, lb.x)
		canv.minY = min(canv.minY, lb.y)
		rt := rec.RightTop()
		canv.maxX = max(canv.maxX, rt.x)
		canv.maxY = max(canv.maxY, rt.y)
	}
}

//CalcSquare - рассчитать площадь канваса
func (canv *canvas) CalcSquare() int {
	var minX, minY, maxX, maxY int = MaxInt32, MaxInt32, 0, 0
	for _, rec := range canv.rectangles {
		lb := rec.LeftBottom()
		minX = min(minX, lb.x)
		minY = min(minY, lb.y)
		rt := rec.RightTop()
		maxX = max(maxX, rt.x)
		maxY = max(maxY, rt.y)
	}
	canv.Square = (maxX - minX) * (maxY - minY)
	return canv.Square
}

//Add - Добавить новый прямоугольник на канвас
// возвращает true  если успешно добавил
func (canv *canvas) Add(width int, height int) bool {
	return false
}

//IsBusy - провериить влазит ли сюда прямоугольник
func (canv *canvas) IsBusy(x int, y int, width int, height int) bool {
	return false
}
