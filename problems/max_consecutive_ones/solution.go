func findMaxConsecutiveOnes(nums []int) int {
    result,max := 0,0
    for i:=0;i<len(nums);i++{
        if nums[i]==0{
            result = 0
        }else{

            result++
            if result > max{
                max = result
            }
            
        }
    }
    return max
}