package question_831_840

import (
	"testing"
)

func Test_maskPII(t *testing.T) {
	tests := []struct {
		S    string
		want string
	}{
		{"LeetCode@LeetCode.com", "l*****e@leetcode.com"},
		{"AB@qq.com", "a*****b@qq.com"},
		{"1(234)567-890", "***-***-7890"},
		{"11(234)567-890", "+*-***-***-7890"},
		{"861(234)567-890", "+**-***-***-7890"},
		{"3861(234)567-890", "+***-***-***-7890"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maskPII(tt.S); got != tt.want {
				t.Errorf("maskPII() = %v, want %v", got, tt.want)
			}
		})
	}
}
