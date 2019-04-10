package question_41_50

import "testing"

func Test_multiply(t *testing.T) {
	tests := []struct {
		name string
		num1 string
		num2 string
		want string
	}{
		{"test#0", "", "", ""},
		{"test#1", "0", "", ""},
		{"test#2", "", "0", ""},
		{"test#3", "0", "0", "0"},
		{"test#4", "0", "1", "0"},
		{"test#5", "1", "0", "0"},
		{"test#6", "1", "1", "1"},
		{"test#7", "2", "3", "6"},
		{"test#8", "123", "456", "56088"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.num1, tt.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
