
func defangIPaddr(address string) string {
    r := []rune{}
    for _,item := range address{
        if item == '.'{
            r = append(r,'[','.',']')
        }else{
            r = append(r,rune(item))
        }
    }
    return string(r)
}