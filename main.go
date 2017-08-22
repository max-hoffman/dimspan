package main

import (
	"log"
)

func main() {
	rawLorenz := createLorenzData()
	formattedLorenz := formatPlotData(rawLorenz)
	err := createSVG(formattedLorenz, "Lorenz sample", "Y", "Z")
	if err != nil {
		log.Fatalf("Create sample plot failed with error: %v", err)
	}
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

// 6. clean up and document before moving to neuron tests
