package question_1471_1480

import (
	"fmt"
	"testing"
)

func TestBrowserHistory(t *testing.T) {
	browserHistory := Constructor("leetcode.com")
	browserHistory.Visit("google.com")
	browserHistory.Visit("facebook.com")
	browserHistory.Visit("youtube.com")
	fmt.Printf("%s\n", browserHistory.Back(1))
	fmt.Printf("%s\n", browserHistory.Back(1))
	fmt.Printf("%s\n", browserHistory.Forward(1))
	browserHistory.Visit("linkedin.com")
	fmt.Printf("%s\n", browserHistory.Forward(2))
	fmt.Printf("%s\n", browserHistory.Back(2))
	fmt.Printf("%s\n", browserHistory.Back(7))
}
