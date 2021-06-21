func check(nums []int) bool {
    lesser := 0
    for i:=0;i<len(nums)-1;i++{
        if nums[i] > nums[i+1]{
            lesser++
            if lesser > 1{
                return false
            }
        }
    }
    if nums[len(nums)-1] > nums[0]{
            lesser++
            if lesser > 1{
                return false
            }
    }
    return true
}