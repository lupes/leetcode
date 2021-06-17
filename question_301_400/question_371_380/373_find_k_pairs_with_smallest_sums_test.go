package question_371_380

import (
	"reflect"
	"testing"
)

func Test_kSmallestPairs(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		k     int
		want  [][]int
	}{
		{[]int{1, 7, 11}, []int{2, 4, 6}, 3,
			[][]int{{1, 2}, {1, 4}, {1, 6}},
		},
		{
			[]int{1, 1, 2},
			[]int{1, 2, 3}, 2,
			[][]int{{1, 1}, {1, 1}},
		},
		{
			[]int{1, 2},
			[]int{3}, 3,
			[][]int{{1, 3}, {2, 3}},
		},
		{
			[]int{1, 1, 2},
			[]int{1, 2, 3}, 9,
			[][]int{{1, 1}, {1, 1}, {1, 2}, {1, 2}, {2, 1}, {1, 3}, {1, 3}, {2, 2}, {2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := kSmallestPairs(tt.nums1, tt.nums2, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kSmallestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
