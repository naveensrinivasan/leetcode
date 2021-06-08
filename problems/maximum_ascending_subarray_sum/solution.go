func maxAscendingSum(nums []int) int {
    l,r,max,m := 0,1,nums[0],0
    
    for l <= r && r < len(nums){
        if nums[l] < nums[r]{
           m += nums[l]
        }else{
            if l >0 && nums[l-1] < nums[l]{
                m+= nums[l]
                if m > max {
                    max = m
                }
            }
            m = 0
        }
        l,r = r,r+1
        if m > max {
            max = m
        }
    }
    if l >0 && nums[l-1] < nums[l]{
        m+= nums[l]
    }
    if m > max {
        max = m
    }
    return max
}