func totalMoney(n int) int {
    sum,counter,start:=0,n/7,4
    
    for i:=1;i<=counter;i++{
       sum,start = sum+7*start, start +1
    }
    
    for i :=1;i<=n%7;i++{
       sum+=counter+i 
    }
    return sum
}