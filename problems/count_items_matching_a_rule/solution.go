func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    item := 0
    counter := 0
    if ruleKey == "type"{
        item = 0
    }else if ruleKey == "color"{
        item = 1
    }else{
        item = 2
    }
    
    for i := range items{
        if items[i][item] == ruleValue{
            counter++
        }
    }
    return counter
}