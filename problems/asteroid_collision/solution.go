func asteroidCollision(asteroids []int) []int {
    result := []int{}
    for _, item := range asteroids{
        if len(result) == 0{
            result= append(result, item)
            continue
        }
        result = foo(item,result)
    }
    return result
}

func foo(item int, result []int) []int{
     x:= result[len(result)-1]
      switch x>0{
            case true:
            if item > 0{ 
              result= append(result, item)  
            }else{ 
                if x == item * -1{
                    result = result[:len(result)-1]  
                }else if x < item * -1{
                    result = result[:len(result)-1]  
                    result= append(result, item)
                }
                if len(result) > 1{
                    return foo(result[len(result)-1],result[:len(result)-1])
                }
            } 
            case false:
                  result= append(result, item)  
        }
    return result
}