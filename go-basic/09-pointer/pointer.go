package main

import (
	"fmt"
	"github.com/meteedev/learnapp/pointer/employee"	
)

func main() {
	teleport()
	raise_salary()
}

// func warp( now *string,to string)string{
// 	now := to
// }

func raise_salary(){
	 emp :=  employee.Employee {
			Name:"lieang",
			Salary:1200.0,
	 }

	 emp.Info()

	 employee.Raise_salary(&emp,100.0)

	 emp.Info()


}

func teleport(){
	variable_now := "Bangkok"
	fmt.Println("I'm in ",variable_now)
	display_addr(variable_now)
	display_value(&variable_now)
	//display_value(&variable_now)
	addr := &variable_now;
	warp_location(addr,"Chiang Mai")
	fmt.Println("I' in ",variable_now)

}


func display_addr(variable string){
	fmt.Println(&variable)
}

func display_value(address *string){
	value :=  *address
	fmt.Println(value)
}


func warp_location(from  *string, destination string){
	*from = destination
}
