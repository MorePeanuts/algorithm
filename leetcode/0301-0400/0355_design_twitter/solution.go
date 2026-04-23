// Package leetcode0355 solves LeetCode 355. Design Twitter
package leetcode0355

import "container/heap"

type Tweet struct {
	id, time int
}

type TweetHeap []Tweet

func (h TweetHeap) Len() int           { return len(h) }
func (h TweetHeap) Less(i, j int) bool { return h[i].time > h[j].time }
func (h TweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *TweetHeap) Push(x any)        { *h = append(*h, x.(Tweet)) }

func (h *TweetHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Set map[int]bool

func (s Set) Add(x int) {
	s[x] = true
}

func (s Set) Find(x int) bool {
	_, ok := s[x]
	return ok
}

func (s Set) Remove(x int) {
	delete(s, x)
}

type Twitter struct {
	userTweet    map[int][]Tweet
	userFollowee map[int]Set
	time         int
}

func Constructor() Twitter {
	return Twitter{make(map[int][]Tweet), make(map[int]Set), 0}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.userTweet[userId] = append(this.userTweet[userId], Tweet{tweetId, this.time})
	this.time++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	h := make(TweetHeap, 0, 10*len(this.userFollowee[userId])+10)
	// Tweets are appended chronologically, so the most recent are at the end.
	// Take the last up to 10 tweets from each user.
	takeLast := func(tweets []Tweet) []Tweet {
		if len(tweets) <= 10 {
			return tweets
		}
		return tweets[len(tweets)-10:]
	}
	h = append(h, takeLast(this.userTweet[userId])...)
	for followeeId := range this.userFollowee[userId] {
		h = append(h, takeLast(this.userTweet[followeeId])...)
	}
	heap.Init(&h)
	res := make([]int, 0, 10)
	for range min(10, len(h)) {
		tweet := heap.Pop(&h).(Tweet)
		res = append(res, tweet.id)
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.userFollowee[followerId] == nil {
		this.userFollowee[followerId] = make(Set)
	}
	this.userFollowee[followerId].Add(followeeId)
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	this.userFollowee[followerId].Remove(followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
