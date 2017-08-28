package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
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
func poolData(data [][]float64, n, polyorder int, usesine bool) *mat64.Dense {
	rowCount := len(data)
	colCount := 1 + n + (n * (n + 1) / 2) + (n * (n + 1) * (n + 2) / (2 * 3)) + 11
	theta := mat64.NewDense(rowCount, colCount, nil)
	colIdx := 0

	// zero order
	for row := range data {
		theta.Set(row, colIdx, 1)
	}
	colIdx++

	// first order
	for ridx, row := range data {
		for cidx, val := range row {
			theta.Set(ridx, colIdx+cidx, val)
		}
	}
	colIdx += 3

	// second order
	if polyorder >= 2 {
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				for ridx, row := range data {
					theta.Set(ridx, colIdx, row[i]*row[j])
				}
				colIdx++
			}
		}
	}

	// third order
	if polyorder >= 3 {
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				for k := j; k < n; k++ {
					for ridx, row := range data {
						theta.Set(ridx, colIdx, row[i]*row[j]*row[k])
					}
					colIdx++
				}
			}
		}
	}

	// fourth order
	if polyorder >= 4 {
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				for k := j; k < n; k++ {
					for l := k; l < n; l++ {
						for ridx, row := range data {
							theta.Set(ridx, colIdx, row[i]*row[j]*row[k]*row[l])
						}
						colIdx++
					}
				}
			}
		}
	}

	// fith order
	if polyorder >= 5 {
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				for k := j; k < n; k++ {
					for l := k; l < n; l++ {
						for m := l; m < n; m++ {
							for ridx, row := range data {
								theta.Set(ridx, colIdx, row[i]*row[j]*row[k]*row[l]*row[m])
							}
							colIdx++
						}
					}
				}
			}
		}
	}

	//sines
	// if usesine {

	// }

	return theta
}

// func normalize(data [][]float64) [][]float64 {

// }

// params: Theta, dx, lambda and n
// returns: Xi
// func PLS() {

// }

// func stringifyXi() (solution string) {

// }
