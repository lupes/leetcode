package question_51_60

import "testing"

func Test_getPermutation(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want string
	}{
		{"test#0", 0, 0, ""},
		{"test#1", 1, 1, "1"},
		{"test#2", 2, 2, "21"},
		{"test#3", 2, 3, ""},
		{"test#4", 3, 3, "213"},
		{"test#5", 4, 9, "2314"},
		{"test#6", 3, 6, "321"},
		{"test#7", 3, 2, "132"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPermutation(tt.n, tt.k); got != tt.want {
				t.Errorf("getPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
