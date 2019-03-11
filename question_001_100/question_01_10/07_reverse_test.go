package question_01_10

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test#1", args{0}, 0},
		{"test#2", args{123}, 321},
		{"test#3", args{-123}, -321},
		{"test#4", args{120}, 21},
		{"test#5", args{2147483648}, 0},
		{"test#6", args{-2147483648}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
