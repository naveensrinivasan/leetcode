func buildArray(target []int, n int) []string {
    result := []string{} 
    max := target[len(target)-1]+1
    for i:=1;i<max;i++{
        if target[0] == i{
            result ,target= append(result,"Push"),target[1:]
            continue
        }
        result = append(result,[]string{"Push","Pop"}...)
    }
    return result
}