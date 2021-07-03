func isArmstrong(n int) bool {
    l := int(math.Log10(float64(n)))+1
    x := n
    sum :=0 
    for x > 0{
        sum+= pow(x%10,l)
        x/=10
    }
    return sum==n
}

func pow(x,y int)int{
   return int(math.Pow(float64(x),float64(y)))
}


