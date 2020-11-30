package question_01_10

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test#1", args{"LEETCODEISHIRING", 3}, "LCIRETOESIIGEDHN"},
		{"test#2", args{"LEETCODEISHIRING", 4}, "LDREOEIIECIHNTSG"},
		{"test#3", args{"LEETCODEISHIRING", 1}, "LEETCODEISHIRING"},
		{"test#3", args{"", 1}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
