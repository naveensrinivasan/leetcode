func sortedSquares(nums []int) []int {
    
    result := []int{}
    left,right := 0,len(nums)-1
    for left<=right{
        if abs(nums[left]) > abs(nums[right]){
            result = append([]int{abs(nums[left]) * abs(nums[left])},result...)
            left++
            
        }else{
            result = append([]int{abs(nums[right]) * abs(nums[right])},result...)
            right--
            
        } 
    }
    
    return result
}

func abs(x int)int{
    if x < 0{
        return x * -1
    }
    return x
}