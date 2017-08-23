package main

import (
	"fmt"

	"github.com/gonum/matrix"
	"github.com/gonum/matrix/mat64"
)

func henkelSVD(data []float64, delta int) (s []float64, u, v *mat64.Dense, err error) {
	for i := len(data); i < len(data)+len(data)%delta; i++ {
		data = append(data, 0)
	}
	rows := len(data) / delta

	var svd mat64.SVD
	ok := svd.Factorize(mat64.NewDense(rows, delta, data), matrix.SVDThin)
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
