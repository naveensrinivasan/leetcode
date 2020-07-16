

func twoSum(nums []int, target int) []int {
    left:= 0
    right:= len(nums) -1
    for left<right{
        total := nums[left] + nums[right]
        if total == target{
            return []int{left+1,right+1}
        }else if total > target{
            right--
        }else{
            left++
        }
    }
    return nil
}