package engine

//Start - инициализирует первую поппуляцию
func Start(poppulatioCnt int) *Population {
	var pop = new(Population)
	pop.Init(1, [...]Size{
		Size{150, 110},
		Size{50, 80},
		Size{120, 100},
		Size{140, 110}})
}
