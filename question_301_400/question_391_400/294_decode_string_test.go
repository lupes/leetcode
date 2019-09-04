package question_391_400

import "testing"

func Test_decodeString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"test", "3[a]2[bc]", "aaabcbc"},
		{"test", "3[a2[c]]", "accaccacc"},
		{"test", "3[a2[cd4[ef]]]", "acdefefefefcdefefefefacdefefefefcdefefefefacdefefefefcdefefefef"},
		{"test", "2[abc]3[cd]ef", "abcabccdcdcdef"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
