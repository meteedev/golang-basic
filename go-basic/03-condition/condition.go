package main

import "fmt"

func main() {
	if_condition()
	if_switch()
}

func if_condition() {

	score := 65
	grade_str := "Your grade is "

	if score < 50 {
		fmt.Println(grade_str, "F")
	} else if score <= 50 {
		fmt.Println(grade_str, "D")
	} else if score <= 60 {
		fmt.Println(grade_str, "C")
	} else if score <= 70 {
		fmt.Println(grade_str, "B")
	} else if score >= 80 {
		fmt.Println(grade_str, "A")
	}

}

func if_switch() {
	num := 1

	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Default Case")
	}
}
