package main

import (
	"bufio"
	"fmt"
	"golang.org/x/image/bmp"
	"os"
	"strconv"
    "math"
)

func main() {
	//operation, c or d for compression/decompression respectively
	op := os.Args[1]
	input := os.Args[2]
	output := os.Args[3]
	//threshold for compression
	thr, err := strconv.ParseFloat(os.Args[4], 32)
	errCheck(err)
	if op == "c" {
		f, err := os.Open(input)
		errCheck(err)
		r := bufio.NewReader(f)
		fi, err := f.Stat()
		errCheck(err)
		fmt.Println("The file is %d bytes long before performing HAAR ", fi.Size())
		btmp, err := bmp.Decode(r)
		errCheck(err)
		pixels := make([][]float32, btmp.Bounds().Size().X)
		fmt.Println("Bitmap dimensions : ", btmp.Bounds().Size())
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
		haar := blocks(pixels, btmp.Bounds().Size().X, btmp.Bounds().Size().Y, float32(thr))
		//Prepend
		//data = append([]string{"Prepend Item"}, data...)
		by := float32ToBytes(haar)
		bound := math.Log2(float64(btmp.Bounds().Size().X))
        byteBound := byte(bound)
		by = append(byteBound, by...)
		fmt.Println("Haar length", len(by))
		file, err := os.OpenFile(output, os.O_RDWR, 0644)
		defer file.Close()
		errCheck(err)
		_, err = file.Write(by)
		errCheck(err)
	}
}
