func findDuplicate(nums []int) int {
    m := make(map[int]int)
    for i:=0;i<len(nums);i++{
        if _,ok := m[nums[i]];!ok{
            m[nums[i]] =1
        }else{
            return nums[i]
        }
    }
    return 0
}