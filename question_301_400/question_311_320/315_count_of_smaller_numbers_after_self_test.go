package question_311_320

import (
	"reflect"
	"testing"
)

func Test_countSmaller(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{5, 2, 6, 1}, []int{2, 1, 1, 0}},
		{[]int{-1, -1}, []int{0, 0}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := countSmaller(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSmaller() = %v, want %v", got, tt.want)
			}
		})
	}
}
