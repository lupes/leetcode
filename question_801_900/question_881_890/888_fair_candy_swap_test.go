package question_881_890

import (
	"reflect"
	"testing"
)

func Test_fairCandySwap(t *testing.T) {
	tests := []struct {
		A    []int
		B    []int
		want []int
	}{
		{[]int{1, 1}, []int{2, 2}, []int{1, 2}},
		{[]int{1, 2}, []int{2, 3}, []int{1, 2}},
		{[]int{2}, []int{1, 3}, []int{2, 3}},
		{[]int{1, 2, 5}, []int{2, 4}, []int{5, 4}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := fairCandySwap(tt.A, tt.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
