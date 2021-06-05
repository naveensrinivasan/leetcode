type ProductOfNumbers struct {
    L []int
}


func Constructor() ProductOfNumbers {
    return ProductOfNumbers{}
}


func (this *ProductOfNumbers) Add(num int)  {
    this.L = append(this.L,num)
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    x := len(this.L) -k
    sum := 1
    for i:=len(this.L)-1;i>=x;i--{
        sum*= this.L[i]
    }
    return sum
}


/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */