package question_361_370

import (
	"reflect"
	"sort"
	"testing"
)

func Test_largestDivisibleSubset(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{1, 2, 4, 8}},
		{[]int{4, 8, 10, 240}, []int{4, 8, 240}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := largestDivisibleSubset(tt.nums)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestDivisibleSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
