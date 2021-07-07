func findSpecialInteger(arr []int) int {
    if len(arr) == 1{
        return arr[0]
    }
    max,count,maxitem:= 0,0,-1
    for i :=1;i<len(arr);i++{
        if arr[i] == arr[i-1] {
            count++
            if count > max{
                max = count
                maxitem = arr[i]
            }
        }else{
            count = 0
        }
    }
    return maxitem
}