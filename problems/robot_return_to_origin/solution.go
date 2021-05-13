func judgeCircle(moves string) bool {
    x,y := 0,0
    for _, i := range moves{
        if i == 'L'{
            x--
        }else if i == 'R'{
            x++
        }else if i == 'U'{
            y++
        }else{
            y--
        }
    }
    return x == 0 && y == 0
}