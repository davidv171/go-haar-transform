package main

import (
	"sort"
)

type Decoder struct {
	/*
		- inputBits : the bool array of input bits, that gets shifted for error calculations
		- symbolInterval: is the calculated symbol interval, that falls between a symbol's high and low value in the model definition
		- high table is encoded in the file, low table is derived from the high table
		- step,low,high : step, calculation helper used here for data retention
		- output: the decoded byte array, decompressed file
		- numberOfAllSymbols: number of all, even non-unique symbols of the encoded file
		- currentInput: the field/array of bools, upon which we od operations on, next bit is taken out of inputBits
		- quarters: 4 sized array containing certain quarters used for error calculation
	*/
	inputBits          []bool
	highTable          []uint32
	lowTable           []uint32
	symbolInterval     uint32
	step               uint32
	low                uint32
	high               uint32
	output             []byte
	numberOfAllSymbols uint32
	currentInput       []bool
	quarters           []uint32
	index              uint32
}

//The first 256 32-bit bytes are our frequency table, or rather out HIGH table, containing the high value of each symbol
func (d *Decoder) readFreqTable(data []uint8) {
	highTable := make([]uint32, 256)
	for i := 4; i < 256*4; i += 4 {
		readBytes := data[i-4 : i]
		convertedInt := bytesToUint32(readBytes)
		highTable[(i-4)/4] = convertedInt

	}
	d.highTable = highTable
	//A temporary variable to avoid ovewriting
	hsorted := make([]uint32, 256)
	copy(hsorted, highTable)
	sort.Slice(hsorted, func(i, j int) bool { return hsorted[i] < hsorted[j] })
	//The largest high(last one in the sorted array) tells us how many non-unique symbols we have
	d.numberOfAllSymbols = hsorted[255]
	d.step = (d.high + 1) / d.numberOfAllSymbols
	d.generateLowTable(hsorted)
	lsorted := make([]uint32, 256)
	copy(lsorted, d.lowTable)
	sort.Slice(lsorted, func(i, j int) bool { return lsorted[i] < lsorted[j] })
	hmap := keep(highTable)
	d.initializeField(data)
	d.quarterize(d.high)
	d.intervalGeneration(hmap, lsorted, hsorted)
	//Sort the array so we can generate low table much easier
}

/*Generate low table out of the high table
1. Find the lowest high
2. The lowest high's symbol's low = 0
3. Find the second lowest high, that symbol's low is the previous symbol's high
* This could also be used for interval definition
*/
func (d *Decoder) generateLowTable(hsorted []uint32) {
	//Generate low table based on the sorted table and high table that is unsorted
	//Put the value of non-0 indexes of sorted table into the index corresponding to the value in hightable
	//IN other words, put the low table's values into the right
	highTable := d.highTable
	lowTable := make([]uint32, 256)
	//hsorted is sorted from lowest to highest
	//The first low value of a symbol is always 0
	//The next value of low is equal to the next high value
	for i := 0; i < 256; i++ {
		currSorted := hsorted[i]
		if currSorted != 0 {
			//Find the index with value in unsorted table, don't loop from start every time
			for j := 0; j < 256; j++ {
				if currSorted == highTable[j] {
					lowTable[j] = hsorted[i-1]
					break
				}
			}
		}
	}
	d.lowTable = lowTable

}

//Take the first 7 bits(or 32 bits, depending on the arithmetic coder size)
func (d *Decoder) initializeField(data []uint8) {
	var i uint32 = 0
	bitSlice := make([]bool, 0)
	//Load every bit, that is not in the model(first 256*4 bytes) onwards
	for i = 256 * 4; i < uint32(len(data)); i++ {
		currByte := data[i]
		bitSlice = append(bitSlice, byteToBitSlice(uint32(currByte), 8)...)
	}
	d.inputBits = bitSlice
	//First 32 bits
	d.currentInput = bitSlice[0:32]

}

/*
1. read probability table
2. intialize the array of first N bits
3. Decode on every iteration:
	v = (field - low) / step
	step = roundDown(high - low + 1) / n
	high = low + step * high(simbol) - 1
	low = low + step * low(simbol)
4. Check for errors on each iteration:
	while((mHigh < secondQuarter) || (mLow >= secondQuarter))
		- E1: high < secondQuarter
			- low = 2*low
			- high = 2*high + 1
			- field = 2 * field + next bit(0 or 1)
		- E2: low >= 2ndQaurter
			- low = 2*(low - 2ndQuarter)
			- high = 2*(high - 2ndQuarter) + 1
			- field = 2*(field - secondQuarter) + next bit
		- E3: (low>=firstQuarter && high < thirdQuarter)
			- low = 2*(low - firstQuarter)
			- high = 2*(high - firstQuarter)+1
			- field = 2*(field - 1stQuarter) + next bit
*/
func (d *Decoder) intervalGeneration(hmap map[uint32]int, lsorted []uint32, hsorted []uint32) {
	quarters := d.quarters
	numberOfAllSymbols := d.numberOfAllSymbols
	//An index to keep track of getting the value of the next bit from input bits and adding it to the symbol
	index := d.index
	inputBits := d.inputBits
	for i := 256 * 4; uint32(i) < 1024+d.numberOfAllSymbols; i++ {
		low := d.low
		high := d.high
		step := d.step
		currentBits := d.currentInput
		currentByte := arbitraryBitsToByte(&currentBits)
		step = uint32((uint64(high) - uint64(low) + 1) / uint64(numberOfAllSymbols))
		symbolInterval := (currentByte - low) / step
		//Calculate the symbol that relates to the symbolInterval
		symbol := d.intervalToSymbol(symbolInterval, hmap, lsorted, hsorted)
		//Each decoded symbol is added to then be written into the file
		d.output = append(d.output, symbol)
		high = low + step*d.highTable[symbol] - 1
		low = low + step*d.lowTable[symbol]
		//Error intervals
		for (high < quarters[1]) || low >= quarters[1] {
			//E1
			if high < quarters[1] {
				low = 2 * low
				high = 2*high + 1
				//Turn bool into 0 or 1 then add it
				//TODO: Turn it into a function
				var add uint8 = 0
				if index < uint32(len(inputBits)) {
					if inputBits[index] {
						add = 1
					}
				}
				tempByte := currentByte
				tempByte = 2*tempByte + uint32(add)
				currentByte = tempByte
				index++

			} else if low >= quarters[1] {
				//E2
				low = 2 * (low - quarters[1])
				high = 2*(high-quarters[1]) + 1
				var add uint8 = 0
				if index < uint32(len(inputBits)) {
					if inputBits[index] {
						add = 1
					}
				}
				tempByte := currentByte
				tempByte = 2*(uint32(tempByte)-quarters[1]) + uint32(add)
				currentByte = tempByte
				index++
			}

		}
		for (low >= quarters[0]) && high < quarters[2] {
			low = 2 * (low - quarters[0])
			high = 2*(high-quarters[0]) + 1
			var add uint8 = 0
			if index < uint32(len(inputBits)) {
				if inputBits[index] {
					add = 1
				}
			}
			tempByte := currentByte
			tempByte = 2*(uint32(tempByte)-quarters[0]) + uint32(add)
			currentByte = tempByte
			index++

		}

		d.currentInput = byteToBitSlice(uint32(currentByte), 32)
		d.high = high
		d.low = low
		d.step = step
		d.index = index
	}

}

//TODO: Cache the found symbol as it will be the same symbol look up a lot, or use a better lookup algorithm
//Returns the symbol that is represented by an interval number
func (d *Decoder) intervalToSymbol(symbolInterval uint32, hmap map[uint32]int,
	lsorted []uint32, hsorted []uint32) uint8 {

	//Binary search
	start := 0
	end := len(lsorted) - 1
	match := 0
	//If the interval is 0, then we're looking for the first non-zero symbol, or the symbol with 0 low and non 0 high
	for start <= end {
		middle := (start + end) / 2

		if symbolInterval >= lsorted[middle] && symbolInterval < hsorted[middle] {
			match = middle
			//Used for later lookup(to get the index back)
			hvalue := hsorted[match]
			//Look up this value in the saved map(either is fine)
			index := hmap[hvalue]
			return uint8(index)
		}
		if symbolInterval >= hsorted[middle] {
			start = middle + 1
		}
		if symbolInterval < lsorted[middle] {
			end = middle - 1
		}
	}
	return 0
}

//In case we want to build without a coder
func (d *Decoder) quarterize(upperLimit uint32) []uint32 {
	for i := 0; i < 3; i++ {
		// turn to uint64 then back
		upperLimitConverted := uint64(upperLimit)
		d.quarters[i] = uint32(((upperLimitConverted + 1) / 4) * uint64(i+1))
	}
	//2^32 or 2^32 - 1 shouldnt matter, we avoid a lot of conversion with it
	d.quarters[3] = upperLimit
	return d.quarters

}
