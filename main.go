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

func watchEnergyCalculation(iterations int) {
	lat := &lattice.Lattice{}
	lat.Fill()
	fmt.Printf("normal: ")
	stopwatch(lat.SimpleCalcEnergy, iterations)
	fmt.Printf("checker board:")
	stopwatch(lat.CalcEnergyCB, iterations)
	fmt.Printf("bakapara:")
	stopwatch(lat.CalcEnergyBakapara, iterations)
	fmt.Printf("checker board para:")
	stopwatch(lat.CalcEnergyCBGoroutine, iterations)
}

func watchUpdate(temperature float64) {
	lat := &lattice.Lattice{}
	lat.Initialize()
	energy := lat.SimpleCalcEnergy()
	fmt.Printf("E=%v\n", energy)
	do := func(update func(float64, float64) float64, iters int) {
		start := time.Now()
		for i:=0; i<iters;i++ {
			update(energy, temperature)
		}
		fmt.Printf("time=%v\n", time.Now().Sub(start))
	}
	fmt.Printf("random: ")
	do(lat.RandomUpdate, 10)
	fmt.Println(lat.SimpleCalcEnergy())
}

func main() {
	watchUpdate(1.0)
}
