package question_481_490

import "testing"

func Test_smallestGoodBase(t *testing.T) {
	tests := []struct {
		name string
		n    string
		want string
	}{
		{"test", "13", "3"},
		{"test", "4681", "8"},
		{"test", "4681", "8"},
		{"test", "1000000000000000000", "999999999999999999"},
		//{"test", "987654321", "0"},ij
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestGoodBase(tt.n); got != tt.want {
				t.Errorf("smallestGoodBase() = %v, want %v", got, tt.want)
			}
		})
	}
}
