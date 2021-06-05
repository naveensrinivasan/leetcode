type MyQueue struct {
    L []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.L = append([]int{x},this.L...)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    x:= this.L[len(this.L)-1]
    this.L = this.L[:len(this.L)-1]
    return x
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    return this.L[len(this.L)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.L) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */