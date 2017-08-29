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
func pls(dx, theta *mat64.Dense, lambda float64) (*mat64.Dense, error) {
	m, n := dx.Dims()
	_, p := theta.Dims()
	xi := mat64.NewDense(p, n, nil)
	var tempXi mat64.Dense

	//initial guess for xi
	err := xi.Solve(dx, theta)
	if err != nil {
		return nil, fmt.Errorf("least-squares regression failed with error: %v", err)
	}

	// optimize
	for col := 0; col < n; col++ {
		for i := 0; i < 10; i++ {

			// find small indices, set them equal to zero
			bigIdx := []int{}
			for row := 0; row < p; row++ {
				val := xi.At(row, col)
				if val < lambda && val > -lambda {
					xi.Set(row, col, 0)
				}
				bigIdx = append(bigIdx, col)
			}

			// collect theta columns that were large in xi
			tempTheta := mat64.NewDense(m, len(bigIdx), nil)
			for i, colIdx := range bigIdx {
				newCol := getRawCol(theta, colIdx)
				tempTheta.SetCol(i, newCol)
			}

			// get new approximations for xi
			err := tempXi.Solve(dx, tempTheta)
			if err != nil {
				return nil, fmt.Errorf("least-squares regression failed in cycle %v with error: %v", i, err)
			}

			// replace updated approximations for xi
			// MARK : do the indexes still match up with the correct rows?
			for idx, val := range bigIdx {
				xi.SetRow(val, tempXi.RawRowView(idx))
			}
		}
	}
	return xi, nil
}

// func stringifyXi() (solution string) {

// }

func getRawCol(m *mat64.Dense, col int) (newCol []float64) {
	rowCount, _ := m.Dims()
	for row := 0; row < rowCount; row++ {
		newCol = append(newCol, m.At(row, col))
	}
	return
}

func regularize(i, j int, v float64) float64 {
	var lambda float64
	switch j {
	case 0:
		lambda = .01
	default:
		lambda = .2
	}
	if v < lambda && v > -lambda {
		return 0
	}
	return v
}
