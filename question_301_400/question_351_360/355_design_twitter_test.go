package question_351_360

import (
	"reflect"
	"testing"
)

func TestTwitter(t *testing.T) {
	twitter := Constructor5()
	twitter.PostTweet(1, 5)
	if got, want := twitter.GetNewsFeed(1), []int{5}; !reflect.DeepEqual(got, want) {
		t.Errorf("GetNewsFeed() = %v, want %v", got, want)
	}

	twitter.Follow(1, 2)
	if got, want := twitter.GetNewsFeed(1), []int{5}; !reflect.DeepEqual(got, want) {
		t.Errorf("GetNewsFeed() = %v, want %v", got, want)
	}

	twitter.PostTweet(2, 6)
	if got, want := twitter.GetNewsFeed(1), []int{6, 5}; !reflect.DeepEqual(got, want) {
		t.Errorf("GetNewsFeed() = %v, want %v", got, want)
	}

	twitter.Unfollow(1, 2)
	if got, want := twitter.GetNewsFeed(1), []int{5}; !reflect.DeepEqual(got, want) {
		t.Errorf("GetNewsFeed() = %v, want %v", got, want)
	}
}
