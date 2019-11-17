package question_01_10

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test#1", args{""}, 0},
		{"test#2", args{"abcabcbb"}, 3},
		{"test#3", args{"bbbbb"}, 1},
		{"test#4", args{"pwwkew"}, 3},
		{"test#5", args{" "}, 1},
		{"test#6", args{"au"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
