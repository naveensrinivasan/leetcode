func runningSum(arr []int) []int {
    for i:=1;i<len(arr);i++{
        arr[i] += arr[i-1]
    }
    return arr
}