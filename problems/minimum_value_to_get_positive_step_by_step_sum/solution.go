func minStartValue(nums []int) int {
    sum := 0
    i:= 1
    for {
        sum = i
        wasBelowZero := false
        for i:=0;i<len(nums);i++{
            sum += nums[i]
            if sum < 1{
                break
                wasBelowZero = true
            }
        }
        if !wasBelowZero && sum > 0{
            return i
        }
        i++
    }
    return 0
}