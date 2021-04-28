package question_1651_1660

import (
	"reflect"
	"testing"
)

func Test_decrypt(t *testing.T) {
	tests := []struct {
		code []int
		k    int
		want []int
	}{
		{[]int{5, 7, 1, 4}, 3, []int{12, 10, 16, 13}},
		{[]int{1, 2, 3, 4}, 0, []int{0, 0, 0, 0}},
		{[]int{2, 4, 9, 3}, -2, []int{12, 5, 6, 13}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := decrypt(tt.code, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
