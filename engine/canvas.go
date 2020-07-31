package engine

import (
	"math/rand"
	"time"
)

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

//Canvas - холст с расположеными на нём прямоугольниками
type Canvas struct {
	rectangles []Rectangle
	Square     int
	minX       int
	minY       int
	maxX       int
	maxY       int
}

func (canv *Canvas) Init(sizes [...]Size) {
	for _, size := range sizes {
		canv.Add(size.w, size.h)
	}
	canv.CalcSize()
	canv.CalcSquare()
}

//CalcSize - рассчитать верхний и нихний угол канваса
func (canv *Canvas) CalcSize() {
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
func (canv *Canvas) CalcSquare() int {
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
func (canv *Canvas) Add(width int, height int) bool {
	vercots := [...][2]int{ //направления по который пробуем добавлять новый прямоугольник
		[2]int{1, 0},
		[2]int{-1, 0},
		[2]int{0, 1},
		[2]int{0, -1},
		[2]int{1, 1},
		[2]int{-1, -1},
		[2]int{-1, 1},
		[2]int{1, -1}}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(vercots), func(i, j int) { vercots[i], vectors[j] = vercots[j], vectors[i] })
	centerX := (canv.maxX - canv.minX) / 2
	w := canv.maxX - canv.minX
	centerY := (canv.maxY - canv.minY) / 2
	h := canv.maxY - canv.minY
	for _, vector := range vercots {
		if !canv.IsBusy(
			centerX+vector[0]*w/2+width/2+1,
			centerY+vector[1]*h/2+height/2+1,
			width,
			height) {
			//свободно
			canv.rectangles = append(canv.rectangles, Rectangle{
				vector[0]*canv.maxX + width/2 + 1,
				vector[1]*canv.maxY + height/2 + 1,
				height,
				width})
			return true
		}
	}
	return false
}

//IsBusy - провериить влазит ли сюда прямоугольник
func (canv *Canvas) IsBusy(x int, y int, width int, height int) bool {
	var (
		x1 = x - width/2
		y1 = y - height/2
		x2 = x + width/2
		y2 = y + height/2
		x3 = x - width/2
		y3 = y + height/2
		x4 = x + width/2
		y4 = y - height/2
	)
	for _, rec := range canv.rectangles {
		lb := rec.LeftBottom()
		rt := rec.RightTop()
		if lb.x < x1 && x1 < rt.x && lb.y < y1 && y1 < rt.y {
			return true
		}
		if lb.x < x2 && x2 < rt.x && lb.y < y2 && y2 < rt.y {
			return true
		}
		if lb.x < x3 && x3 < rt.x && lb.y < y3 && y3 < rt.y {
			return true
		}
		if lb.x < x4 && x4 < rt.x && lb.y < y4 && y4 < rt.y {
			return true
		}
	}
	return false
}
