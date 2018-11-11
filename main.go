package main

import (
	"fmt"
	"isingo/lattice"
)

func main() {
	lat := &lattice.Lattice{}
	lat.Initialize()
	fmt.Println("normal:", lat.SimpleCalcEnergy())
}
