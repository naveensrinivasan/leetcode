func minStoneSum(piles []int, k int) int {
    h := &IntHeap{}
    for i:=0;i<len(piles);i++{
        *h = append(*h,piles[i])
    }
	heap.Init(h)
    for k > 0{
        x :=  heap.Pop(h).(int)
        if x %2 == 1{
            heap.Push(h, (x/2)+1)
        }else{
            heap.Push(h,x/2)
        }
        k--
    }
    sum := 0
    for h.Len() > 0{
        x := heap.Pop(h).(int)
        sum += x
    }
    return sum
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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
