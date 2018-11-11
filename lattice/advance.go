package lattice

import (
	"runtime"
	"sync"
)

// 添字の偶奇が同じケースのループ。「黒」とよぶことにする
func (lat *Lattice) LoopSumBlack(function func(x, y int) float64) (sum float64) {
	sum = 0.0
	for x := 0; x < L; x += 2 {
		for y := 0; y < L; y += 2 {
			sum += function(x, y)
		}
	}
	for x := 1; x < L; x += 2 {
		for y := 1; y < L; y += 2 {
			sum += function(x, y)
		}
	}
	return
}

// 添字の偶奇が異なるケースのループ。「白」とよぶことにする
func (lat *Lattice) LoopSumWhite(function func(x, y int) float64) (sum float64) {
	sum = 0.0
	for x := 0; x < L; x += 2 {
		for y := 1; y < L; y += 2 {
			sum += function(x, y)
		}
	}
	for x := 1; x < L; x += 2 {
		for y := 0; y < L; y += 2 {
			sum += function(x, y)
		}
	}
	return
}

// checkerboard の順でエネルギー総和を計算
func (lat *Lattice) CalcEnergyCB() (sum float64) {
	sum = 0.0
	sum += lat.LoopSumBlack(lat.SpinEnergy)
	sum += lat.LoopSumWhite(lat.SpinEnergy)
	sum /= 2.0
	return
}

// cpu数分ごとにgoroutineで並列にする
func (lat *Lattice) CalcEnergyBakapara() (sum float64) {
	sum = 0.0
	wg := sync.WaitGroup{}
	cpus := runtime.NumCPU()
	res := make(chan float64, cpus)
	for x := 0; x < L; x++ {
		for y := 0; y < L; y += cpus {
			threadNum := 0
			for ; (y+threadNum < L) && (threadNum < cpus); threadNum++ {
				wg.Add(1)
				go func(r chan<- float64) {
					defer wg.Done()
					r <- lat.SpinEnergy(x, y)
				}(res)
			}
			wg.Wait()
			for i := 0; i < threadNum; i++ {
				sum += <-res
			}
		}
	}
	sum /= 2
	return
}

// チェッカーボードの白黒をそれぞれgoroutineで並列にする
func (lat *Lattice) CalcEnergyCBGoroutine() (sum float64) {
	sum = 0.0
	res := make(chan float64, 2)
	go func(r chan<- float64) {
		r <- lat.LoopSumBlack(lat.SpinEnergy)
	}(res)
	go func(r chan<- float64) {
		r <- lat.LoopSumWhite(lat.SpinEnergy)
	}(res)
	for i := 0; i < 2; i++ {
		sum += <-res
	}
	sum /= 2.0
	return
}
