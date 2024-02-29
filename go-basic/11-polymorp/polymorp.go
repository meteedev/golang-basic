package main

import "fmt"

type Employee interface {
	Calculation()
}

type Employee_monthly struct {
	Name string
	Salary float32
}

type Employee_daily struct {
	Name string
	Salary float32
}

func (em Employee_monthly) Calculation() {
	fmt.Println("Call Employee_monthly method")
}

func (ed Employee_daily) Calculation() {
	fmt.Println("Call Employee_daily method")
}


func main() {
	em  := Employee_monthly{
		Name:"Mr. monthly",
		Salary: 12000.00,
	}


	ed  := Employee_daily{
		Name:"Mr. daily",
		Salary: 400.00,
	}

	employees :=[]Employee{em,ed} 

	for _ , em := range employees{
		em.Calculation()
	}

}