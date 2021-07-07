func findSpecialInteger(arr []int) int {
    if len(arr) == 1{
        return arr[0]
    }
    x, max,count,maxitem:= -1,0,0,-1
    x = arr[0]
    for i :=1;i<len(arr);i++{
        if arr[i] == x {
            count++
        }else{
            count = 0
        }
        x = arr[i]
        if count > max{
            max = count
            maxitem = arr[i]
        }
    }
    return maxitem
}