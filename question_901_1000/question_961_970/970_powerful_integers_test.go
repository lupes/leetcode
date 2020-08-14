package question_961_970

import (
	"reflect"
	"testing"
)

func Test_powerfulIntegers(t *testing.T) {
	tests := []struct {
		x     int
		y     int
		bound int
		want  []int
	}{
		{2, 3, 10, []int{2, 3, 4, 5, 7, 9, 10}},
		{2, 1, 10, []int{2, 3, 4, 5, 7, 9, 10}},
		{2, 3, 1000000, []int{2, 3, 4, 5, 7, 9, 10}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := powerfulIntegers(tt.x, tt.y, tt.bound); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("powerfulIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
