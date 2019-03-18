package main
//TODO: make this a goroutine
func pairing(input []float32) []float32{
	//Sums and subtraction array, later we append subtract to the sums
	sums := make([]float32,0, len(input)/2)
	subtr := make([]float32,0, len(input)/2)
	//Input should always be just a row
	for i := 1; i < len(input); i++ {
		//Calculate averages and differences
		if !(i%2 == 0) {
			sums = append(sums, (input[i-1]+input[i])/2)
			subtr = append(subtr, (input[i-1]-input[i])/2)
		}
	}
	return append(sums,subtr...)
}
func blocks(input []byte) [][]float32{
	//Generate 8x8 blocks out of the input array
	return nil
}
