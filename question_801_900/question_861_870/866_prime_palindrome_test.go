package question_861_870

import (
	"testing"
)

func Test_primePalindrome(t *testing.T) {
	tests := []struct {
		N    int
		want int
	}{
		{1, 2},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 5},
		{6, 7},
		{7, 7},
		{8, 11},
		{11, 11},
		{12, 101},
		{13, 101},
		{9989899, 9989899},
		{8570_9140, 100030001},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := primePalindrome(tt.N); got != tt.want {
				t.Errorf("primePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
