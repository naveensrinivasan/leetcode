func removeVowels(s string) string {
    for i:=0;i<len(s);i++{
        if s[i] == 'a' || s[i] == 'e' || s[i] == 'o' || s[i] == 'u' || s[i] == 'i'{
            s= s[:i]+s[i+1:]
            i--
        }
    }
    return s
}