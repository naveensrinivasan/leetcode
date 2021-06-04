func maximumTime(time string) string {
    data := strings.Split(time,"")
    if data[0] == "?" && data[1] == "?"{
        data[0] = "2"
        data[1] = "3"
    }
    
    if data[1] == "?" {
         if data[0] == "2"{
            data[1] = "3"
        }else{

            data[1] = "9"
        }
    }
    
    if data[0] == "?" {
        x , _ := strconv.Atoi(data[1])
        if x <= 3{
            data[0] = "2" 
        }else{
            data[0] = "1"
        }
    }
    
    if data[3] == "?" && data[4] == "?"{
        data[3] = "5"
        data[4] = "9"
    }
    
    if data[4] == "?"{
        data[4] = "9"
    }
    if data[3] == "?"{
        data[3] = "5"
    }

    return strings.Join(data,"")
}