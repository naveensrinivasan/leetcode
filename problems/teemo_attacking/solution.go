func findPoisonedDuration(timeSeries []int, duration int) int {
    if len(timeSeries) == 1{
        return duration
    }
    sum := 0
    for i:=0;i<len(timeSeries)-1;i++{
        if timeSeries[i]+duration <= timeSeries[i+1]{
            sum+= duration
        }else{
            sum += timeSeries[i+1] - timeSeries[i]
        }
    }

    if len(timeSeries) >1 && timeSeries[len(timeSeries)-2]+duration <= timeSeries[len(timeSeries)-1]{
            sum+= duration
        }else if len(timeSeries) >1{
            if (timeSeries[len(timeSeries)-1] - timeSeries[len(timeSeries)-2]) == 1{
                sum += duration + (timeSeries[len(timeSeries)-1] - timeSeries[len(timeSeries)-2]) -1
            }else{
               sum += duration
            }
        } 

    return sum
}