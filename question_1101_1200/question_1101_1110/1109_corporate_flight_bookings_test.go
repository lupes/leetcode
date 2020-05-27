package question_1101_1110

import (
	"reflect"
	"testing"
)

func Test_corpFlightBookings(t *testing.T) {
	tests := []struct {
		bookings [][]int
		n        int
		want     []int
	}{
		{[][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5, []int{10, 55, 45, 25, 25}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := corpFlightBookings(tt.bookings, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("corpFlightBookings() = %v, want %v", got, tt.want)
			}
		})
	}
}
