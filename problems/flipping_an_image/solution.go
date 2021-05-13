func flipAndInvertImage(image [][]int) [][]int {
    for _,j := range image{
        l,r := 0,len(j)-1   
        for l<r{
            j[l],j[r] = j[r],j[l]
            if j[l] == 0{
                j[l] = 1
            }else{
                j[l]=0
            }
            if j[r] == 0{
                j[r] = 1
            }else{
                j[r]=0
            }
             l++
             r--
        }
      
        if len(j)%2 == 1{
              mid := (len(j)/2) 
            if j[mid] == 0{
                j[mid] = 1
            }else{
                j[mid] = 0
            }
        }
    }
    return image
}