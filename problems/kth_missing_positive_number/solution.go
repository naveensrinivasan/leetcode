func findKthPositive(arr []int, k int) int {
    m := make(map[int]int)
    for i:=0;i<len(arr);i++{
        m[arr[i]]= arr[i]
    }
    for i:=1;i<arr[len(arr)-1];i++{
        if _,ok:=m[i];!ok{
            k--
            if k == 0{
                return i
            }
        }
    }
    num := arr[len(arr)-1]
    for k >0{
        num++
        k--
    }
    return num
}