package main

import (
	"fmt"

	ga "./engine"
)

func main() {
	fmt.Println(ga.Max(4, 5))
	var canv ga.Canvas
	canv.CalcSize()
}
