package question_491_500

import (
	"reflect"
	"testing"
)

func Test_constructRectangle(t *testing.T) {
	tests := []struct {
		name string
		area int
		want []int
	}{
		{"test", 1, []int{1, 1}},
		{"test", 2, []int{2, 1}},
		{"test", 4, []int{2, 2}},
		{"test", 6, []int{3, 2}},
		{"test", 15, []int{5, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructRectangle(tt.area); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
