package question_561_570

import "testing"

func Test_arrayPairSum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 4, 3, 2}, 4},
		{[]int{1, 4, 3, 2, 5, 4}, 8},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := arrayPairSum(tt.nums); got != tt.want {
				t.Errorf("arrayPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
