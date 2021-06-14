type MovingAverage struct {
    items []int
    size int
}


/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
    m := MovingAverage{}
    m.size = size
    m.items = []int{}
    return m
}


func (this *MovingAverage) Next(val int) float64 {
    this.items = append(this.items,val)
    d := 1.0
    if len(this.items) < this.size{
        d = float64(len(this.items)) 
    }else{
        d = float64(this.size)
    }
    temp := this.items[len(this.items)-int(d):]
    sum := 0.0
    for i:=0;i<len(temp);i++{
        sum += float64(temp[i])
    }
    return sum / d
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */