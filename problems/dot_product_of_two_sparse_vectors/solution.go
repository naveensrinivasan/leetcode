type SparseVector struct {
    X []int
}

func Constructor(nums []int) SparseVector {
    return SparseVector{nums}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    sum := 0
    for i:=0;i<len(this.X);i++{
        sum += this.X[i] * vec.X[i]
    }
    return sum
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */