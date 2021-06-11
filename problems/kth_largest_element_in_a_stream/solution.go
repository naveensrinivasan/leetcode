type KthLargest struct {
    i int
    l *IntHeap
}


func Constructor(k int, nums []int) KthLargest {
    h := &IntHeap{}
    for _ , i := range nums{
        h.Push(i)   
    }
	heap.Init(h)
    for len(*h) > k{
        heap.Pop(h)
    }
    return KthLargest{k,h}
}


func (this *KthLargest) Add(val int) int {
  if len(*this.l) < this.i {
		heap.Push(this.l, val)
	} else if val > (*this.l)[0] {
		heap.Push(this.l, val)
		heap.Pop(this.l)
	}
    return (*this.l)[0]
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i]  <h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
