package question_971_980

import (
	"testing"
)

func Test_subarraysDivByK(t *testing.T) {
	tests := []struct {
		A    []int
		K    int
		want int
	}{
		{[]int{4, 5, 0, -2, -3, 1}, 5, 7},
		{[]int{-1, 2, 9}, 2, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := subarraysDivByK(tt.A, tt.K); got != tt.want {
				t.Errorf("subarraysDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
