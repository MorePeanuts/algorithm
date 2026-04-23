---
link: https://leetcode.cn/problems/design-twitter/
tags:
  - 中等
  - 设计
  - 哈希表
  - 链表
  - 堆（优先队列）
---
## 题目描述
设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近 `10` 条推文。

实现 `Twitter` 类：

- `Twitter()` 初始化简易版推特对象
- `void postTweet(int userId, int tweetId)` 根据给定的 `tweetId` 和 `userId` 创建一条新推文。每次调用此函数都会使用一个不同的 `tweetId` 。
- `List<Integer> getNewsFeed(int userId)` 检索当前用户新闻推送中最近  `10` 条推文的 ID 。新闻推送中的每一项都必须是由用户关注的人或者是用户自己发布的推文。推文必须 **按照时间顺序由最近到最远排序** 。
- `void follow(int followerId, int followeeId)` ID 为 `followerId` 的用户开始关注 ID 为 `followeeId` 的用户。
- `void unfollow(int followerId, int followeeId)` ID 为 `followerId` 的用户不再关注 ID 为 `followeeId` 的用户。

**示例：**

```
输入
["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"]
[[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
输出
[null, null, [5], null, null, [6, 5], null, [5]]

解释
Twitter twitter = new Twitter();
twitter.postTweet(1, 5); // 用户 1 发送了一条新推文 (用户 id = 1, 推文 id = 5)
twitter.getNewsFeed(1);  // 用户 1 的获取推文应当返回一个列表，其中包含一个 id 为 5 的推文
twitter.follow(1, 2);    // 用户 1 关注了用户 2
twitter.postTweet(2, 6); // 用户 2 发送了一个新推文 (推文 id = 6)
twitter.getNewsFeed(1);  // 用户 1 的获取推文应当返回一个列表，其中包含两个推文，id 分别为 -> [6, 5] 。推文 id 6 应当在推文 id 5 之前，因为它是在 5 之后发送的
twitter.unfollow(1, 2);  // 用户 1 取消关注了用户 2
twitter.getNewsFeed(1);  // 用户 1 获取推文应当返回一个列表，其中包含一个 id 为 5 的推文。因为用户 1 已经不再关注用户 2
```

**提示：**

- `1 <= userId, followerId, followeeId <= 500`
- `0 <= tweetId <= 104`
- 所有推特的 ID 都互不相同
- `postTweet`、`getNewsFeed`、`follow` 和 `unfollow` 方法最多调用 `3 * 104` 次
- 用户不能关注自己

## 题目解析
### 解法1
**哈希表 + 最大堆**：用哈希表维护用户的推文列表和关注关系，用全局时间戳标记推文先后顺序，获取新闻推送时利用最大堆按时间倒序合并所有相关用户的推文。

**原理：**
为每条推文分配全局递增时间戳，用哈希表分别存储每个用户的推文列表和关注集合；获取新闻推送时，从用户自身及所有关注者中各取最近 10 条推文，入最大堆后依次弹出堆顶，即可得到按时间倒序排列的前 10 条。

**步骤：**
1. `PostTweet`：将推文追加到对应用户的推文列表，记录全局递增时间戳
2. `GetNewsFeed`：收集用户自身及所有关注者各最近 10 条推文，构建最大堆，依次弹出堆顶取前 10 条
3. `Follow`：将关注对象 ID 加入用户的关注集合
4. `Unfollow`：从用户的关注集合中移除关注对象 ID

**示例：**
- 用户 1 发推文 5（时间 0），用户 2 发推文 6（时间 1）
- 用户 1 关注用户 2 后获取新闻推送：从用户 1 的推文 [5] 和用户 2 的推文 [6] 中取最近 10 条，入堆后按时间倒序弹出 → [6, 5]
- 用户 1 取关用户 2 后获取新闻推送：仅从用户 1 自身推文获取 → [5]

```embed-go
PATH: "vault://leetcode/0301-0400/0355_design_twitter/solution.go"
TITLE: "leetcode 0355.设计推特"
```
