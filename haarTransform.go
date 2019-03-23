package main

import "fmt"

//TODO: make this a goroutine

/*
Takes an input row, calculates haar transform on it, depth elements just get appended without calculation
We send rows and columns alike as a 1D array, and return a 1D array which is the haar transform of it
We only append, without calculating from depth onwards
First run: depth = 0-> don't append anything,
Second run: depth = 4 -> append last 4 elements without calculaitng anything
Third run depth = 6, append last 6 elements without calculating anything
number of runs is decided by log2(8)
*/
func haar(input []float32, thr float32, depth int) []float32 {
	//Sums and subtraction array, later we append subtract to the sums
	sums := make([]float32, 0, len(input)/2)
	subtr := make([]float32, 0, len(input)/2)
	//Input should always be just a row
	//So we're expecting 8 x of 8 sized rows to be inputted N times
	for i := 1; i < len(input)-depth; i++ {
		//Calculate averages and differences
		if !(i%2 == 0) {
			var sum = (input[i-1] + input[i]) / 2
			var sub = (input[i-1] - input[i]) / 2
			if sum < thr {
				sum = 0
			}
			if sub < thr {
				sub = 0
			}
			sums = append(sums, sum)
			subtr = append(subtr, sub)
		}
	}

	subtr = append(subtr, input[len(input)-depth:]...)
	return append(sums, subtr...)
}

//Get all the pixels and width and height of the picture(x), turn it into 8x8 blocks and perform haar transform on the rows and columns
func blocks(pixels [][]float32, x, y int) [][]float32 {
	block := make([][]float32, 8)
	for i := 0; i < 8; i++ {
		block[i] = make([]float32, 8)
	}
	for i := 0; i < y; i += 8 {
		//Transform 8x8 block by transforming all rows, then transforming all columns
		for j := 0; j < x; j += 8 {
			//get 8x8 block
			block = pixels[j : j+8]
		}
	}
	fmt.Println("")
	haar(block[0], -151, 0)

	return block
}
