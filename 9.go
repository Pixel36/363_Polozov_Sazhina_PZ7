package main

import "fmt"

type Employee struct {
 Name     string
 Position string
 Salary   float64
}

type Department struct {
 Employees []Employee
}

func (d *Department) AddEmployee(e Employee) {
 d.Employees = append(d.Employees, e)
}

func (d *Department) RemoveEmployee(name string) {
 newList := []Employee{}
 for _, e := range d.Employees {
  if e.Name != name {
   newList = append(newList, e)
  }
 }
 d.Employees = newList
}

func (d *Department) TotalSalary() float64 {
 total := 0.0
 for _, e := range d.Employees {
  total += e.Salary
 }
 return total
}

func (d *Department) GetByPosition(pos string) []Employee {
 var res []Employee
 for _, e := range d.Employees {
  if e.Position == pos {
   res = append(res, e)
  }
 }
 return res
}

func main() {
 dept := Department{}
 dept.AddEmployee(Employee{"Карина", "Программист", 100000})
 dept.AddEmployee(Employee{"Матвей", "Тестировщик", 70000})
 fmt.Println("Фонд зарплат:", dept.TotalSalary())
}
