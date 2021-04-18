func repeatedNTimes(A []int) int {
    sort.Ints(A)
    if len(A) == 1{
        return A[0]
    }
    for i:=1;i<len(A);i++{
        if A[i-1] == A[i]{
            return A[i]
        }
    }
    return A[0]
}