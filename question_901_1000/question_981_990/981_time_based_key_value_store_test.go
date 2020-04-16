package question_981_990

import (
	"testing"
)

func TestTimeMap(t *testing.T) {
	timeMap := Constructor()
	timeMap.Set("foo", "bar", 1)

	if got := timeMap.Get("foo", 1); got != "bar" {
		t.Errorf("Get() = %v", got)
	}

	if got := timeMap.Get("foo", 3); got != "bar" {
		t.Errorf("Get() = %v", got)
	}

	timeMap.Set("foo", "bar2", 4)
	if got := timeMap.Get("foo", 4); got != "bar2" {
		t.Errorf("Get() = %v", got)
	}

	if got := timeMap.Get("foo", 5); got != "bar2" {
		t.Errorf("Get() = %v", got)
	}

	timeMap = Constructor()
	timeMap.Set("love", "high", 10)
	timeMap.Set("love", "low", 20)

	if got := timeMap.Get("love", 5); got != "" {
		t.Errorf("Get() = %v", got)
	}

	if got := timeMap.Get("love", 10); got != "high" {
		t.Errorf("Get() = %v", got)
	}

	if got := timeMap.Get("love", 15); got != "high" {
		t.Errorf("Get() = %v", got)
	}

	if got := timeMap.Get("love", 20); got != "low" {
		t.Errorf("Get() = %v", got)
	}

	if got := timeMap.Get("love", 25); got != "low" {
		t.Errorf("Get() = %v", got)
	}
}
