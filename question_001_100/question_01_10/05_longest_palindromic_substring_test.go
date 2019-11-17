package question_01_10

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test#1", args{""}, ""},
		{"test#2", args{"a"}, "a"},
		{"test#3", args{"babad"}, "bab"},
		{"test#4", args{"cbbd"}, "bb"},
		{"test#5", args{"ccc"}, "ccc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
