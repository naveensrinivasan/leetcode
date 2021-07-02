func minStartValue(nums []int) int {
    sum := 1
    min := 1
    for i:=0;i<len(nums);i++{
        sum += nums[i]   
        if sum < min{
            min = sum
        }
    }
    if min < 1{
        return 2-min
    }
    return 1
}