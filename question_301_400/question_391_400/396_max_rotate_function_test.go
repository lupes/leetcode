package question_391_400

import "testing"

func Test_maxRotateFunction(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		want int
	}{
		{"test", []int{4, 3, 2, 6}, 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRotateFunction(tt.A); got != tt.want {
				t.Errorf("maxRotateFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
