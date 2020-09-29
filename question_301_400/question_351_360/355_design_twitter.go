package question_351_360

import (
	"sort"
)

// 355. 设计推特
// https://leetcode-cn.com/problems/design-twitter
// Topics: 堆 设计 哈希表

type Tweet struct {
	Id int
	ts int64
}

type Twitter struct {
	UserMap  map[int]map[int]struct{}
	TweetMap map[int][]Tweet
	id       int64
}

/** Initialize your data structure here. */
func Constructor5() Twitter {
	return Twitter{
		UserMap:  make(map[int]map[int]struct{}, 0),
		TweetMap: make(map[int][]Tweet, 0),
		id:       0,
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	tweets := this.TweetMap[userId]
	tweets = append(tweets, Tweet{
		Id: tweetId,
		ts: this.id,
	})
	this.id++
	if len(tweets) > 10 {
		tweets = tweets[1:]
	}
	this.TweetMap[userId] = tweets
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	var tweets []Tweet
	tweets = append(tweets, this.TweetMap[userId]...)
	for id := range this.UserMap[userId] {
		tweets = append(tweets, this.TweetMap[id]...)
	}
	sort.Slice(tweets, func(i, j int) bool {
		return tweets[i].ts > tweets[j].ts
	})
	var res []int
	for i := 0; i < len(tweets) && i < 10; i++ {
		res = append(res, tweets[i].Id)
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	_, ok := this.UserMap[followerId]
	if !ok {
		this.UserMap[followerId] = make(map[int]struct{}, 0)
	}
	this.UserMap[followerId][followeeId] = struct{}{}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.UserMap[followerId], followeeId)
}
