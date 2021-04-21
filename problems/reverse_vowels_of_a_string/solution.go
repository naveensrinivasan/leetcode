func reverseVowels(s string) string {
    l,r :=0,len(s)-1
    result := []rune(s)
    
    for l<r {
        foundLeft,foundRight := false,false
        for l<=r  {
            
            if isVowel(rune(result[l])){
                foundLeft = true
                break
            }
            l++
        }
        
        for r>=0{
            if isVowel(rune(result[r])){
                foundRight = true
                break
            }
            r--
        }
        // if l == r || r==0{
        //     return string(result)
        // }
        if foundLeft && foundRight {
            result[l],result[r] = result[r],result[l]    
        }
        
        l++
        r--
    }
    return string(result)
}

func isVowel(r rune) bool{
    if r=='a' || r == 'e' || r == 'i' || r== 'o' || r == 'u' || r=='A' || r == 'E' || r == 'I' || r== 'O' || r == 'U'{
        return true
    }
    return false
}