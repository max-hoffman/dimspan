package main

func main() {}

// requirements:

// 1. pull in and format data into array structures from text file

// 2. form the hebian matrix, given offset and row #

// 3. perform SVD on that hebian matrix
// - need to pull in the gonum library to do SVD

// 4. pull the top n rows from SVD V* matrix for plotting
// - need gonum plotting library
// - will probably need context -> renderer -> texture maybe
// - sdl as a backup, but will be harder
// - worst case scenario just make it with D3 or React

// 5. find out the SINDy algorithm, make it its own library
// - @param : n rows from SVD (in our case)
// - @param : polynomial order (n)
// - @param : sin and cos ? (bool)
// - @return : equations representing underlying dist (form TBD, but could be string, array, or {}interface)

// 6. clean up and document before moving to real neuron data
