package question_1631_1640

import (
	"reflect"
	"testing"
)

func Test_frequencySort(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 1, 2, 2, 2, 3}, []int{3, 1, 1, 2, 2, 2}},
		{[]int{2, 3, 1, 3, 2}, []int{1, 3, 3, 2, 2}},
		{[]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}, []int{5, -1, 4, 4, -6, -6, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := frequencySort(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
