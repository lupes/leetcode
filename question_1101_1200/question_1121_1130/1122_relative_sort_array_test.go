package question_1121_1130

import (
	"reflect"
	"testing"
)

func Test_relativeSortArray(t *testing.T) {
	tests := []struct {
		arr1 []int
		arr2 []int
		want []int
	}{
		{[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}, []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := relativeSortArray(tt.arr1, tt.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relativeSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
