package problem

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type uid = int

type tid = int

type Twitter struct {
	users map[uid]*User
	idx   int
}

type User struct {
	id       uid
	followee map[uid]struct{} // 关注列表
	tweets   []Tweet          // 推送列表
}

type Tweet struct {
	id  tid
	idx int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		users: make(map[uid]*User, 2048),
		idx:   0,
	}
}

func (this *Twitter) isUserValid(userId int) bool {
	return this.users[userId] != nil
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if !this.isUserValid(userId) {
		this.users[userId] = &User{
			id:       userId,
			followee: make(map[uid]struct{}, 2048),
			tweets:   make([]Tweet, 0, 2048),
		}
	}
	this.users[userId].tweets = append(this.users[userId].tweets, Tweet{
		id:  tweetId,
		idx: this.idx,
	})
	this.idx++
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	if !this.isUserValid(userId) {
		return []int{}
	}
	user := this.users[userId]
	tweets := append([]Tweet{}, user.tweets...)
	for folId, _ := range user.followee {
		if this.users[folId] == nil || this.users[folId].tweets == nil {
			continue
		}
		tweets = append(tweets, this.users[folId].tweets...)
	}
	sort.Slice(tweets, func(i, j int) bool {
		return tweets[i].idx > tweets[j].idx
	})
	l := 10
	if l_ := len(tweets); l_ < l {
		l = l_
	}
	res := make([]int, l)
	for i := 0; i < l; i++ {
		res[i] = tweets[i].id
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if !this.isUserValid(followerId) {
		this.users[followerId] = &User{
			id:       followerId,
			followee: make(map[uid]struct{}, 2048),
			tweets:   make([]Tweet, 0, 2048),
		}
	}
	this.users[followerId].followee[followeeId] = struct{}{}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if !this.isUserValid(followerId) {
		this.users[followerId] = &User{
			id:       followerId,
			followee: make(map[uid]struct{}, 2048),
			tweets:   make([]Tweet, 0, 2048),
		}
	}
	if _, ok := this.users[followerId].followee[followeeId]; ok {
		delete(this.users[followerId].followee, followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
func TestTwitter(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		ast := assert.New(t)
		obj := Constructor()
		obj.PostTweet(1, 5)
		ast.Equal([]int{5}, obj.GetNewsFeed(1))
		obj.Follow(1, 2)
		ast.Equal([]int{}, obj.GetNewsFeed(2))
		ast.Equal([]int{5}, obj.GetNewsFeed(1))
		obj.PostTweet(2, 6)
		ast.Equal([]int{6, 5}, obj.GetNewsFeed(1))
		obj.Unfollow(1, 2)
		ast.Equal([]int{5}, obj.GetNewsFeed(1))
	})
	t.Run("Example 2", func(t *testing.T) {
		// ["Twitter","postTweet","follow","getNewsFeed"]
		// [[],[1,5],[1,1],[1]]
		ast := assert.New(t)
		obj := Constructor()
		obj.PostTweet(1, 5)
		obj.Follow(1, 1)
		ast.Equal([]int{5}, obj.GetNewsFeed(1))
	})
}
