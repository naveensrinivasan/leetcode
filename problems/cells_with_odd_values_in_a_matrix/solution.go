func oddCells(m int, n int, indices [][]int) int {
    x := make([][]int,m)
    result := 0
    for i := range x{
        x[i]= make([]int,n)
    }   
    fmt.Println(x)
    for i:= 0;i<len(indices);i++{
        row := indices[i][0]
        col := indices[i][1]
        
        for i:=0;i<len(x[0]);i++{
            x[row][i]++
        }
        for i:=0;i<len(x);i++{
              x[i][col]++
        }
      
       // fmt.Println(i,x)
//         for i:=0;i<len(x[0]);i++{
//             fmt.Println(len(x),i,col)
//             if len(x) <=i{
                
//             }
            
//         }
        
        
        
    }
    fmt.Println(x)
   
    for i:=0;i<len(x);i++{
        for j:=0;j<len(x[0]);j++{
            if x[i][j] %2 ==1{
                result++
            }
        }
    }
    return result
}
