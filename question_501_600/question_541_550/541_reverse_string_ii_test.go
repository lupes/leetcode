package question_541_550

import "testing"

func Test_reverseStr(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want string
	}{
		{"t", 1, "t"},
		{"t", 2, "t"},
		{"te", 2, "et"},
		{"tes", 2, "ets"},
		{"test", 2, "etst"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := reverseStr(tt.s, tt.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
