package main

import (
	"math"
)

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
		sqrt := float32(math.Sqrt(2))
		if !(i%2 == 0) {
			var sum = (input[i-1] + input[i]) / 2
			//Round to the 4th decimal
			sum *= sqrt
			var sub = (input[i-1] - input[i]) / 2
			sub *= sqrt
			if sum < thr {
				sum = 0
			}
			if sub < thr {
				sub = 0
			}
			//Round to 4 decimals
			sums = append(sums, float32(math.Round(float64(sum*10000))/10000))
			subtr = append(subtr, float32(math.Round(float64(sub*10000)/10000)))
		}
	}

	subtr = append(subtr, input[len(input)-depth:]...)
	rowhaar := append(sums, subtr...)
	//Unless we're at the last depth, recurse
	//Ugly as shit recursion, please FIXME
	switch depth {
	case 0:
		return haar(rowhaar, thr, 4)
	case 4:
		return haar(rowhaar, thr, 6)
	case 6:
		return rowhaar
	default:
		return nil
	}
}

//Get all the pixels and width and height of the picture(x = height , y = width)
//We can expect heigh and width to be identical in our future test cases
//Pass down threshhold(thr) from the user input
//turn it into 8x8 blocks and perform haar transform on the rows and columns
//TODO: Insert the transformed blocks where it's supposed to be
func blocks(pixels [][]float32, x, y int, thr float32) []float32 {
	//Count how many 8x8 blocks there can be in an x times y matrix
	count := (x * y) / 64
	blocks := make([][][]float32, count)
	for i := range blocks {
		blocks[i] = make([][]float32, 8)
		for j := range blocks[i] {
			blocks[i][j] = make([]float32, 8)
		}
	}
	z := 0
	//Extra iteration counter, so we can keep resetting m
	m := 0
	//TODO: Stop when reaching the end
	for d := 0; d < len(blocks)-1; d++ {
		//A single 8x8 block
		for i := z * 8; i/8 < z+1; i++ {
			//f and g are in-block trackers
			f := 0
			for j := m * 8; j/8 < m+1; j++ {
				g := 0
				currPixel := pixels[i][j]
				blocks[d][f][g] = currPixel
				g++
			}
			f++
		}
		blocks[d] = blocksT(blocks[d], thr)
		m++
		//Reached right corner
		//Restart algorithm, one row down
		if x/8 == m {
			z++
			m = 0

		}
		//The transformed 8x8 block
		//transformed := make([][]float32, x, y)
		//Transform 8x8 block by transforming all rows, then transforming all columns
		//get 8xY sized block
		//Get [0,0], [8,8],[8,16] etc. every 8th tile in the 2D array

		//build a matrix of 8x8 blocks
		//Transform H into orthogonal matrix-> Inverse is faster
		//Normalize each colmn of the starting matrix to length 1
	}
	field := zigZag(blocks)
	return field
}

//Receive 8x8 block, return 1D array of size 8, based on index value
//index 0 -> get row 0 in 8x8 block
func getRow(block [][]float32, index int) []float32 {
	row := block[index][:]
	return row
}

//Receive 8x8 block, return 1D array of size 8, representing the indexth column
func getColumn(block [][]float32, index int) []float32 {
	column := make([]float32, 8)
	//alternative := block[:][index]
	for i := range column {
		column[i] = block[i][index]
	}
	return column
}

//Transform the received block
//TODO: Can use goroutine for this
func blocksT(block [][]float32, thr float32) [][]float32 {
	transformedBlock := make([][]float32, 8, 8)
	for i := 0; i < 8; i++ {
		transformedBlock[i] = haar(getRow(block, i), thr, 0)
	}
	//Transform blocks after, used the already transformed matrix...
	for i := 0; i < 8; i++ {
		//Get column as a row-> insert it as a column
		currColumn := getColumn(transformedBlock, i)
		for j := 0; j < 8; j++ {
			transformedBlock[j][i] = haar(currColumn, thr, 0)[j]
		}
	}
	//Transpose it because I have low iq
	return transformedBlock
}

//Global zig zag, takes in an array of 8x8 blocks, then zig zags all of them simultaneously, offsetting the indexes by each matrix index
//arrayOfBlocks[blockIndex][i][j] -> result[blockIndex + iterator]
func zigZag(block [][][]float32) []float32 {
	//Resulting a NxN length 1D matrix from an 8x8 block
	l := len(block)
	result := make([]float32, l*64)
	for z := 0; z < l; z++ {
		i := 1
		j := 0
		iterator := 2 * l
		result[z] = block[z][0][0]
		result[z+l] = block[z][0][1]
		//Last one is predetermined
		result[(63*l)+z] = block[z][7][7]
		firsthalf := true
		for iterator+z < 64*l {
			if i == 8 && j == 0 {
				i = 7
				j++
				firsthalf = !firsthalf
			} else if i == 8 && j == 7 {
				i = 7
			}
			if firsthalf {
				if j%2 != 0 || i%2 != 0 {
					for j != 0 {
						result[iterator+z] = block[z][i][j]
						iterator += l
						i++
						j--
					}
					result[iterator+z] = block[z][i][j]
					iterator += l
					i++
				} else if i%2 == 0 {
					for i != 0 {
						result[iterator+z] = block[z][i][j]
						iterator += l
						i--
						j++

					}
					result[iterator+z] = block[z][i][j]
					iterator += l
					j++
				}
				//Second half of the matrix
			} else {
				if j%2 != 0 && i == 7 {
					//Go up the matrix until you reach the inverse(from 5,0 -> 0,5)
					for j != 7 {
						result[iterator+z] = block[z][i][j]
						iterator += l
						i--
						j++

					}
					result[iterator+z] = block[z][i][j]
					iterator += l
					i++
				} else if i%2 == 0 {
					for i != 7 {
						result[iterator+z] = block[z][i][j]
						iterator += l
						j--
						i++
					}
					result[iterator+z] = block[z][i][j]
					iterator += l
					j++
				}
			}
		}
	}

	return result
}
