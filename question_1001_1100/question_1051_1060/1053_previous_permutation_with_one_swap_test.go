package question_1051_1060

import (
	"reflect"
	"testing"
)

func Test_prevPermOpt1(t *testing.T) {
	tests := []struct {
		A    []int
		want []int
	}{
		{[]int{3, 2, 1}, []int{3, 1, 2}},
		{[]int{1, 1, 5}, []int{1, 1, 5}},
		{[]int{1, 9, 4, 6, 7}, []int{1, 7, 4, 6, 9}},
		{[]int{3, 1, 1, 3}, []int{1, 3, 1, 3}},
		{[]int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := prevPermOpt1(tt.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prevPermOpt1() = %v, want %v", got, tt.want)
			}
		})
	}
}
