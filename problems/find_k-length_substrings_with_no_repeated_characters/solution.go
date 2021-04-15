func numKLenSubstrNoRepeats(S string, K int) int {
    if K > len(S){
        return 0
    }
    count := 0
    
    for i:=0;i<len(S)-K+1;i++{
        m := make(map[rune]int)
        hasDupes := false  
        for _ ,x := range S[i:i+K]{
            m[x]++
        } 
        for k,_ := range m{
            if m[k] > 1{
                hasDupes = true
                break
            } 
        }
        if !hasDupes{
            count++
        }
    } 
    return count
}