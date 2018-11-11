/*
simple

愚直な計算関数たち
 */
package lattice

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

//func (lat *Lattice) PairwiseLoopSum(function func(x, y int) (float64)) (sum float64) {
//	sum = 0.0
//	for x := 0; x < L; x++ {
//		for y := x; y < L; y++ {
//			sum += function(x, y)
//		}
//	}
//	return
//}
//
//func (lat *Lattice) CalcEnergy() (energy float64) {
//	energy = lat.PairwiseLoopSum(lat.SpinEnergy)
//	return
//}
