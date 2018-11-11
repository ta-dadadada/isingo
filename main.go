package main

import (
	"fmt"
	"isingo/lattice"
	"time"
)

// エネルギー計算の時間計測
func stopwatch(function func() float64, iters int) {
	start := time.Now()
	res := 0.0
	for i := 0; i < iters; i++ {
		res = function()
	}
	fmt.Printf("result=%v time=%v\n", res, time.Now().Sub(start))
}

func main() {
	lat := &lattice.Lattice{}
	lat.Fill()
	iterations := 1000
	fmt.Printf("normal: ")
	stopwatch(lat.SimpleCalcEnergy, iterations)
	fmt.Printf("checker board:")
	stopwatch(lat.CalcEnergyCB, iterations)
	fmt.Printf("bakapara:")
	stopwatch(lat.CalcEnergyBakapara, iterations)
	fmt.Printf("checker board para:")
	stopwatch(lat.CalcEnergyCBGoroutine, iterations)
}
