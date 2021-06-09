func numberOfDays(Y int, M int) int {
    odd := make(map[int]int)
    odd[1]=1
    odd[3]=3
    odd[5]=5
    odd[7]=7
    odd[8]=8
    odd[10]=10
    odd[12]=12
    
    isLeapYear := false
    if Y %4 == 0{
        isLeapYear = true
        if Y%100 == 0{
           isLeapYear = false 
        }
        if Y%400 == 0{
            isLeapYear = true
        }
    }
    
    if _,ok := odd[M];ok{
        return 31
    }
    
    if M==2 && isLeapYear {
        return 29
    }
    if M == 2{
        return 28
    }
    return 30
}