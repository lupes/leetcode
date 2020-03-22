package question_931_940

import (
	"testing"
)

func TestRecentCounter_Ping(t *testing.T) {
	this := Constructor()

	if got := this.Ping(1); got != 1 {
		t.Errorf("want 1 got %d", got)
	}

	if got := this.Ping(100); got != 2 {
		t.Errorf("want 1 got %d", got)
	}

	if got := this.Ping(3001); got != 3 {
		t.Errorf("want 1 got %d", got)
	}

	if got := this.Ping(3002); got != 3 {
		t.Errorf("want 1 got %d", got)
	}
}
