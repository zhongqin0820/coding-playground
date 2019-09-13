package problem

type intMinHeap []int

func (hs intMinHeap) Len() int { return len(hs) }

func (hs intMinHeap) Swap(i, j int) { hs[i], hs[j] = hs[j], hs[i] }

func (hs intMinHeap) Less(i, j int) bool { return hs[i] < hs[j] }

func (hs *intMinHeap) Push(e interface{}) { *hs = append(*hs, e.(int)) }

func (hs *intMinHeap) Pop() interface{} {
	e := (*hs)[hs.Len()-1]
	*hs = (*hs)[:hs.Len()-1]
	return e
}
