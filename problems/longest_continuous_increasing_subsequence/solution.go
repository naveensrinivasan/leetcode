func findLengthOfLCIS(nums []int) int {
    if len(nums) ==1{
        return 1
    }
    result,max,x:= 1,1,nums[0]
    
    for i:=1;i<len(nums);i,x = i+1,nums[i]{
        if nums[i] > x{
            result++
            if result > max{
                max = result
            }
        }else{
            result =1
        }
    }
    return max
}