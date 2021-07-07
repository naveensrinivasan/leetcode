func isMonotonic(nums []int) bool {
    isIncreasing := false
    isSet := false
    for i:=0;i<len(nums)-1;i++{
        if nums[i] == nums[i+1]{
            continue
        }
        if nums[i+1] - nums[i] > 0 {
            if !isSet{
                isSet = true
                isIncreasing = true
                continue
            }else{
                if !isIncreasing{
                    return false
                }
            }
        }else if nums[i] - nums[i+1] > 0{
            if !isSet{
                isSet = true
                isIncreasing = false
                continue
            }else{
                if isIncreasing{
                    return false
                }
            }
        }
    }
    return true
}