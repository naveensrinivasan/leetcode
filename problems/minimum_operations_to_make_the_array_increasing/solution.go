func minOperations(nums []int) int {
    if len(nums) ==1{
        return 0
    }
    result := 0
    for i:=0;i<len(nums)-1;i++{
        if nums[i] == nums[i+1]{
            result +=1
             nums[i+1] =  nums[i+1] +1
        } else if nums[i] > nums[i+1]{
            result += (nums[i] - nums[i+1]) + 1
            nums[i+1] +=  (nums[i] - nums[i+1]) + 1
        }
    }
    return result
}