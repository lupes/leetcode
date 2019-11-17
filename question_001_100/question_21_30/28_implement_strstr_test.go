package question_0011_0020

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test#0", args{"", ""}, -1},
		{"test#1", args{"", "l"}, -1},
		{"test#2", args{"l", ""}, -1},
		{"test#3", args{"l", "ll"}, -1},
		{"test#4", args{"l", "l"}, 0},
		{"test#5", args{"hello", "ll"}, 2},
		{"test#6", args{"aaaaa", "bba"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
