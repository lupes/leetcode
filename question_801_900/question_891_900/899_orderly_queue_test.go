package question_891_900

import (
	"testing"
)

func Test_orderlyQueue(t *testing.T) {
	tests := []struct {
		S    string
		K    int
		want string
	}{
		{"cba", 1, "acb"},
		{"baaca", 3, "aaabc"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := orderlyQueue(tt.S, tt.K); got != tt.want {
				t.Errorf("orderlyQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
