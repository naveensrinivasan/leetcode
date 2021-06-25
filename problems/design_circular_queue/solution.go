type MyCircularQueue struct {
    L []int
    P int
    M int
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{L:[]int{},P:0,M:k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.M == this.P{
        return false
    }
    this.L = append(this.L,value)
    this.P++
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if len(this.L) == 0{
        return false
    }
    this.L = this.L[1:]
    this.P--
    return true
}


func (this *MyCircularQueue) Front() int {
    if len(this.L) == 0{
        return -1
    }
     x := this.L[0]
    return x
}



func (this *MyCircularQueue) Rear() int {
    if len(this.L) == 0{
        return -1
    }
    x := this.L[len(this.L)-1]
    return x
}


func (this *MyCircularQueue) IsEmpty() bool {
    return len(this.L) == 0
}

func (this *MyCircularQueue) IsFull() bool {
    return this.M == this.P 
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */