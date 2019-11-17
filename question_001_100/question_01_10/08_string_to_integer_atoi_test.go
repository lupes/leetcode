package question_01_10

import (
	"math"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test#1", args{""}, 0},
		{"test#2", args{"w1"}, 0},
		{"test#3", args{"   "}, 0},
		{"test#4", args{"5"}, 5},
		{"test#5", args{"+5"}, 5},
		{"test#6", args{"-5"}, -5},
		{"test#7", args{"+w5"}, 0},
		{"test#8", args{"-w5"}, 0},
		{"test#9", args{"-5w"}, -5},
		{"test#10", args{"   1"}, 1},
		{"test#11", args{"   1c"}, 1},
		{"test#12", args{"1234567890123456"}, math.MaxInt32},
		{"test#13", args{"-1234567890123456"}, math.MinInt32},
		{"test#14", args{"2147483646"}, 2147483646},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
