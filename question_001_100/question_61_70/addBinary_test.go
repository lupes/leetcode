package question_61_70

import "testing"

func Test_addBinary(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{"test#0", "0", "0", "0"},
		{"test#1", "1", "0", "1"},
		{"test#2", "1", "1", "10"},
		{"test#3", "11", "1", "100"},
		{"test#4", "1010", "1011", "10101"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.a, tt.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
