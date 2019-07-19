package problem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/heaters/

// Winter is coming! Your first job during the contest is to design a standard heater with fixed warm radius to warm all the houses.

// Now, you are given positions of houses and heaters on a horizontal line, find out minimum radius of heaters so that all houses could be covered by those heaters.

// So, your input will be the positions of houses and heaters seperately, and your expected output will be the minimum radius standard of heaters.

// Note:
// Numbers of houses and heaters you are given are non-negative and will not exceed 25000.
// Positions of houses and heaters you are given are non-negative and will not exceed 10^9.
// As long as a house is in the heaters' warm radius range, it can be warmed.
// All the heaters follow your radius standard and the warm radius will the same.

func findRadius(houses []int, heaters []int) int {
	if len(houses) == 0 {
		return 0
	}
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var max = func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	//
	res := 0
	sort.Ints(houses)
	sort.Ints(heaters)
	iHeater := sort.SearchInts(heaters, houses[0])
	//
	for iHouse := 0; iHouse < len(houses); iHouse++ {
		for iHeater < len(heaters) && houses[iHouse] > heaters[iHeater] {
			iHeater++
		}
		if iHeater == len(heaters) {
			return max(res, houses[len(houses)-1]-heaters[iHeater-1])
		}

		left := 1<<31 - 1
		if 0 <= iHeater-1 {
			left = houses[iHouse] - heaters[iHeater-1]
		}
		right := heaters[iHeater] - houses[iHouse]

		res = max(res, min(left, right))
	}
	return res
}

func findRadiusAdv(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	var dist = func(a, b int) int {
		if d := a - b; d < 0 {
			return -d
		} else {
			return d
		}
	}

	hi := 0 // closest index of heater
	maxDist := 0
	for i := 0; i < len(houses); i++ {
		house := houses[i]
		for hi < len(heaters)-1 && dist(house, heaters[hi]) >= dist(house, heaters[hi+1]) {
			hi++
		}

		if curDist := dist(house, heaters[hi]); curDist > maxDist {
			maxDist = curDist
		}
	}
	return maxDist
}

func TestFindRadius(t *testing.T) {
	tests := []struct {
		houses, heaters []int
		output          int
	}{
		{[]int{1, 2, 3}, []int{2}, 1},
		// Explanation: The only heater was placed in the position 2, and if we use the radius 1 standard, then all the houses can be warmed.
		{[]int{1, 2, 3, 4}, []int{1, 4}, 1},
		// Explanation: The two heater was placed in the position 1 and 4. We need to use radius 1 standard, then all the houses can be warmed.
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findRadius(ts.houses, ts.heaters))
		})
		t.Run(fmt.Sprintf("Example %d Adv", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, findRadiusAdv(ts.houses, ts.heaters))
		})
	}
}
