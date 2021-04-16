func numKLenSubstrNoRepeats(S string, K int) int {
    if K > len(S){
        return 0
    }
    
    count := 0
    dupes := 0
    
    
    m := make(map[string]int)
    first := string(S[0])
    
    for i:=0;i<K;i++{
        m[string(S[i])]++
        if m[string(S[i])] == 2 {
            dupes++
        }
    }
    if dupes == 0{
        count++
    }
    
    
    for i:=K;i<len(S);i++{      
        m[first]--
        if (m[first] ==1){
                dupes --
        }
        
        first= string(S[i-K+1])
        
        m[string(S[i])]++
        if m[string(S[i])] == 2 {
            dupes++
        }
        
        if dupes== 0{
                count++
        }
    } 
    return count
}