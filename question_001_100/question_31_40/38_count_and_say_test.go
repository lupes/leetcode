package question_31_40

import "testing"

func Test_countAndSay(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{"test#1", 1, "1"},
		{"test#2", 2, "11"},
		{"test#3", 3, "21"},
		{"test#4", 4, "1211"},
		{"test#5", 5, "111221"},
		{"test#6", 6, "312211"},
		{"test#7", 7, "13112221"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
