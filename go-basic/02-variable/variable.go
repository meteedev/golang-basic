package main

import "fmt"

var name string = "Lieang"

func main() {
	display_variable()
	display_const_normal()
	display_const_iota()
}

func display_variable() {
	string_var := "This is a book"
	int_var := 5
	float_var := 3.7
	boolean_var := true
	rune_var := '界'

	fmt.Println(name)
	fmt.Printf("name: %q\n", name)
	fmt.Printf("type: %T\n", name)

	fmt.Printf("int: %v\n", int_var)
	fmt.Printf("float64: %v\n", float_var)
	fmt.Printf("bool: %v\n", boolean_var)
	fmt.Printf("string: %v\n", string_var)
	fmt.Printf("rune: %v\n", rune_var)

	fmt.Printf("int: %v\n", int_var)
	fmt.Printf("float64: %T\n", float_var)
	fmt.Printf("bool: %T\n", boolean_var)
	fmt.Printf("string: %T\n", string_var)
	fmt.Printf("rune: %T\n", rune_var)

}

func display_const_normal() {
	const (
		sunday    = 0
		monday    = 1
		tuesday   = 2
		wednesday = 3
		thrusday  = 4
		friday    = 5
		saturday  = 6
	)

	fmt.Println("tuesday :", tuesday)

}

func display_const_iota() {
	const (
		sunday = iota
		monday
		tuesday
		wednesday
		thrusday
		friday
		saturday
	)

	fmt.Println("tuesday :", tuesday)

}
