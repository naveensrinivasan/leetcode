func stringShift(s string, shift [][]int) string {
    for _, item := range shift{
        for i := 0;i<item[1];i++{
            if item[0] == 0{
                s = s[1:] + string(s[0])
            }else{
                s =  string(s[len(s)-1]) + s[:len(s)-1] 
            }
        }
    }
    return s
}