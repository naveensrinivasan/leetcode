func distributeCandies(candies int, num_people int) []int {
    counter := 1
    result := make([]int,num_people)
    for candies > 0{
        for i:=0;i<num_people;i++{
            if candies  > counter{
                candies -= counter
            }else{
                counter = candies
                candies =0
            }
            result[i]+=counter
            if candies == 0{
                break
            }
            counter++
        }
    }
    return result 
}