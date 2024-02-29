package main

import "fmt"

func main() {
	variable_now := "Bangkok"
	
	fmt.Println("I'm in ",variable_now)
	display_addr(variable_now)
	display_value(&variable_now)
	//display_value(&variable_now)
	addr := &variable_now;
	warp(addr,"Chiang Mai")
	fmt.Println("I' in ",variable_now)
}

// func warp( now *string,to string)string{
// 	now := to
// }

func display_addr(variable string){
	fmt.Println(&variable)
}

func display_value(address *string){
	value :=  *address
	fmt.Println(value)
}


func warp(from  *string, destination string){
	*from = destination
}
