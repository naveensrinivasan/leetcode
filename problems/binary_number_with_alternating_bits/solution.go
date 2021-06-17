func hasAlternatingBits(n int) bool {
    prev := -1
    for n > 0{
        if n &1 == prev{
            return false
        }
        prev,n = n & 1, n >> 1
    }
    return true
}