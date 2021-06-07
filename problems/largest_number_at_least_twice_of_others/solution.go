func dominantIndex(nums []int) int {
    max,maxpos := 0,0
    for i:=0;i<len(nums);i++{
        if nums[i] > max {
            max = nums[i]
            maxpos = i
        }
    }
    
    sort.Ints(nums)
    
    for i:=0;i<len(nums);i++{
        m := 0
        if i != len(nums)-1{
            m = nums[i] *2
        }else if nums[i] <2{
            m = nums[i] *nums[i]
        }
        if m > max{
            return -1
        }
    }
    return maxpos
}