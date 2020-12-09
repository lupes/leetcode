package question_0011_0020

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		nums  []int
		wants []int
		want  int
	}{
		{[]int{}, []int{}, 0},
		{[]int{1}, []int{1}, 1},
		{[]int{1, 1}, []int{1}, 1},
		{[]int{1, 1, 2}, []int{1, 2}, 2},
		{[]int{1, 1, 2, 2}, []int{1, 2}, 2},
		{[]int{1, 1, 2, 2, 3}, []int{1, 2, 3}, 3},
		{[]int{1, 2}, []int{1, 2}, 2},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 3},
		{[]int{1, 1, 2, 2, 3, 4}, []int{1, 2, 3, 4}, 4},
		{[]int{1, 1, 2, 2, 3, 4}, []int{1, 2, 3, 4}, 4},
		{[]int{1, 1, 2, 2, 3, 4, 5, 5}, []int{1, 2, 3, 4, 5}, 5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.nums[:got], tt.wants) {
				t.Errorf("removeDuplicates() = %v, want %v", tt.nums[:got], tt.wants)
			}
		})
	}
}
