package question_901_910

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParity(t *testing.T) {
	tests := []struct {
		A    []int
		want []int
	}{
		{[]int{3, 1, 2, 4}, []int{4, 2, 1, 3}},
		{[]int{3, 1, 1, 2, 4}, []int{4, 2, 1, 1, 3}},
		{[]int{2, 1, 3, 2, 4}, []int{2, 4, 2, 3, 1}},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := sortArrayByParity(tt.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}
