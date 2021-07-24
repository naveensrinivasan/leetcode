func maxSubArray(nums []int) int {
    n := len(nums)
    localmax := 0
    globalmax := -1000000
    
    for i:=0;i<n;i++{
        localmax = int(math.Max(float64(nums[i]),float64(nums[i]+localmax)))
        if localmax > globalmax{
            globalmax = localmax
        }
    }
    return globalmax
}