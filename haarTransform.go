package main

import "fmt"

//TODO: make this a goroutine

/* Takes an input row, calculates haar transform on it, depth elements just
get appended without calculation We send rows and columns alike as
a 1D array, and return a 1D array
which is the haar transform of it We only append, without calculating from depth onwards
First run: depth = 0-> don't append anything,
Second run: depth = 4 -> append last 4 elements without calculaitng anything
Third run depth = 6, append last 6 elements without calculating anything
number of runs is decided by log2(8) -> we're dealing with 8 sized rows/columns
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
	rowhaar := append(sums, subtr...)
	fmt.Println("Recursing...", depth, rowhaar)
	//Unless we're at the last depth, recurse
	//Ugly as shit recursion, please FIXME
	switch depth {
	case 0:
		return haar(rowhaar, thr, 4)
	case 4:
		return haar(rowhaar, thr, 6)
	case 6:
		fmt.Println("Ending recursion... ", rowhaar)
		return rowhaar
	default:
		return nil
	}
	return nil
}

//Get all the pixels and width and height of the picture(x = height , y = width)
//Pass down threshhold(thr) from the user input
//turn it into 8x8 blocks and perform haar transform on the rows and columns
func blocks(pixels [][]float32, x, y int, thr float32) [][]float32 {
	block := make([][]float32, 8, 8)
	//The transformed 8x8 block
	transformed := make([][]float32, x, y)
	for i := 0; i < 8; i++ {
		block[i] = make([]float32, 8, 8)
		transformed[i] = make([]float32, 8, 8)
	}
	//Transform 8x8 block by transforming all rows, then transforming all columns
	for i := 0; i < x; i++ {
		//get 8xY sized block
		//Get [0,0], [8,8],[8,16] etc. every 8th tile in the 2D array
		for j := 0; j < y; j++ {
			block[i%8][j%8] = pixels[i][j]
			//TODO: Haar transform on block row

		}

	}
	fmt.Println("Done with blocks, ", transformed[0])
	return block
}

//Receive 8x8 block, return 1D array of size 8, based on column value
//index 0 -> get row 0 in 8x8 block
func getRow(block [][]float32, index int) []float32 {
	row := block[index][0:8]
	return row
}
