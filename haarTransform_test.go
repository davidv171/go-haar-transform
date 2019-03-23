package main

import (
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
		{"1D TEST", args{[]float32{88, 88, 89, 90, 92, 94, 96, 97}, -532, 0}, []float32{88, 89.5, 93, 96.5, 0, -0.5, -1, -0.5}},
		{"2 DEPTH", args{[]float32{88, 89.5, 93, 96.5, 0, -0.5, -1, -0.5}, -532, 4}, []float32{88.75, 94.75, -0.75, -1.75, 0, -0.5, -1, -0.5}},
		{"3 DEPTH", args{[]float32{88.75, 94.75, -0.75, -1.75, 0, -0.5, -1, -0.5}, -532, 6}, []float32{91.75, -3, -0.75, -1.75, 0, -0.5, -1, -0.5}},
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
		{"BASIC TEST", args{[][]float32{{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16}, {1, 2, 3, 4, 5, 6, 7, 24},
			{1, 2, 3, 4, 5, 6, 7, 32}, {1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
			{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 7, 64},
		}, 8, 8}, [][]float32{{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16}, {1, 2, 3, 4, 5, 6, 7, 24},
			{1, 2, 3, 4, 5, 6, 7, 32}, {1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
			{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 7, 64},
		}},
		{"Double loops", args{[][]float32{{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 16}, {1, 2, 3, 4, 5, 6, 7, 24},
			{1, 2, 3, 4, 5, 6, 7, 32}, {1, 2, 3, 4, 5, 6, 7, 40}, {1, 2, 3, 4, 5, 6, 7, 48},
			{1, 2, 3, 4, 5, 6, 7, 56}, {1, 2, 3, 4, 5, 6, 7, 64},
			{1, 2, 3, 4, 5, 6, 7, 72}, {1, 2, 3, 4, 5, 6, 7, 80}, {1, 2, 3, 4, 5, 6, 7, 88},
			{1, 2, 3, 4, 5, 6, 7, 96}, {1, 2, 3, 4, 5, 6, 7, 104}, {1, 2, 3, 4, 5, 6, 7, 112},
			{1, 2, 3, 4, 5, 6, 7, 120}, {1, 2, 3, 4, 5, 6, 7, 128},
		}, 16, 16}, [][]float32{{1, 2, 3, 4, 5, 6, 7, 72}, {1, 2, 3, 4, 5, 6, 7, 80}, {1, 2, 3, 4, 5, 6, 7, 88},
			{1, 2, 3, 4, 5, 6, 7, 96}, {1, 2, 3, 4, 5, 6, 7, 104}, {1, 2, 3, 4, 5, 6, 7, 112},
			{1, 2, 3, 4, 5, 6, 7, 120}, {1, 2, 3, 4, 5, 6, 7, 128},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := blocks(tt.args.pixels, tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("blocks() = %v, want %v", got, tt.want)
			}
		})
	}
}
