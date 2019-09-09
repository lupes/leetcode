package question_231_240

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"test", []int{1, 2}, []int{2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
