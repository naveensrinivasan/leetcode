func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
    m := make(map[int]int)
    m2 := make(map[int]int)
    result := []int{}
    for _,i := range arr1{
        m[i] = 1
    }
    
    for _, i := range arr2{
        if _,ok:= m[i];ok{
            m2[i] = 1
        }
    }
    
       
    for _, i := range arr3{
        if _,ok:= m2[i];ok{
            result = append(result,i)
        }
    }
  
    return result
}