package problem

import (
	"log"
	"sort"
)

type peopleSlice [][]int

func (p peopleSlice) Len() int { return len(p) }

func (p peopleSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// ascend height, descend k num
func (p peopleSlice) Less(i, j int) bool {
	if p[i][0] == p[j][0] {
		return p[i][1] < p[j][1]
	}
	return p[i][0] > p[j][0]
}

func ReconstructQueue(people [][]int) [][]int {
	return reconstructQueue(people)
}

func reconstructQueue(people [][]int) [][]int {
	sort.Stable(peopleSlice(people))
	lenP := len(people)
	if lenP == 0 {
		return people
	}
	log.Println(people)
	for i := 0; i < lenP; i++ {
		p := people[i]
		if p[1] != i {
			people = append(people[:i], people[i+1:]...) // delete index i
			// insert at index p[1]
			rear := append([][]int{}, people[p[1]:]...)
			people = append(people[:p[1]], p)
			people = append(people, rear...)
		}
	}
	return people
}
