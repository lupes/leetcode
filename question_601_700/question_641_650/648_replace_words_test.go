package question_641_650

import (
	"testing"
)

func Test_replaceWords(t *testing.T) {
	tests := []struct {
		dict     []string
		sentence string
		want     string
	}{
		{[]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery", "the cat was rat by the bat"},
		{[]string{"cat", "bat", "rat"}, "the cattle was rattled by the bat", "the cat was rat by the bat"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := replaceWords(tt.dict, tt.sentence); got != tt.want {
				t.Errorf("replaceWords() = [%v], want [%v]", got, tt.want)
			}
		})
	}
}
