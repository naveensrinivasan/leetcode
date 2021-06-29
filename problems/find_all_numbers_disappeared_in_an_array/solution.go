func findDisappearedNumbers(nums []int) []int {
    result := []int{}
    m := make(map[int]int)
    for _,i := range nums{
        m[i]=i
    }
    for i:=1;i<=len(nums);i++{
        if _,ok:=m[i];!ok{
            result = append(result,i)
        }
    }
    return result
}