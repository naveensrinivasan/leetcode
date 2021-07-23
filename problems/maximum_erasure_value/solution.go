func maximumUniqueSubarray(nums []int) int {
    l,r,max,mapsize :=0,0,0,0
    m := make(map[int]int)
    for (r<len(nums)){
        if _, ok := m[nums[r]];!ok{
            m[nums[r]] = nums[r]
            mapsize += nums[r]
            r++
            max = Max(max,mapsize)
        }else{
            delete(m,nums[l])
            mapsize -= nums[l]
            l++
        }
    }
    return max
}



// Max returns the larger of x or y.
func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}