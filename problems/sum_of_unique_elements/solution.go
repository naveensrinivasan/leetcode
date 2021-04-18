func sumOfUnique(nums []int) int {
    m := make(map[int]int)
    sum := 0
    for i:= 0;i<len(nums);i++{
        if v,ok:= m[nums[i]];!ok{
            sum+=nums[i]
        }else{
            if v ==1{
                sum -= nums[i]
            }
        } 
        m[nums[i]]++
    }
    return sum
}