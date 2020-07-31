package main

import (
	"fmt"

	"./engine"
)

func main() {
	fmt.Println(engine.Max(4, 5))
	engine.Start()
}
