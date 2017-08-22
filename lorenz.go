package main

import (
	"math/rand"
	"time"

	"github.com/sj14/ode"
)

func lorenz(t float64, y []float64) []float64 {
	sigma := float64(10)
	beta := float64(8 / 3)
	rho := float64(28)

	result := make([]float64, 3)
	result[0] = sigma * (y[1] - y[0])
	result[1] = y[0]*(rho-y[2]) - y[1]
	result[2] = y[0]*y[1] - beta*y[2]

	return result
}

func createLorenzData() [][]float64 {
	initCond := []float64{-8, 8, 27}

	y := ode.RungeKutta4(.001, .001, 100, initCond, lorenz)

	// for _, val := range y {
	// 	fmt.Println(val)
	// }

	return y
}

func addNoise(data [][]float64, value int) (noisyData [][]float64) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range data {
		noisyData[i][1] += float64(rand.Intn(value))
		noisyData[i][2] += float64(rand.Intn(value))
		noisyData[i][3] += float64(rand.Intn(value))
	}
	return noisyData
}
