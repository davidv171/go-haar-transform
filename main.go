package main

import (
	"bufio"
	"fmt"
	"golang.org/x/image/bmp"
	"os"
)

func main() {
	//operation, c or d for compression/decompression respectively
	op := os.Args[1]
	input := os.Args[2]
	output := os.Args[3]
	//threshold for compression
	thr := os.Args[4]
	fmt.Println("ARgs: ", op, input, output, thr)
	f,err := os.Open(input)
	errCheck(err)
	r := bufio.NewReader(f)
	btmp,err := bmp.Decode(r)
	errCheck(err)
	fmt.Println("RESUKLT", btmp.Bounds())
	for i := 0; i < btmp.Bounds().Max.X; i++ {
		for j:=0; j < btmp.Bounds().Max.Y; j++ {
		}
	}
	fte := "temp2.bmp"
	ft, err := os.OpenFile(fte, os.O_RDWR|os.O_CREATE, 0755)
	defer ft.Close()
	errCheck(err)
	w := bufio.NewWriter(ft)
	defer w.Flush()
	err = bmp.Encode(w,btmp)
	errCheck(err)
}

