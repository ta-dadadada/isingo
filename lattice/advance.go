package lattice

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

