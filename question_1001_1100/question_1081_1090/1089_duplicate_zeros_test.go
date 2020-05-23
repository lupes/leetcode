package question_1081_1090

import (
	"reflect"
	"testing"
)

func Test_duplicateZeros(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		//{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		//{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{8, 4, 5, 0, 0, 0, 0, 7}, []int{8, 4, 5, 0, 0, 0, 0, 0}},
		//{[]int{8, 4, 5, 0, 0, 0, 0, 0, 7}, []int{8, 4, 5, 0, 0, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			duplicateZeros(tt.arr)
			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Errorf("duplicateZeros() = %v,\n want %v", tt.arr, tt.want)
			}
		})
	}
}
