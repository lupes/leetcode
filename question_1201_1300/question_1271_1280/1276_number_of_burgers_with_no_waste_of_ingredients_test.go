package question_1261_1270

import (
	"reflect"
	"testing"
)

func Test_numOfBurgers(t *testing.T) {
	tests := []struct {
		tomatoSlices int
		cheeseSlices int
		want         []int
	}{
		{16, 7, []int{1, 6}},
		{17, 4, nil},
		{4, 17, nil},
		{0, 0, []int{0, 0}},
		{2, 1, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := numOfBurgers(tt.tomatoSlices, tt.cheeseSlices)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numOfBurgers() = %v, want %v", got, tt.want)
			}
		})
	}
}
