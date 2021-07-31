type NumArray struct {
    X []int
}


func Constructor(nums []int) NumArray {
    sum := 0
    x := make([]int,len(nums))
    for i, item := range nums{
        sum += item
        x[i] = sum
    }
    return NumArray{x}
}


func (this *NumArray) SumRange(left int, right int) int {
    if left == 0{
        return this.X[right]
    }
    
    return this.X[right] - this.X[left-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */