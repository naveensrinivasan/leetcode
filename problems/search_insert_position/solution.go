func searchInsert(nums []int, target int) int {
    left,right,ans := 1,len(nums),0
    
    if target > nums[len(nums)-1]{
        return len(nums)
    }
    
    if target <= nums[0]{
        return 0
    }
    for left<= right{
        mid := (left+right)/2
        
        if nums[mid] == target{
            return mid
        }
        if nums[mid] < target{
            left = mid +1
            ans = left 
        }else{
            right = mid -1
            ans = right + 1 
        }
    }
    return ans
  
}