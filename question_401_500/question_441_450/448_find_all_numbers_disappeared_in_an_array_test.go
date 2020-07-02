package question_441_450

import (
	"reflect"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
		{[]int{1, 1}, []int{2}},
		{[]int{2, 2}, []int{1}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findDisappearedNumbers(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
