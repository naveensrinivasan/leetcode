func validIPAddress(IP string) string {
    if ipv4(IP){
    return "IPv4"
    }else if ipv6(IP){
        return "IPv6"
    }
    return "Neither"
}

func ipv4(ip string)bool{
    cidr := strings.Split(ip,".")
    if len(cidr) != 4{
        return false
    }
    for _,c := range cidr{
        if len(c) > 1 && string(c[0]) == "0"{
            return false
        }
        n,err := strconv.Atoi(c)
        if err != nil{
            return false
        }
        if n < 0 || n > 255{
            return false
        }
    }
    return true
}
func ipv6(ip string) bool{
    cidr := strings.Split(ip,":")
    if len(cidr) != 8{
        return false
    }
     for _,c := range cidr{
         if len(c) > 4 || len(c) ==0{
             return false
         }
         for _, i := range c{
             if (rune(i) >= 65 && rune(i) <= 70) || 
                (rune(i) >=97 && rune(i) <= 102) || 
                (rune(i) >=48 &&  rune(i) <=57){
                 continue
             }
             return false
         }   
    }
    return true
}