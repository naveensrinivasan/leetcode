func findMaxAverage(nums []int, k int) float64 {
    sum := 0
    for i:=0;i<k;i++{
         sum+= nums[i]   
    }
    max := sum
    for p:=k;p <len(nums);p++{
        sum = sum - nums[p-k] + nums[p]
        if  sum > max{
            max =  sum
        }
    }
    return float64(max) / float64(k)
}