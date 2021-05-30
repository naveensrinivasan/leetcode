func maxNumberOfBalloons(text string) int {
    m := make(map[string]int)
    for i:=0;i<len(text);i++{
        switch text[i]{
            case 'b':
            fallthrough
            case 'a':
            fallthrough
            case 'l':
            fallthrough
            case 'o':
            fallthrough
            case 'n':
            m[string(text[i])]++
        }
    } 
    result := 0
    for {
        word := "balloon"
        for _ , i := range word{
            if _,ok:= m[string(i)];!ok{
                return result
            }else{
                m[string(i)]--
                if m[string(i)] == 0{
                    delete(m,string(i))
                }
            }
            
        }
        result++
    }
    
    return 0
}