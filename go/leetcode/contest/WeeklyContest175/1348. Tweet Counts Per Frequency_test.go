package contest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TweetCounts struct {
	m map[string]map[int]int
}

func Constructor() TweetCounts {
	return TweetCounts{
		m: make(map[string]map[int]int, 10000),
	}
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	if _, ok := this.m[tweetName]; !ok {
		this.m[tweetName] = make(map[int]int, 10000)
	}
	this.m[tweetName][time]++
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	res := []int{}
	if _, ok := this.m[tweetName]; !ok {
		return res
	}
	interval := 60
	if freq == "hour" {
		interval = 3600
	} else if freq == "day" {
		interval = 86400
	}
	size := ((endTime - startTime) / interval) + 1
	buckets := make([]int, size)

	tweets := this.m[tweetName]
	for tweetTime, cnt := range tweets {
		if tweetTime < startTime || tweetTime > endTime {
			continue
		}
		idx := (tweetTime - startTime) / interval
		buckets[idx] += cnt
	}
	for _, cnt := range buckets {
		res = append(res, cnt)
	}
	return res
}

/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */
func TestTweetCounts(t *testing.T) {
	t.Run(fmt.Sprintf("Example %d", 1), func(t *testing.T) {
		ast := assert.New(t)
		obj := Constructor()
		obj.RecordTweet("tweet3", 0)
		obj.RecordTweet("tweet3", 60)
		obj.RecordTweet("tweet3", 10)
		ast.Equal(obj.GetTweetCountsPerFrequency("minute", "tweet3", 0, 59), []int{2})    // [2]
		ast.Equal(obj.GetTweetCountsPerFrequency("minute", "tweet3", 0, 60), []int{2, 1}) //[2,1]
		obj.RecordTweet("tweet3", 120)
		ast.Equal(obj.GetTweetCountsPerFrequency("hour", "tweet3", 0, 210), []int{4}) //[4]
	})

	t.Run(fmt.Sprintf("Example %d", 2), func(t *testing.T) {
		ast := assert.New(t)
		obj := Constructor()
		obj.RecordTweet("tweet0", 99)
		obj.RecordTweet("tweet1", 80)
		obj.RecordTweet("tweet2", 29)
		obj.RecordTweet("tweet3", 81)
		obj.RecordTweet("tweet4", 56)
		ast.Equal(obj.GetTweetCountsPerFrequency("day", "tweet4", 56, 2667), []int{1}) // [1]
		obj.RecordTweet("tweet1", 56)
		obj.RecordTweet("tweet0", 96)
		obj.RecordTweet("tweet4", 35)
		obj.RecordTweet("tweet1", 51)
		obj.RecordTweet("tweet2", 45)
		obj.RecordTweet("tweet0", 74)
		obj.RecordTweet("tweet1", 63)
	})

}
