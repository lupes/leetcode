package question_681_690

import (
	"reflect"
	"testing"
)

func Test_maxSumOfThreeSubarrays(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 1, 2, 6, 7, 5, 1}, 2, []int{0, 3, 5}},
		{[]int{1, 2, 1, 2, 1, 2, 1, 2, 1}, 2, []int{0, 2, 4}},
		{[]int{4, 5, 10, 6, 11, 17, 4, 11, 1, 3}, 1, []int{4, 5, 7}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxSumOfThreeSubarrays(tt.nums, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSumOfThreeSubarrays() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
