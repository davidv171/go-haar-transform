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
	f, err := os.Open(input)
	errCheck(err)
	r := bufio.NewReader(f)
	btmp, err := bmp.Decode(r)
	errCheck(err)
	pixels := make([][]float32, btmp.Bounds().Size().X)
	fmt.Println("RESUKLT", btmp.Bounds().Size())
	for i := 0; i < btmp.Bounds().Size().X; i++ {
		pixels[i] = make([]float32, btmp.Bounds().Size().Y)
		for j := 0; j < btmp.Bounds().Size().Y; j++ {
			pix, _, _, _ := btmp.At(i, j).RGBA()
			//we're dealing with n bit depth gray pixel, the library always does 0-65635
			pix = pix >> 8
			pixels[i][j] = float32(pix)
		}
	}
	//We assume a symmetric BMP
	blocks(pixels, btmp.Bounds().Size().X, btmp.Bounds().Size().Y)
	ft, err := os.OpenFile(output, os.O_RDWR|os.O_CREATE, 0755)
	defer ft.Close()
	errCheck(err)
	w := bufio.NewWriter(ft)
	defer w.Flush()
	err = bmp.Encode(w, btmp)
	errCheck(err)
}
