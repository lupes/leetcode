package question_501_510

import (
	"reflect"
	"testing"
)

func Test_findRelativeRanks(t *testing.T) {
	tests := []struct {
		nums []int
		want []string
	}{
		{[]int{5, 4, 3, 2, 1}, []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findRelativeRanks(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRelativeRanks() = %v, want %v", got, tt.want)
			}
		})
	}
}
