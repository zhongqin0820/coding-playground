package contest

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-days-between-two-dates/
var temp = "2016-01-01"

func daysBetweenDates(date1 string, date2 string) int {
	t1 := parse(date1)
	t2 := parse(date2)
	h1 := time.Since(t1).Hours()
	h2 := time.Since(t2).Hours()
	if h1 > h2 {
		h1, h2 = h2, h1
	}
	return int((h2 - h1 + 1) / 24)
}

func parse(s string) time.Time {
	s1 := strings.Split(s, "-")
	y1, _ := strconv.Atoi(s1[0])
	m1, _ := strconv.Atoi(s1[1])
	d1, _ := strconv.Atoi(s1[2])
	t1 := time.Date(y1, time.Month(m1), d1, 0, 0, 0, 0, time.UTC)
	return t1
}

func TestDaysBetweenDates(t *testing.T) {
	tests := []struct {
		date1  string
		date2  string
		output int
	}{
		{"2019-06-29", "2019-06-30", 1},
		{"2020-01-15", "2019-12-31", 15},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, daysBetweenDates(ts.date1, ts.date2))
		})
	}
}
