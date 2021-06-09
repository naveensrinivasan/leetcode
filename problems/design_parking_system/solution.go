type ParkingSystem struct {
    b int
    m int
    s int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{b:big,m:medium,s:small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
    switch carType {
        case 1:
        this.b--
        if this.b <0{
            return false
        }
    case 2:
        this.m--
        if this.m < 0{
            return false
        }
    case 3:
        this.s--
        if this.s<0{
            return false
        }
    }
    return true
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */