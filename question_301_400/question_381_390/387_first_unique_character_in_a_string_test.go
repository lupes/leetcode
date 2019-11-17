package question_381_390

import "testing"

func Test_firstUniqChar(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test", "leetcode", 0},
		{"test", "cc", -1},
		{"test", "loveleetcode", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
