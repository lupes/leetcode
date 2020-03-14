package question_911_920

import (
	"reflect"
	"testing"
)

func Test_sortArray(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
		{[]int{5, 2, 3, 1}, []int{1, 2, 3, 5}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := sortArray(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
