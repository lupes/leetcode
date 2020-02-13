package question_821_830

import (
	"testing"
)

func Test_numFriendRequests(t *testing.T) {
	tests := []struct {
		ages []int
		want int
	}{
		{[]int{16, 16}, 2},
		{[]int{16, 16, 16}, 6},
		{[]int{16, 17, 18}, 2},
		{[]int{20, 30, 100, 110, 120}, 3},
		{[]int{101, 56, 69, 48, 30}, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numFriendRequests(tt.ages); got != tt.want {
				t.Errorf("numFriendRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}
