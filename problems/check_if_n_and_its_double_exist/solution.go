func checkIfExist(arr []int) bool {
    m := make(map[int]int)
    for i:=0;i<len(arr);i++{
        m[arr[i]*2]= i
        if v,ok := m[arr[i]];ok && v != i{
            return true
        }
    }
     for i:=0;i<len(arr);i++{
         if v,ok := m[arr[i]];ok && v != i{
            return true
        }
     }
    return false
}