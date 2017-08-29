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

// param: svd V basis
// returns: dx
func derivate(data [][]float64, dt float64, n int) *mat64.Dense {
	// dV(i-2,k) = (1/(12*dt))*(-V(i+2,k)+8*V(i+1,k)-8*V(i-1,k)+V(i-2,k))
	rowCount := len(data)
	dV := mat64.NewDense(rowCount, n, nil)
	for r := 0; r < rowCount; r++ {
		for c := 0; c < n; c++ {
			if r < 2 || r >= rowCount {
				dV.Set(r, c, 0)
				continue
			}
			deltaVal := (1 / (12 * dt)) * (-data[r+2][c] + 8*data[r+1][c] - 8 - data[r-1][c] + data[r-2][c])
			dV.Set(r, c, deltaVal)
		}
	}

	// fmt.Printf("rows in dV: %v\n", len(dV))
	// fmt.Printf("cols in dV: %v\n", len(dV[0]))
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
func pls(dx, theta *mat64.Dense) (*mat64.Dense, error) {
	psi := mat64.NewDense(20, 3, nil)

	//init psi with Solve(theta, dx)
	err := psi.Solve(dx, theta)
	if err != nil {
		return nil, fmt.Errorf("failed to perform least-squares regression with error: %v", err)
	}

	for i := 0; i < 10; i++ {
		// find small indices

		// set them equal to zero

		// perform solve on the remaining big indices
	}
	return psi, nil
}

// func stringifyXi() (solution string) {

// }
