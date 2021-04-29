func findPeakElement(nums []int) int {
    for i:=1;i<len(nums);i++{
        if nums[i-1]<nums[i]{
             continue
        }else{
            return i-1
        } 
    }
    return len(nums)-1
}