/*
Lattice module

モデルの定義とutil
*/

package lattice

import "math/rand"

// エッジサイズ
const L = 512

// 磁化
const Mag = -1.0

// 本体
type Lattice struct {
	Spin [L][L]float64 `label:"スピンの値"`
}

// Loop 単純な要素の走査
func (lat *Lattice) Loop(function func(x, y int)) {
	for x := 0; x < L; x++ {
		for y := 0; y < L; y++ {
			function(x, y)
		}
	}
}

func (lat *Lattice) RandomizeSpin(x, y int) {
	lat.Spin[x][y] = float64(2*rand.Intn(2)) - 1.0
}

// ランダムに初期化する
func (lat *Lattice) Initialize() {
	lat.Loop(lat.RandomizeSpin)
}

func (lat *Lattice) Fill() {
	lat.Loop(func(x, y int) { lat.Spin[x][y] = 1 })
}

// あるスピン1つのエネルギー
func (lat *Lattice) SpinEnergy(x, y int) (energy float64) {
	nears := lat.Spin[x][(y-1+L)%L] +
		lat.Spin[x][(y+1)%L] +
		lat.Spin[(x-1+L)%L][y] +
		lat.Spin[(x+1)%L][y]
	energy = Mag * lat.Spin[x][y] * nears
	return
}
