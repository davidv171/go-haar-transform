package main

import (
	"fmt"
	"os"
)

func readBinaryFile(arithmeticCoder *Coder, operation string, modelCreation bool, arithmeticDecoder *Decoder, outputFile string, data []byte) {
	fileSize := len(data)
	var bufferSize int64
	//YOLO
	bufferSize = int64(fileSize)
	var bufferOverflow int64 = 0
	//Data where we put the read bytes into

	//for _,aByte := range data {
	//Add the new 8 booleans to the end of the bits array
	//bits = append(bits,byteToBitSlice(&aByte)...)
	//}
	if operation == "c" {
		if modelCreation {
			arithmeticCoder.frequencyTableGenerator(data)
		} else {
			arithmeticCoder.intervalCalculation(data)

		}
	} else if operation == "d" {
		arithmeticDecoder.readFreqTable(data)
		outputBytes := arithmeticDecoder.output
		writeBin(outputFile, outputBytes, 0)

	}
	bufferOverflow += bufferSize
	//So we're aware of indexes if the file is large
	/*After we are done encoding the values byte by byte, we look at the rest:
	- if low < firstQuarter : output 01 and E3_COUNTER times bit 1
	- else : output 10 and E3_COUNTER times bit 0
	*/
	if !modelCreation && arithmeticCoder != nil {
		writeEncoded(arithmeticCoder, outputFile)
		//fmt.Println("The rest:")

	}

}
func writeBin(fileName string, bytesToWrite []byte, bufferOverflow int64) {
	if bufferOverflow == 0 {
		_, err := os.Create(fileName)
		errCheck(err)
	}
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	defer file.Close()
	errCheck(err)
	_, err = file.WriteAt(bytesToWrite, bufferOverflow)
	errCheck(err)
	os.Exit(0)
}
func writeEncoded(arithmeticCoder *Coder, fileName string) {
	if arithmeticCoder.low < arithmeticCoder.quarters[0] {
		arithmeticCoder.outputBits = append(arithmeticCoder.outputBits, false, true)
		for i := 0; uint32(i) < arithmeticCoder.e3Counter; i++ {
			arithmeticCoder.outputBits = append(arithmeticCoder.outputBits, true)
		}
	} else {
		arithmeticCoder.outputBits = append(arithmeticCoder.outputBits, true, false)

		for i := 0; uint32(i) < arithmeticCoder.e3Counter; i++ {
			arithmeticCoder.outputBits = append(arithmeticCoder.outputBits, false)
		}
	}
	fmt.Println("")
	//Write the 32uint[256] high table into file
	//If the value in high table is 0,
	// we can just write 4 0 bytes into the table, this saves us some time when doing compression on a small amount of unique symbols
	outputBytes := make([]byte, 0)
	//TODO: decide on number of written bytes based on the highest value
	for i := 0; i < 256; i++ {
		currentElement := arithmeticCoder.highTable[i]
		if currentElement != 0 {
			tempSlice := byteToBitSlice(currentElement, 32)
			outputBytes = append(outputBytes, bitSliceToByte(&tempSlice, 4)...)

		} else {
			outputBytes = append(outputBytes, 0, 0, 0, 0)
		}

	}
	for i := 0; i < len(arithmeticCoder.outputBits); i += 8 {
		tempSlice := arithmeticCoder.outputBits[i : i+8]
		outputBytes = append(outputBytes, bitSliceToByte(&tempSlice, 1)[0])

	}
	fmt.Println("Compressed output file size ", len(outputBytes))
	writeBin(fileName, outputBytes, 0)
}
