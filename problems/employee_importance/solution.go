/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
    m := make(map[int]*Employee)
    for _,item := range employees{
        m[item.Id] = item
    }
    x := m[id]
    sum :=  x.Importance
    f:= []int{}
    f = append(x.Subordinates,f...)
    for i:=0;i<len(f);i++{
        for _, x := range m[f[i]].Subordinates{
            f= append(f,x)
        }
        sum += m[f[i]].Importance
    }
    return sum
}