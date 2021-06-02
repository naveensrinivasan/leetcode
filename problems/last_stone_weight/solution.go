func lastStoneWeight(stones []int) int {
    for len(stones) > 1{
        sort.Sort(sort.Reverse(sort.IntSlice(stones)))
        if stones[0] == stones[1]{
            stones = stones[2:]
        }else{
            result := stones[0] - stones[1]
            stones = stones[2:]
            stones = append(stones,result)
        }
    }
    
    if len(stones) == 1{
        return stones[0]
    }
    
    return 0
}