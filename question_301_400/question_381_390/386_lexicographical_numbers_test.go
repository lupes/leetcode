package question_381_390

import (
	"reflect"
	"testing"
)

func Test_lexicalOrder(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{1, []int{1}},
		{11, []int{1, 10, 11, 2, 3, 4, 5, 6, 7, 8, 9}},
		{13, []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := lexicalOrder(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexicalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
