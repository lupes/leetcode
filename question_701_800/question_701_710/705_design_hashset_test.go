package question_701_710

import (
	"testing"
)

func TestConstructor3(t *testing.T) {
	set := Constructor3()
	set.Add(9)
	set.Remove(19)

	set.Add(14)
	set.Remove(19)
	set.Remove(9)

	if got, want := set.Contains(1), true; got != want {
		t.Errorf("Contains(1) = %v, want %v", got, want)
	}

	if got, want := set.Contains(2), true; got != want {
		t.Errorf("Contains(2) = %v, want %v", got, want)
	}

	if got, want := set.Contains(3), false; got != want {
		t.Errorf("Contains(3) = %v, want %v", got, want)
	}

	set.Add(2)

	if got, want := set.Contains(2), true; got != want {
		t.Errorf("Contains(2) = %v, want %v", got, want)
	}

	set.Remove(2)
	if got, want := set.Contains(2), false; got != want {
		t.Errorf("Contains(2) = %v, want %v", got, want)
	}
}
