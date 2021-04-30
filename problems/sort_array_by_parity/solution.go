func sortArrayByParity(A []int) []int {
    result := []int{}
    for i:=0;i<len(A);i++{
        if A[i] %2 != 0{
            result = append(result,A[i])
        }else{
            result = append([]int{A[i]},result...)
        }
    }
    return result
}