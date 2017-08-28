package main

import (
	"fmt"
)

// param: svd V basis
// returns: dx
func derivateMatrix(data [][]float64, dt float64) [][]float64 {
	// dV(i-2,k) = (1/(12*dt))*(-V(i+2,k)+8*V(i+1,k)-8*V(i-1,k)+V(i-2,k))
	var dV [][]float64
	for ridx, row := range data {
		if ridx < 2 || ridx >= len(data)-2 {
			continue
		}
		dV = append(dV, []float64{})
		for cidx := range row {
			deltaVal := (1 / (12 * dt)) * (-data[ridx+2][cidx] + 8*data[ridx+1][cidx] - 8 - data[ridx-1][cidx] + data[ridx-2][cidx])
			dV[ridx-2] = append(dV[ridx-2], deltaVal)
		}
	}
	fmt.Printf("rows in dV: %v\n", len(dV))
	fmt.Printf("cols in dV: %v\n", len(dV[0]))
	return dV
}

// params: x, n, polyorder and usesine
// returns: Theta
func poolData() {

}

// params: Theta, dx, lambda and n
// returns: Xi
func sparseRegress() {

}

// func stringifyXi() (solution string) {

// }
