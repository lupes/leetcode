package question_1581_1590

import (
	"reflect"
	"testing"
)

func Test_sumOddLengthSubarrays(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 4, 2, 5, 3}, 58},
		{[]int{1, 2}, 3},
		{[]int{10, 11, 12}, 66},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := sumOddLengthSubarrays(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumOddLengthSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
