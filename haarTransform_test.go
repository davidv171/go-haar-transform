package main

import (
	"reflect"
	"testing"
)

func Test_pairing(t *testing.T) {
	type args struct {
		input []float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{"1D TEST", args{[]float32{88,88,89,90,92,94,96,97}}, []float32{88,89.5,93,96.5,0,-0.5,-1,-0.5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairing(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pairing() = %v, want %v", got, tt.want)
			}
		})
	}
}
