func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
    type F struct{
        ID, Rating int
    }
    l := []F{}
   
    
    for _, item := range restaurants{
        if veganFriendly == 1 && item[2] != veganFriendly{
            continue
        }
        if item[3] > maxPrice{
            continue
        }
        if item[4] > maxDistance{
            continue
        }
        l = append(l,F{item[0],item[1]})
    }
    
    
    sort.Slice(l,func(i,j int)bool{
        if l[i].Rating == l[j].Rating{
            return l[i].ID > l[j].ID
        }
        return l[i].Rating > l[j].Rating
    })
    
    result :=[]int{}
    for _, i:= range l{
        result = append(result,i.ID)
    }
    
    return result
}