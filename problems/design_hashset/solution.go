type MyHashSet struct {
    M map[int]int
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
    x := MyHashSet{}
    x.M = make(map[int]int)
    return x
}


func (this *MyHashSet) Add(key int)  {
    if _,ok := this.M[key];!ok{
        this.M[key]=key
    }
}


func (this *MyHashSet) Remove(key int)  {
    delete(this.M,key)
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
    _,ok := this.M[key]
    return ok
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */