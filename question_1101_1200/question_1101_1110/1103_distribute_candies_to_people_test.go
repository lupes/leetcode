package question_1101_1110

import (
	"reflect"
	"testing"
)

func Test_distributeCandies(t *testing.T) {
	tests := []struct {
		candies    int
		num_people int
		want       []int
	}{
		{7, 4, []int{1, 2, 3, 1}},
		{100000000, 1000, []int{1, 2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := distributeCandies(tt.candies, tt.num_people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
