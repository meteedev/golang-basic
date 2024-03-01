package employee

import "fmt"

type Employee struct {
	Name   string
	Salary float32
}

func Raise_salary(emp *Employee, amount float32) {
	emp.Salary += amount
}

func (emp Employee) Info() {
	fmt.Println(emp.Name)
	fmt.Println(emp.Salary)
}