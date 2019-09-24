package question_481_490

import "testing"

func Test_licenseKeyFormatting(t *testing.T) {
	tests := []struct {
		name string
		S    string
		K    int
		want string
	}{
		{"test", "5F3Z-2e-9-w", 4, "5F3Z-2E9W"},
		{"test", "2-5g-3-J", 2, "2-5G-3J"},
		{"test", "2-5g-3-J", 6, "25G3J"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := licenseKeyFormatting(tt.S, tt.K); got != tt.want {
				t.Errorf("licenseKeyFormatting() = %v, want %v", got, tt.want)
			}
		})
	}
}
