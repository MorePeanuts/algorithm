package leetcode0355

import (
	"slices"
	"testing"
)

func TestTwitterExamples(t *testing.T) {
	t.Run("example_1", func(t *testing.T) {
		twitter := Constructor()
		twitter.PostTweet(1, 5)
		if got := twitter.GetNewsFeed(1); !slices.Equal(got, []int{5}) {
			t.Errorf("GetNewsFeed(1) = %v, want [5]", got)
		}
		twitter.Follow(1, 2)
		twitter.PostTweet(2, 6)
		if got := twitter.GetNewsFeed(1); !slices.Equal(got, []int{6, 5}) {
			t.Errorf("GetNewsFeed(1) = %v, want [6, 5]", got)
		}
		twitter.Unfollow(1, 2)
		if got := twitter.GetNewsFeed(1); !slices.Equal(got, []int{5}) {
			t.Errorf("GetNewsFeed(1) = %v, want [5]", got)
		}
	})
}

func TestTwitter(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{"single_user_single_tweet", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(1, 100)
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{100}) {
				t.Errorf("got %v, want [100]", got)
			}
		}},
		{"single_user_multiple_tweets", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(1, 10)
			tw.PostTweet(1, 20)
			tw.PostTweet(1, 30)
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{30, 20, 10}) {
				t.Errorf("got %v, want [30 20 10]", got)
			}
		}},
		{"news_feed_max_10", func(t *testing.T) {
			tw := Constructor()
			for i := 1; i <= 15; i++ {
				tw.PostTweet(1, i)
			}
			if got := tw.GetNewsFeed(1); len(got) != 10 {
				t.Errorf("got %d tweets, want 10", len(got))
			}
			// Most recent first
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6}) {
				t.Errorf("got %v, want [15 14 13 12 11 10 9 8 7 6]", got)
			}
		}},
		{"follow_then_post", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(1, 1)
			tw.Follow(2, 1)
			if got := tw.GetNewsFeed(2); !slices.Equal(got, []int{1}) {
				t.Errorf("got %v, want [1]", got)
			}
		}},
		{"unfollow_removes_tweets", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(1, 1)
			tw.Follow(2, 1)
			tw.Unfollow(2, 1)
			if got := tw.GetNewsFeed(2); len(got) != 0 {
				t.Errorf("got %v, want []", got)
			}
		}},
		{"unfollow_nonexistent_followee", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(1, 1)
			tw.Unfollow(1, 999) // no-op, should not panic
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{1}) {
				t.Errorf("got %v, want [1]", got)
			}
		}},
		{"follow_same_user_twice", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(2, 5)
			tw.Follow(1, 2)
			tw.Follow(1, 2) // duplicate follow
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{5}) {
				t.Errorf("got %v, want [5]", got)
			}
		}},
		{"empty_feed_no_tweets_no_follows", func(t *testing.T) {
			tw := Constructor()
			if got := tw.GetNewsFeed(1); len(got) != 0 {
				t.Errorf("got %v, want []", got)
			}
		}},
		{"feed_includes_own_tweets", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(1, 10)
			tw.Follow(1, 2)
			tw.PostTweet(2, 20)
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{20, 10}) {
				t.Errorf("got %v, want [20 10]", got)
			}
		}},
		{"multiple_followees_merged_feed", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(2, 1)
			tw.PostTweet(3, 2)
			tw.PostTweet(4, 3)
			tw.Follow(1, 2)
			tw.Follow(1, 3)
			tw.Follow(1, 4)
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{3, 2, 1}) {
				t.Errorf("got %v, want [3 2 1]", got)
			}
		}},
		{"interleaved_posts_from_multiple_users", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(1, 100)
			tw.PostTweet(2, 200)
			tw.PostTweet(1, 101)
			tw.PostTweet(2, 201)
			tw.Follow(1, 2)
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{201, 101, 200, 100}) {
				t.Errorf("got %v, want [201 101 200 100]", got)
			}
		}},
		{"tweet_id_zero", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(1, 0)
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{0}) {
				t.Errorf("got %v, want [0]", got)
			}
		}},
		{"max_tweet_id", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(1, 10000)
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{10000}) {
				t.Errorf("got %v, want [10000]", got)
			}
		}},
		{"min_user_id", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(1, 50)
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{50}) {
				t.Errorf("got %v, want [50]", got)
			}
		}},
		{"max_user_id", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(500, 50)
			if got := tw.GetNewsFeed(500); !slices.Equal(got, []int{50}) {
				t.Errorf("got %v, want [50]", got)
			}
		}},
		{"follow_multiple_then_unfollow_one", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(2, 10)
			tw.PostTweet(3, 20)
			tw.Follow(1, 2)
			tw.Follow(1, 3)
			tw.Unfollow(1, 2)
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{20}) {
				t.Errorf("got %v, want [20]", got)
			}
		}},
		{"followee_with_no_tweets", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(1, 1)
			tw.Follow(1, 2) // user 2 has no tweets
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{1}) {
				t.Errorf("got %v, want [1]", got)
			}
		}},
		{"get_news_feed_without_following", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(1, 10)
			tw.PostTweet(1, 20)
			// User 1 not following anyone, feed only contains own tweets
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{20, 10}) {
				t.Errorf("got %v, want [20 10]", got)
			}
		}},
		{"multiple_twitter_instances_independent", func(t *testing.T) {
			tw1 := Constructor()
			tw2 := Constructor()
			tw1.PostTweet(1, 100)
			if got := tw2.GetNewsFeed(1); len(got) != 0 {
				t.Errorf("got %v, want []", got)
			}
		}},
		{"follow_unfollow_refollow", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(2, 42)
			tw.Follow(1, 2)
			tw.Unfollow(1, 2)
			tw.Follow(1, 2) // re-follow
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{42}) {
				t.Errorf("got %v, want [42]", got)
			}
		}},
		{"exactly_10_tweets_returns_all", func(t *testing.T) {
			tw := Constructor()
			for i := 1; i <= 10; i++ {
				tw.PostTweet(1, i)
			}
			if got := tw.GetNewsFeed(1); len(got) != 10 {
				t.Errorf("got %d tweets, want 10", len(got))
			}
		}},
		{"9_tweets_returns_all", func(t *testing.T) {
			tw := Constructor()
			for i := 1; i <= 9; i++ {
				tw.PostTweet(1, i)
			}
			if got := tw.GetNewsFeed(1); len(got) != 9 {
				t.Errorf("got %d tweets, want 9", len(got))
			}
		}},
		{"feed_from_followee_with_many_tweets", func(t *testing.T) {
			tw := Constructor()
			for i := 1; i <= 15; i++ {
				tw.PostTweet(2, i)
			}
			tw.Follow(1, 2)
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6}) {
				t.Errorf("got %v, want [15 14 13 12 11 10 9 8 7 6]", got)
			}
		}},
		{"merged_feed_caps_at_10", func(t *testing.T) {
			tw := Constructor()
			for i := 1; i <= 8; i++ {
				tw.PostTweet(1, i)
			}
			for i := 100; i <= 105; i++ {
				tw.PostTweet(2, i)
			}
			tw.Follow(1, 2)
			got := tw.GetNewsFeed(1)
			if len(got) != 10 {
				t.Errorf("got %d tweets, want 10", len(got))
			}
			// Most recent 10: user2's 105-100, then user1's 8
			if !slices.Equal(got, []int{105, 104, 103, 102, 101, 100, 8, 7, 6, 5}) {
				t.Errorf("got %v", got)
			}
		}},
		{"self_tweet_always_visible", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(1, 77)
			// Even without following anyone, own tweets appear
			if got := tw.GetNewsFeed(1); !slices.Equal(got, []int{77}) {
				t.Errorf("got %v, want [77]", got)
			}
		}},
		{"many_users_all_follow_one", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(1, 999)
			for uid := 2; uid <= 10; uid++ {
				tw.Follow(uid, 1)
				if got := tw.GetNewsFeed(uid); !slices.Equal(got, []int{999}) {
					t.Errorf("user %d: got %v, want [999]", uid, got)
				}
			}
		}},
		{"post_after_unfollow_not_in_feed", func(t *testing.T) {
			tw := Constructor()
			tw.PostTweet(2, 1)
			tw.Follow(1, 2)
			tw.Unfollow(1, 2)
			tw.PostTweet(2, 2)
			if got := tw.GetNewsFeed(1); len(got) != 0 {
				t.Errorf("got %v, want []", got)
			}
		}},
	}

	for _, tst := range tests {
		t.Run(tst.name, tst.fn)
	}
}
