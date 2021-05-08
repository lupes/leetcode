package question_1681_1690

import (
	"reflect"
	"testing"
)

func Test_getSumAbsoluteDifferences(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{2, 3, 5}, []int{4, 3, 5}},
		{[]int{1, 4, 6, 8, 10}, []int{24, 15, 13, 15, 21}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := getSumAbsoluteDifferences(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSumAbsoluteDifferences() = %v, want %v", got, tt.want)
			}
		})
	}
}
