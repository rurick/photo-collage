package engine

//Population - популяция
type Population struct {
	Cnt      int //количество особей
	Canvases []Canvas
}

//Init - создает новую случайную поппуляцию канвасов
func (p *Population) Init(cnt int, sizes [...]Size) {
	p.Cnt = cnt
	for i := 0; i < cnt; i++ {
		canv = new Canvas()
		canv.Init(sizes)
	}
}
