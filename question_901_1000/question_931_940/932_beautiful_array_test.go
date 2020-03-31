package question_931_940

import (
	"reflect"
	"testing"
)

func Test_beautifulArray(t *testing.T) {
	tests := []struct {
		N    int
		want []int
	}{
		{1, []int{1}},
		{2, []int{1, 2}},
		{3, []int{1, 3, 2}},
		{4, []int{1, 3, 2, 4}},
		{5, []int{1, 5, 3, 2, 4}},
		{6, []int{1, 5, 3, 2, 6, 4}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := beautifulArray(tt.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("beautifulArray() = %#v, want %v", got, tt.want)
			}
		})
	}
}
