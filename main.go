package main

import (
	"fmt"
	"isingo/lattice"
)

func main() {
	lat := &lattice.Lattice{}
	lat.Initialize()
	lat.Fill()
	fmt.Println(lat.Spin)
	fmt.Println("normal:", lat.SimpleCalcEnergy())
//	fmt.Println("pairwise:", lat.CalcEnergy())
	fmt.Println("cheker board:", lat.CalcEnergyCB())
}
