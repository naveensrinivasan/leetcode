func lengthOfLastWord(s string) int {
    x := strings.Fields(s)
    if len(x) ==0 {
        return 0
    }
    return len(x[len(x)-1])
}