package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func writeBinaryFile(fileName string, bytesToWrite []byte ){
	if len(bytesToWrite) == 0{
		fmt.Print("No bytes to write")
	}
		file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
		defer file.Close()
		ErrCheck(err)
		_, err = file.Write(bytesToWrite)
		ErrCheck(err)
		fmt.Println("Done writing")
		os.Exit(0)
}

func float32ToBytes(f []float32) []byte {
	by := make([]byte,len(f)*4)
	//for i :=0; i < len(f); i++ {
		var buf bytes.Buffer
		err := binary.Write(&buf, binary.LittleEndian, f)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
		by = buf.Bytes()
	//}
	return by
}

