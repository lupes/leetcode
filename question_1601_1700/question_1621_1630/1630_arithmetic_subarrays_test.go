package question_1621_1630

import (
	"reflect"
	"testing"
)

func Test_checkArithmeticSubarrays(t *testing.T) {
	tests := []struct {
		nums []int
		l    []int
		r    []int
		want []bool
	}{
		{[]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}, []bool{true, false, true}},
		{[]int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}, []int{0, 1, 6, 4, 8, 7}, []int{4, 4, 9, 7, 9, 10}, []bool{false, true, false, false, true, true}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := checkArithmeticSubarrays(tt.nums, tt.l, tt.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkArithmeticSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
