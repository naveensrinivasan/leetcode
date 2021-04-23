func balancedStringSplit(s string) int {
    balance,counter := 0, 0
    for _, i:= range s{
        if i == 'L'{
            balance++
        }else{
            balance --
        }
        if balance == 0{
            counter++
        }
    }
    return counter
}