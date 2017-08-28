package main

import (
	"fmt"
	"log"
)

func main() {
	rawLorenz := createLorenzData()
	formattedLorenz := formatPlotData(rawLorenz)
	err := createSVG(formattedLorenz, "Lorenz sample", "lorenz.png", "Y", "Z")
	if err != nil {
		log.Fatalf("Create sample plot failed with error: %v", err)
	}

	noisyLorenz := addNoise(rawLorenz, 1)
	formattedNoisyLorenz := formatPlotData(noisyLorenz)
	err = createSVG(formattedNoisyLorenz, "Noisy Lorenz sample", "noisy-lorenz.png", "Y", "Z")
	if err != nil {
		log.Fatalf("Create sample plot failed with error: %v", err)
	}

	const (
		rowLength   = 1000
		vectorCount = 3
		rows        = 10
		dt          = .01
	)

	var singleVarStream []float64
	for _, tuple := range noisyLorenz {
		singleVarStream = append(singleVarStream, tuple[1])
	}

	// s, u, v, err := henkelSVD(singleVarStream, rowLength)
	s, _, v, err := henkelSVD(singleVarStream, rowLength, rows)
	if err != nil {
		log.Fatalf("Failed to perform SVD on data: %v\n", err)
	}

	var lorenzSVGData [][]float64
	rowCount, colCount := v.Dims()
	if vectorCount > colCount {
		log.Fatalf("Requested more dimensions than SVD returned: %v, %v\n", colCount, vectorCount)
	}
	for row := 0; row < rowCount; row++ {
		lorenzSVGData = append(lorenzSVGData, []float64{})
		currentRow := v.RowView(row)
		for col := 0; col < vectorCount; col++ {
			lorenzSVGData[row] = append(lorenzSVGData[row], currentRow.At(col, 0))
		}
	}

	formattedSVGLorenz := formatPlotData(lorenzSVGData)
	err = createSVG(formattedSVGLorenz, "Lorenz after SVD", "svd-lorenz.png", "Y", "Z")
	if err != nil {
		log.Fatalf("Create sample plot failed with error: %v", err)
	}

	fmt.Printf("s: %v\n", s)
	// fmt.Printf("u: %v\n", u)
	// fmt.Printf("%v\n", lorenzSVGData)

	dV := derivateMatrix(lorenzSVGData, dt)
	fmt.Print(len(dV))
}

// requirements:

// 1. create Lorenz data
// - write equation
// - get data for x
// - add noise

// 2. form the henkel matrix, given desired time offset
// - make sure to cut off excess data, or fill in with zeros

// 3. perform SVD on that hebian matrix
// - need to pull in the gonum library to do SVD

// 4. pull the top n rows from SVD V* matrix for plotting
// - need gonum plotting library
// - will probably need context -> renderer -> texture maybe
// - sdl as a backup, but will be harder
// - worst case scenario just make it with D3 or React

// -> I'm here
// 5. get derivatives of the n input rows
// - use "total-variation regularized derivative," need to figure out what that is

// 6. SINDy algorithm, make it its own library (reference paper here)
// - @param : n rows from SVD (in our case)
// - @param : polynomial order (n)
// - @param : sin and cos ? (bool)
// - @return : equations representing underlying dist (form TBD, but could be string, array, or {}interface)

// 6a. make necessary matrices
// - Theta(X) (m data x p candidate functions) amd Theta(xT) (1 x p)
// - Xi (p x n number of input streams)

// 6b. optimize Xi : (still need to figure out what this is doing)
// function Xi = sparsifyDynamics(Theta,dXdt,lambda,n)
// % Copyright 2015, All Rights Reserved
// % Code by Steven L. Brunton
// % For Paper, "Discovering Governing Equations from Data:
// %        Sparse Identification of Nonlinear Dynamical Systems"
// % by S. L. Brunton, J. L. Proctor, and J. N. Kutz

// % compute Sparse regression: sequential least squares
// Xi = Theta\dXdt;  % initial guess: Least-squares

// % lambda is our sparsification knob.
// for k=1:10
//     smallinds = (abs(Xi)<lambda);   % find small coefficients
//     Xi(smallinds)=0;                % and threshold
//     for ind = 1:n                   % n is state dimension
//         biginds = ~smallinds(:,ind);
//         % Regress dynamics onto remaining terms to find sparse Xi
//         Xi(biginds,ind) = Theta(:,biginds)\dXdt(:,ind);
//     end
// end

// 6c. Fnction to print optimization result as string (sparseGalerkin)
// 6. clean up and document before moving to neuron tests
