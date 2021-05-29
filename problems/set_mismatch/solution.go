func findErrorNums(nums []int) []int {
    m := make(map[int]int)
    missing,dupe := 0,0
    for _,n := range nums{
        m[n]++
    }
    for i:=1;i<=len(nums);i++{
        if v,ok := m[i];!ok{
            missing = i
        }else if v > 1{
            dupe = i
        }
    }
    return []int{dupe,missing}
}