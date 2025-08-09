package solution

import (
	"container/heap"
)

func SumaGolosa(set []int) int {
	if len(set) < 2 {
		return  0
	}

	h := &IntHeap{}
	heap.Init(h)
	for _, num := range set {
		heap.Push(h, num)
	}

	totalCost := 0

	for h.Len() > 1 {
		first := heap.Pop(h).(int)
		second := heap.Pop(h).(int)

		cost := first + second
		totalCost += cost

		heap.Push(h, cost)
	}

	return totalCost
}


type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}