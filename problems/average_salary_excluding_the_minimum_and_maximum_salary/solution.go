func average(salary []int) float64 {
    sort.Ints(salary)
    salary = salary[1:len(salary)-1]
    x := float64(0)
    for _,item := range salary{
        x+= float64(item)
    }
    return x/float64(len(salary))
}