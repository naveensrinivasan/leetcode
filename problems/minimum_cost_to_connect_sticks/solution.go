func connectSticks(sticks []int) int {
  minHeap := &MinHeap{}
	for _, i := range sticks {
		heap.Push(minHeap, i)
	}
    
    total :=0
    
    for minHeap.Len() >1  {
        x:= heap.Pop(minHeap)
        y:= heap.Pop(minHeap)
        total += x.(int)+ y.(int)
        heap.Push(minHeap,x.(int)+y.(int))
    }
    return total
}
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() int          { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
