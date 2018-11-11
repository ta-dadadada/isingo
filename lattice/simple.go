/*
simple

愚直な計算関数たち
 */
package lattice

import (
	"math"
	"math/rand"
)

// 格子点を1つづつ走査して、関数を適用した結果を足し上げて返す
func (lat *Lattice) LoopSum(function func(x, y int) float64) (sum float64) {
	sum = 0.0
	for x := 0; x < L; x++ {
		for y := 0; y < L; y++ {
			sum += function(x, y)
		}
	}
	return
}

// 愚直に全エネルギーを計算する
// ペアごとに2回足してしまうのでエネルギーは2で割る
func (lat *Lattice) SimpleCalcEnergy() (energy float64) {
	energy = lat.LoopSum(lat.SpinEnergy) / 2.0
	return
}

// 格子の個数分だけランダムに何かをする
// すべてのスピンを走査するわけではない
func (lat *Lattice) RandomLoopSum(function func(x, y int)) (sum float64) {
	for i := 0; i < L*L; i++ {
		x, y := rand.Intn(L), rand.Intn(L)
		function(x, y)
	}
	return
}

// メトロポリス法でスピンを1つ更新
func (lat *Lattice) Mutate(x, y int) (energyDiff float64) {
	before := lat.SpinEnergy(x, y)
	lat.Flip(x, y)
	after := lat.SpinEnergy(x, y)
	diff := after - before
	if diff <= 0.0 || math.Exp(-diff/lat.Temperature) > rand.Float64() {
		energyDiff = diff
	} else {
		lat.Flip(x, y)
		energyDiff = 0.0
	}
	return
}


// ランダムにスピンの数だけMutate
func (lat *Lattice) RandomUpdate(energy float64, temperature float64) (newEnergy float64) {
	lat.Temperature = temperature
	lat.RandomLoopSum(func(x, y int) {
		energy = energy + lat.Mutate(x, y)
	})
	newEnergy = energy
	return
}
