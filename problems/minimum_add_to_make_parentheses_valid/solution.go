func minAddToMakeValid(s string) int {
    stack :=[]string{}
    for _, i := range strings.Split(s,""){
        if i == ")" && len(stack) > 0 && stack[len(stack)-1] == "(" {
            stack = stack[:len(stack)-1]
            continue
        }
        stack = append(stack,string(i))
    }
    return len(stack)
}