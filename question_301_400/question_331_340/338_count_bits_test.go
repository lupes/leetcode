package question_331_340

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want []int
	}{
		{"test", 2, []int{0, 1, 1}},
		{"test", 5, []int{0, 1, 1, 2, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
