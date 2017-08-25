package main

import (
	"fmt"

	"github.com/gonum/matrix"
	"github.com/gonum/matrix/mat64"
)

func henkelSVD(data []float64, delta, rows int) (s []float64, u, v *mat64.Dense, err error) {
	// for i := len(data); i < len(data)+len(data)%delta; i++ {
	// 	data = append(data, 0)
	// }
	// rows := len(data) / delta

	var henkel []float64
	for i := 0; i < rows; i++ {
		for j := i; j < delta+i; j++ {
			henkel = append(henkel, data[j])
		}
	}

	var svd mat64.SVD
	ok := svd.Factorize(mat64.NewDense(rows, delta, henkel), matrix.SVDThin)
	if !ok {
		s, u, v, err = nil, nil, nil, fmt.Errorf("failed to perform SVD on the input data")
		return
	}

	s, u, v = extractSVD(&svd)
	err = nil
	return
}

func extractSVD(svd *mat64.SVD) (s []float64, u, v *mat64.Dense) {
	var um, vm mat64.Dense
	um.UFromSVD(svd)
	vm.VFromSVD(svd)
	s = svd.Values(nil)
	return s, &um, &vm
}
