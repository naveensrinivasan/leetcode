func missingElement(nums []int, k int) int {
    missingCounter := 0
    m := make(map[int]int)
    for i:=0;i<len(nums);i++{
        m[nums[i]] = nums[i]
    }
    for i:= nums[0];i<=nums[len(nums)-1];i++{
        if _,ok:=m[i];!ok{
            missingCounter++
            if missingCounter == k{
                return i
            }
        }
    }
    fmt.Println(nums[len(nums)-1],(k-missingCounter))
    return nums[len(nums)-1]+(k-missingCounter)
}