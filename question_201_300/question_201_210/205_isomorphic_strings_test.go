package question_201_210

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"test", "", "", true},
		{"test", "a", "b", true},
		{"test", "ab", "bb", false},
		{"test", "abb", "cbb", true},
		{"test", "paper", "title", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.s, tt.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
