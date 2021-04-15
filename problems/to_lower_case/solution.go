func toLowerCase(str string) string {
    result := []rune{}
    for _,i := range str{
        if i > 64 && i < 91{
            result = append(result,  rune(i+32))
        }else{
            result = append(result,  rune(i))
        }
    }
    return string(result)
}