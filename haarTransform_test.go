package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_pairing(t *testing.T) {
	type args struct {
		input []float32
		thr   float32
		depth int
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{"1D TEST", args{[]float32{88, 88, 89, 90, 92, 94, 96, 97}, -151, 0},
			[]float32{91.75, -3, -0.75, -1.75, 0, -0.5, -1, -0.5}},
		/*{"2 DEPTH", args{[]float32{88, 89.5, 93, 96.5, 0, -0.5, -1, -0.5}, -532, 4},
			[]float32{88.75, 94.75, -0.75, -1.75, 0, -0.5, -1, -0.5}},
		{"3 DEPTH", args{[]float32{88.75, 94.75, -0.75, -1.75, 0, -0.5, -1, -0.5},
			-532, 6}, []float32{91.75, -3, -0.75, -1.75, 0, -0.5, -1, -0.5}},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := haar(tt.args.input, tt.args.thr, tt.args.depth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pairing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_blocks(t *testing.T) {
	type args struct {
		pixels [][]float32
		x      int
		y      int
	}
	tests := []struct {
		name string
		args args
		want [][]float32
	}{
		{"BASIC TEST", args{[][]float32{{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16},
			{1, 2, 3, 4, 5, 6, 7, 24}, {1, 2, 3, 4, 5, 6, 7, 32},
			{1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
			{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 7, 64},
		}, 8, 8}, [][]float32{{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16},
			{1, 2, 3, 4, 5, 6, 7, 24}, {1, 2, 3, 4, 5, 6, 7, 32},
			{1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
			{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 7, 64},
		}},
		{"Double loops", args{[][]float32{{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16},
			{1, 2, 3, 4, 5, 6, 7, 24},
			{1, 2, 3, 4, 5, 6, 7, 32}, {1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
			{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 7, 64},
			{1, 2, 3, 4, 5, 6, 7, 72}, {1, 2, 3, 4, 5, 6, 7, 80}, {1, 2, 3, 4, 5, 6, 7, 88},
			{1, 2, 3, 4, 5, 6, 7, 96}, {1, 2, 3, 4, 5, 6, 7, 104}, {1, 2, 3, 4, 5, 6, 7, 112},
			{1, 2, 3, 4, 5, 6, 7, 120}, {1, 2, 3, 4, 5, 6, 7, 128},
		}, 16, 16}, [][]float32{{1, 2, 3, 4, 5, 6, 7, 72}, {1, 2, 3, 4, 5, 6, 7, 80},
			{1, 2, 3, 4, 5, 6, 7, 88},
			{1, 2, 3, 4, 5, 6, 7, 96}, {1, 2, 3, 4, 5, 6, 7, 104},
			{1, 2, 3, 4, 5, 6, 7, 112},
			{1, 2, 3, 4, 5, 6, 7, 120}, {1, 2, 3, 4, 5, 6, 7, 128},
		}},
		{"Multiple lines", args{[][]float32{{0, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16},
			{1, 2, 3, 4, 5, 6, 7, 24},
			{1, 2, 3, 4, 5, 6, 7, 32}, {1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
			{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 0, 64},
			{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16}, {1, 2, 3, 4, 5, 6, 7, 24},
			{1, 2, 3, 4, 5, 6, 7, 32}, {1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
			{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 0, 6, 7, 64},
			{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16}, {1, 2, 3, 4, 5, 6, 7, 24},
			{1, 2, 3, 4, 5, 6, 7, 32}, {1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
			{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 7, 64},
			{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16}, {1, 2, 3, 4, 5, 6, 7, 24},
			{1, 2, 3, 4, 5, 6, 7, 32}, {1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
			{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 7, 64},
			{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16}, {1, 2, 3, 4, 5, 6, 7, 24},
			{1, 2, 3, 4, 5, 6, 7, 32}, {1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
			{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 7, 64},
			{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16}, {1, 2, 3, 4, 5, 6, 7, 24},
			{1, 2, 3, 4, 5, 6, 7, 32}, {1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
			{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 7, 64},
			{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16}, {1, 2, 3, 4, 5, 6, 7, 24},
			{1, 2, 3, 4, 5, 6, 7, 32}, {1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
			{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 7, 64},
			{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16}, {1, 2, 3, 4, 5, 6, 7, 124},
			{1, 2, 3, 4, 5, 6, 7, 32}, {1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 148},
			{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 7, 164},
			{1, 2, 3, 4, 5, 6, 7, 72}, {1, 2, 3, 4, 5, 6, 7, 80}, {1, 2, 3, 4, 5, 6, 7, 88},
			{1, 2, 3, 4, 5, 6, 7, 96}, {1, 2, 3, 4, 5, 6, 7, 104}, {1, 2, 3, 4, 5, 6, 7, 112},
			{1, 2, 3, 4, 5, 6, 7, 120}, {1, 2, 3, 4, 5, 6, 7, 128},
		}, 72, 72},
			[][]float32{{1, 2, 3, 4, 5, 6, 7, 72}, {1, 2, 3, 4, 5, 6, 7, 80},
				{1, 2, 3, 4, 5, 6, 7, 88},
				{1, 2, 3, 4, 5, 6, 7, 96}, {1, 2, 3, 4, 5, 6, 7, 104}, {1, 2, 3, 4, 5, 6, 7, 112},
				{1, 2, 3, 4, 5, 6, 7, 120}, {1, 2, 3, 4, 5, 6, 7, 128},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//We can expect normal threshold behavior....r-right?
			if got := blocks(tt.args.pixels, tt.args.x, tt.args.y, 0.5); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("blocks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_haar(t *testing.T) {
	type args struct {
		input []float32
		thr   float32
		depth int
	}
	var tests []struct {
		name string
		args args
		want []float32
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := haar(tt.args.input, tt.args.thr, tt.args.depth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("haar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRow(t *testing.T) {
	type args struct {
		block [][]float32
		index int
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		//GET first row of 2D 8x8 block
		{"Basic test", args{
			[][]float32{{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16},
				{1, 2, 3, 4, 5, 6, 7, 24}, {1, 2, 3, 4, 5, 6, 7, 32},
				{1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
				{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 7, 64},
			}, 0}, []float32{1, 2, 3, 4, 5, 6, 7, 8},
		// TODO: Add test cases.
		},
		//GET 2nd row of 2D 8x8 block
		{"Basic test on second row", args{
			[][]float32{{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16},
				{1, 2, 3, 4, 5, 6, 7, 24}, {1, 2, 3, 4, 5, 6, 7, 32},
				{1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
				{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 7, 64}},
			1}, []float32{1, 2, 3, 4, 5, 6, 7, 16},
		},
		//GET last row of 2D 8x8 block
		{"Basic test on second row", args{
			[][]float32{{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16},
				{1, 2, 3, 4, 5, 6, 7, 24}, {1, 2, 3, 4, 5, 6, 7, 32},
				{1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
				{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 7, 64}},
			7}, []float32{1, 2, 3, 4, 5, 6, 7, 64},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRow(tt.args.block, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getColumn(t *testing.T) {
	type args struct {
		block [][]float32
		index int
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		//GET first column of 2D 8x8 block
		{"Basic test", args{
			[][]float32{{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16},
				{1, 2, 3, 4, 5, 6, 7, 24}, {1, 2, 3, 4, 5, 6, 7, 32},
				{1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
				{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 7, 64},
			}, 0}, []float32{1, 1, 1, 1, 1, 1, 1, 1},
		// TODO: Add test cases.
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getColumn(tt.args.block, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_blocksT(t *testing.T) {
	type args struct {
		block [][][]float32
		thr   float32
	}
	tests := []struct {
		name string
		args args
		want [][][]float32
	}{
		{"Basic test case", args{
			[][][]float32{{
				{88, 88, 89, 90, 92, 94, 96, 97},
				{90, 90, 91, 92, 93, 95, 97, 97},
				{92, 92, 93, 94, 95, 96, 97, 97},
				{93, 93, 94, 95, 96, 96, 96, 96},
				{92, 93, 95, 96, 96, 96, 96, 95},
				{92, 94, 96, 98, 99, 99, 98, 97},
				{94, 96, 99, 101, 103, 103, 102, 101},
				{95, 97, 101, 104, 106, 106, 105, 105},
			}}, -42,
		}, [][][]float32{
			{
				{96, -2.0312, -1.5312, -0.2188, -0.4375, -0.75, -0.3125, 0.125},
				{-2.4375, -0.0312, 0.7812, -0.7812, 0.4375, 0.25, -0.3125, -0.25},
				{-1.125, -0.625, 0, -0.625, 0, 0, -0.375, -0.125},
				{-2.6875, 0.75, 0.5625, -0.0625, 0.125, 0.25, 0, 0.125},
				{-0.6875, -0.3125, 0, -0.125, 0, 0, 0, -0.25},
				{-0.1875, -0.3125, 0, -0.375, 0, 0, -0.25, 0},
				{-0.875, 0.375, 0.25, -0.25, 0.25, 0.25, 0, 0},
				{-1.25, 0.375, 0.375, 0.125, 0, 0.25, 0, 0.25},
			}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := blocksT(tt.args.block, tt.args.thr); !reflect.DeepEqual(got, tt.want) {
				for x := 0; x < len(tt.args.block); x++ {
					for i := 0; i < 8; i++ {
						for j := 0; j < 8; j++ {
							if got[x][i][j] != tt.want[x][i][j] {
								fmt.Println("Difference :  got: ", got[i][j], " ,wanted ", tt.want[i][j], " on [", i, ",", j, "]")
							}
						}
					}
				}
				t.Errorf("blocksT() = %v,\n want %v", got, tt.want)
			}
		})
	}
}

func Test_zigZag(t *testing.T) {
	type args struct {
		block [][][]float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		// TODO: Add test cases.
		{"BASIC TEST", args{[][][]float32{
			{
				{0, 1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20, 21, 22, 23}, {24, 25, 26, 27, 28, 29, 30, 31},
				{32, 33, 34, 35, 36, 37, 38, 39}, {40, 41, 42, 43, 44, 45, 46, 47},
				{48, 49, 50, 51, 52, 53, 54, 55}, {56, 57, 58, 59, 60, 61, 62, 63},
			}, {
				{0, 1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20, 21, 22, 23}, {24, 25, 26, 27, 28, 29, 30, 31},
				{32, 33, 34, 35, 36, 37, 38, 39}, {40, 41, 42, 43, 44, 45, 46, 47},
				{48, 49, 50, 51, 52, 53, 54, 55}, {56, 57, 58, 59, 60, 61, 62, 63},
			},
		}},
			[]float32{
				0, 0,
				1, 1, 8, 8,
				16, 16, 9, 9, 2, 2,
				3, 3, 10, 10, 17, 17, 24, 24,
				32, 32, 25, 25, 18, 18, 11, 11, 4, 4,
				5, 5, 12, 12, 19, 19, 26, 26, 33, 33, 40, 40,
				48, 48, 41, 41, 34, 34, 27, 27, 20, 20, 13, 13, 6, 6,
				7, 7, 14, 14, 21, 21, 28, 28, 35, 35, 42, 42, 49, 49, 56, 56,
				57, 57, 50, 50, 43, 43, 36, 36, 29, 29, 22, 22, 15, 15,
				23, 23, 30, 30, 37, 37, 44, 44, 51, 51, 58, 58,
				59, 59, 52, 52, 45, 45, 38, 38, 31, 31,
				39, 39, 46, 46, 53, 53, 60, 60,
				61, 61, 54, 54, 47, 47,
				55, 55, 62, 62,
				63, 63,
			},
		},
		{"BIG BOI", args{[][][]float32{
			{
				{0, 1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20, 21, 22, 23}, {24, 25, 26, 27, 28, 29, 30, 31},
				{32, 33, 34, 35, 36, 37, 38, 39}, {40, 41, 42, 43, 44, 45, 46, 47},
				{48, 49, 50, 51, 52, 53, 54, 55}, {56, 57, 58, 59, 60, 61, 62, 63},
			}, {
				{0, 1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20, 21, 22, 23}, {24, 25, 26, 27, 28, 29, 30, 31},
				{32, 33, 34, 35, 36, 37, 38, 39}, {40, 41, 42, 43, 44, 45, 46, 47},
				{48, 49, 50, 51, 52, 53, 54, 55}, {56, 57, 58, 59, 60, 61, 62, 63},
			},
			{
				{0, 1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20, 21, 22, 23}, {24, 25, 26, 27, 28, 29, 30, 31},
				{32, 33, 34, 35, 36, 37, 38, 39}, {40, 41, 42, 43, 44, 45, 46, 47},
				{48, 49, 50, 51, 52, 53, 54, 55}, {56, 57, 58, 59, 60, 61, 62, 63},
			}, {
				{0, 1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20, 21, 22, 23}, {24, 25, 26, 27, 28, 29, 30, 31},
				{32, 33, 34, 35, 36, 37, 38, 39}, {40, 41, 42, 43, 44, 45, 46, 47},
				{48, 49, 50, 51, 52, 53, 54, 55}, {56, 57, 58, 59, 60, 61, 62, 63},
			}, {
				{0, 1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20, 21, 22, 23}, {24, 25, 26, 27, 28, 29, 30, 31},
				{32, 33, 34, 35, 36, 37, 38, 39}, {40, 41, 42, 43, 44, 45, 46, 47},
				{48, 49, 50, 51, 52, 53, 54, 55}, {56, 57, 58, 59, 60, 61, 62, 63},
			}, {
				{0, 1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20, 21, 22, 23}, {24, 25, 26, 27, 28, 29, 30, 31},
				{32, 33, 34, 35, 36, 37, 38, 39}, {40, 41, 42, 43, 44, 45, 46, 47},
				{48, 49, 50, 51, 52, 53, 54, 55}, {56, 57, 58, 59, 60, 61, 62, 63},
			}, {
				{0, 1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20, 21, 22, 23}, {24, 25, 26, 27, 28, 29, 30, 31},
				{32, 33, 34, 35, 36, 37, 38, 39}, {40, 41, 42, 43, 44, 45, 46, 47},
				{48, 49, 50, 51, 52, 53, 54, 55}, {56, 57, 58, 59, 60, 61, 62, 63},
			}, {
				{0, 1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20, 21, 22, 23}, {24, 25, 26, 27, 28, 29, 30, 31},
				{32, 33, 34, 35, 36, 37, 38, 39}, {40, 41, 42, 43, 44, 45, 46, 47},
				{48, 49, 50, 51, 52, 53, 54, 55}, {56, 57, 58, 59, 60, 61, 62, 63},
			}, {
				{0, 1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20, 21, 22, 23}, {24, 25, 26, 27, 28, 29, 30, 31},
				{32, 33, 34, 35, 36, 37, 38, 39}, {40, 41, 42, 43, 44, 45, 46, 47},
				{48, 49, 50, 51, 52, 53, 54, 55}, {56, 57, 58, 59, 60, 61, 62, 63},
			}, {
				{0, 1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20, 21, 22, 23}, {24, 25, 26, 27, 28, 29, 30, 31},
				{32, 33, 34, 35, 36, 37, 38, 39}, {40, 41, 42, 43, 44, 45, 46, 47},
				{48, 49, 50, 51, 52, 53, 54, 55}, {56, 57, 58, 59, 60, 61, 62, 63},
			}, {
				{0, 1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20, 21, 22, 23}, {24, 25, 26, 27, 28, 29, 30, 31},
				{32, 33, 34, 35, 36, 37, 38, 39}, {40, 41, 42, 43, 44, 45, 46, 47},
				{48, 49, 50, 51, 52, 53, 54, 55}, {56, 57, 58, 59, 60, 61, 62, 63},
			}, {
				{0, 1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20, 21, 22, 23}, {24, 25, 26, 27, 28, 29, 30, 31},
				{32, 33, 34, 35, 36, 37, 38, 39}, {40, 41, 42, 43, 44, 45, 46, 47},
				{48, 49, 50, 51, 52, 53, 54, 55}, {56, 57, 58, 59, 60, 61, 62, 63},
			}, {
				{0, 1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20, 21, 22, 23}, {24, 25, 26, 27, 28, 29, 30, 31},
				{32, 33, 34, 35, 36, 37, 38, 39}, {40, 41, 42, 43, 44, 45, 46, 47},
				{48, 49, 50, 51, 52, 53, 54, 55}, {56, 57, 58, 59, 60, 61, 62, 63},
			},
		}},
			[]float32{
				0, 0,
				1, 1, 8, 8,
				16, 16, 9, 9, 2, 2,
				3, 3, 10, 10, 17, 17, 24, 24,
				32, 32, 25, 25, 18, 18, 11, 11, 4, 4,
				5, 5, 12, 12, 19, 19, 26, 26, 33, 33, 40, 40,
				48, 48, 41, 41, 34, 34, 27, 27, 20, 20, 13, 13, 6, 6,
				7, 7, 14, 14, 21, 21, 28, 28, 35, 35, 42, 42, 49, 49, 56, 56,
				57, 57, 50, 50, 43, 43, 36, 36, 29, 29, 22, 22, 15, 15,
				23, 23, 30, 30, 37, 37, 44, 44, 51, 51, 58, 58,
				59, 59, 52, 52, 45, 45, 38, 38, 31, 31,
				39, 39, 46, 46, 53, 53, 60, 60,
				61, 61, 54, 54, 47, 47,
				55, 55, 62, 62,
				63, 63,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigZag(tt.args.block); !reflect.DeepEqual(got, tt.want) {
				for i := 0; i < len(tt.want); i++ {
					if tt.want[i] != got[i] {
						fmt.Println("i: ", i, " -> ", got[i], " vs ", tt.want[i])
					}
				}
				t.Errorf("zigZag() = \n%v,\n want \n%v", got, tt.want)
			}
		})
	}
}
