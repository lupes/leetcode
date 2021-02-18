package question_1461_1470

import (
	"reflect"
	"testing"
)

func Test_shuffle(t *testing.T) {
	tests := []struct {
		nums []int
		n    int
		want []int
	}{
		{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := shuffle(tt.nums, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}
