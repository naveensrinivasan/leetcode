type MinStack struct {
    data []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    this.data = append(this.data,val)
    
}


func (this *MinStack) Pop()  {
    if len(this.data) == 0{
        return
    }
    
    this.data = this.data[:len(this.data)-1]
    
}


func (this *MinStack) Top() int {
    return this.data[len(this.data)-1]
}


func (this *MinStack) GetMin() int {
    
    min := this.data[0]
    for i:=0;i<len(this.data);i++{
        if this.data[i]<min{
            min = this.data[i]
        }
    }
    return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */