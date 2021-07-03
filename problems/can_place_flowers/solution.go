func canPlaceFlowers(flowerbed []int, n int) bool {
    sum := 0
    prev3 := false
    for i:=0;i<=len(flowerbed)-3;i++{
        if !prev3 && flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i+2] == 0{
           sum++        
           prev3 = true
           i+=1
        }else if flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i+2] == 0{
           sum++        
           prev3 = true
           i+=1
        }else{
            prev3 = false
        }
    }
    if len(flowerbed) > 2{
        if flowerbed[0] == 0 && flowerbed[1] == 0 && flowerbed[2] == 1{
           sum++        
        }
        if flowerbed[len(flowerbed)-3] == 1 && flowerbed[len(flowerbed)-2] == 0 && flowerbed[len(flowerbed)-1] == 0{
           sum++        
        }
        
    }
    if len(flowerbed) == 5{
        if flowerbed[0] == 0 && flowerbed[1] == 0 && flowerbed[2] == 0 && flowerbed[3] == 0 && flowerbed[4] != 0{
           sum++        
        }
        if flowerbed[len(flowerbed)-4] == 0 && flowerbed[len(flowerbed)-3] == 0 && flowerbed[len(flowerbed)-2] == 0 && flowerbed[len(flowerbed)-1] == 0{
           sum++        
        }
        
    }
    if len(flowerbed) == 2 && flowerbed[0] == 0 && flowerbed[1] == 0{
        sum++
    }
    if len(flowerbed) == 3 && flowerbed[0] == 0 && flowerbed[1] == 0 && flowerbed[2] == 0{
        sum++
    }
    if len(flowerbed) == 4 && flowerbed[0] == 0 && flowerbed[1] == 0 && flowerbed[2] == 0 && flowerbed[3] == 0{
        sum++
    }
    if len(flowerbed) == 1 && flowerbed[0] == 0{
        sum++
    }
    fmt.Println(sum)
    return sum >= n
}