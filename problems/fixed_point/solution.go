func fixedPoint(arr []int) int {
    for i:=0;i<len(arr);i++{
        if arr[i] == i{
            return i
        }
    }
    return -1
}