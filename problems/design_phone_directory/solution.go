type PhoneDirectory struct {
    M map[int]int
}


/** Initialize your data structure here
        @param maxNumbers - The maximum numbers that can be stored in the phone directory. */
func Constructor(maxNumbers int) PhoneDirectory {
    p := PhoneDirectory{}
    p.M = make(map[int]int)
    for i:=0;i<maxNumbers;i++{
        p.M[i]=i
    }
    return p
}


/** Provide a number which is not assigned to anyone.
        @return - Return an available number. Return -1 if none is available. */
func (this *PhoneDirectory) Get() int {
    if len(this.M) == 0{
        return -1
    }
    item := 0
    for K,_ := range this.M{
        item = K
        break
    }
    delete(this.M,item)
    return item
}


/** Check if a number is available or not. */
func (this *PhoneDirectory) Check(number int) bool {
    if _,ok :=this.M[number];ok{
        return true
    }
    return false
}


/** Recycle or release a number. */
func (this *PhoneDirectory) Release(number int)  {
    this.M[number] = number
}


/**
 * Your PhoneDirectory object will be instantiated and called as such:
 * obj := Constructor(maxNumbers);
 * param_1 := obj.Get();
 * param_2 := obj.Check(number);
 * obj.Release(number);
 */