func fizzBuzz(n int) []string {
    result :=[]string{}
    for i:=1;i<=n;i++{
        if i%5==0 &&  i%3==0{
            result = append(result,"FizzBuzz")
        }else if i%3==0{
            result = append(result,"Fizz")
        }else if i%5==0{
            result = append(result,"Buzz")
        }else{
            result = append(result,strconv.Itoa(i))
        }
    }
    return result
}