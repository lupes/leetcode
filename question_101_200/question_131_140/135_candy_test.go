package question_131_140

import (
	"testing"
)

func Test_candy(t *testing.T) {
	tests := []struct {
		ratings []int
		want    int
	}{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := candy(tt.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
