package question_381_390

import (
	"testing"
)

func Test_lengthLongestPath(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext", 20},
		{"a", 0},
		{"file1.txt\nfile2.txt\nlongfile.txt", 12},
		{"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext", 32},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := lengthLongestPath(tt.input); got != tt.want {
				t.Errorf("lengthLongestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
