type TwoSum struct {
    L []int
}


/** Initialize your data structure here. */
func Constructor() TwoSum {
    return TwoSum{[]int{}}
}


/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int)  {
    this.L = append(this.L,number)
}


/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
    sort.Ints(this.L)
    l,r := 0,len(this.L) -1
    for l < r{
        sum := this.L[l] + this.L[r]
        if sum == value{
            return true
        }else if sum < value {
            l++
        }else{
            r--
        }
    }
    return false
}


/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */