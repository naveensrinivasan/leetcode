type CustomStack struct {
	L        []int
	M        int
	Position int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{L: []int{}, M: maxSize}
}

func (this *CustomStack) Push(x int) {
	if this.Position == this.M {
		return
	}
	this.L = append(this.L, x)
	this.Position++
}

func (this *CustomStack) Pop() int {
	if len(this.L) == 0 {
		return -1
	}
    x := this.L[len(this.L)-1]
    this.L = this.L[:len(this.L)-1]
	this.Position--
	return x
}

func (this *CustomStack) Increment(k int, val int) {
	if k > this.Position {
		for i := range this.L {
			this.L[i] += val
		}
        return
	}
	for i:= 0;i<k; i++ {
		this.L[i] += val
	}
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */