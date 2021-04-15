func halvesAreAlike(s string) bool {
    a,b,mid := 0,0, len(s)/2
    s = strings.ToLower(s)
    for i:=0;i<mid;i++{
        if isVowel(rune(s[i])){
            a++
        }
        if isVowel(rune(s[mid+i])){
            b++
        }
    } 
    
    return a == b
}

func isVowel(r rune) bool{
    if r == 'a' || r == 'e' || r== 'i' || r=='o' || r == 'u'{
        return true
    }
    return false
}