package question_641_650

import (
	"reflect"
	"testing"
)

func Test_findErrorNums(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 2, 4}, []int{2, 3}},
		{[]int{3, 2, 3, 4, 6, 5}, []int{3, 1}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findErrorNums(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findErrorNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
