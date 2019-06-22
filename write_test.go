package main

import (
	"reflect"
	"testing"
)

func Test_float32ToBytes(t *testing.T) {
	type args struct {
		input []float32
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"Basic test", args{[]float32{1.0,2.0,3.0,4.0,5.0,6.0,7.0}},[]byte{0,0,0,1,0,0,0,2,0,0,0,3,0,0,0,4,0,0,0,5,0,0,0,6,0,0,0,7}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := float32ToBytes(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("float32ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
