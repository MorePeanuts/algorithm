---
link: https://leetcode.com/problems/design-twitter/
tags:
  - Medium
  - Design
  - Hash_Table
  - Linked_List
  - Heap_(Priority_Queue)
---
## Description
Design a simplified version of Twitter where users can post tweets, follow/unfollow another user, and is able to see the `10` most recent tweets in the user's news feed.

Implement the `Twitter` class:

- `Twitter()` Initializes your twitter object.
- `void postTweet(int userId, int tweetId)` Composes a new tweet with ID `tweetId` by the user `userId`. Each call to this function will be made with a unique `tweetId`.
- `List<Integer> getNewsFeed(int userId)` Retrieves the `10` most recent tweet IDs in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user themself. Tweets must be **ordered from most recent to least recent**.
- `void follow(int followerId, int followeeId)` The user with ID `followerId` started following the user with ID `followeeId`.
- `void unfollow(int followerId, int followeeId)` The user with ID `followerId` started unfollowing the user with ID `followeeId`.

**Example 1:**

```
Input
["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"]
[[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
Output
[null, null, [5], null, null, [6, 5], null, [5]]

Explanation
Twitter twitter = new Twitter();
twitter.postTweet(1, 5); // User 1 posts a new tweet (id = 5).
twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
twitter.follow(1, 2);    // User 1 follows user 2.
twitter.postTweet(2, 6); // User 2 posts a new tweet (id = 6).
twitter.getNewsFeed(1);  // User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
twitter.unfollow(1, 2);  // User 1 unfollows user 2.
twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.
```

**Constraints:**

- `1 <= userId, followerId, followeeId <= 500`
- `0 <= tweetId <= 104`
- All the tweets have **unique** IDs.
- At most `3 * 104` calls will be made to `postTweet`, `getNewsFeed`, `follow`, and `unfollow`.
- A user cannot follow himself.

## Solution
### Approach 1
**Hash Map + Max Heap**: Maintain per-user tweet lists and follow relationships via hash maps, assign a global incrementing timestamp to each tweet, and use a max heap to merge and retrieve the 10 most recent tweets from all relevant users in reverse chronological order.

**Principle:**
Each tweet is assigned a globally incrementing timestamp on posting. Hash maps store each user's tweet list and follow set. To retrieve the news feed, the last up to 10 tweets from the user and each followee are collected, pushed into a max heap keyed by timestamp, and the top 10 are popped to produce the result in reverse chronological order.

**Steps:**
1. `PostTweet`: Append the tweet to the user's tweet list with the current global timestamp, then increment the timestamp.
2. `GetNewsFeed`: Collect the last up to 10 tweets from the user and each followee, build a max heap, and pop the top 10 entries.
3. `Follow`: Add the followee's ID to the follower's followee set.
4. `Unfollow`: Remove the followee's ID from the follower's followee set.

**Example:**
- User 1 posts tweet 5 (time 0), user 2 posts tweet 6 (time 1).
- After user 1 follows user 2 and calls getNewsFeed: collect tweets from user 1 [5] and user 2 [6], push into max heap, pop in reverse chronological order → [6, 5].
- After user 1 unfollows user 2 and calls getNewsFeed: only user 1's own tweets remain → [5].

```embed-go
PATH: "vault://leetcode/0301-0400/0355_design_twitter/solution.go"
TITLE: "leetcode 355. Design Twitter"
```
