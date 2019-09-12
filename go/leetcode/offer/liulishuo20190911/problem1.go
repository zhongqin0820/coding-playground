package main

import (
	"fmt"
	"math"
	"sort"
)

type pairs []struct {
	index int
	diff  int
}

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pairs) Less(i, j int) bool {
	return p[i].diff > p[j].diff
}

// 参考讨论：https://www.nowcoder.com/discuss/256204
func problem1(n, a, b int, c, d []int) int {
	res := 0
	pair := make(pairs, n)
	for i := 0; i < n; i++ {
		pair[i].index = i
		pair[i].diff = int(math.Abs(float64(c[i] - d[i])))
	}
	sort.Sort(pair)
	// 先派送差值大的人
	for i := 0; i < n; i++ {
		index := pair[i].index
		switch {
		case (c[index] > d[index] && b > 0) || (a <= 0):
			// 乙地代价低且乙地还有礼物可以派送 || 甲地礼物派完了
			res += d[index]
			b--
		case (c[index] < d[index] && a > 0) || (b <= 0):
			// 甲地代价低且甲地还有礼物可以派送 || 已地礼物派完了
			res += c[index]
			a--
		}
	}
	return res
}

func problemWrapper1() {
	n, a, b, c, d := input1()
	output := problem1(n, a, b, c, d)
	fmt.Println(output)
}

func input1() (int, int, int, []int, []int) {
	n, a, b := 0, 0, 0
	fmt.Scanf("%d\n%d %d", &n, &a, &b)
	c, d := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &c[i], &d[i])
	}
	return n, a, b, c, d
}

func main() {
	problemWrapper1()
}
