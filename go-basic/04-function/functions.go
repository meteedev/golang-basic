package main

import "fmt"

func main() {
	fmt.Println(sayHello("lieang"))
	fmt.Println(cal(add, 1, 2))
	fmt.Println(cal(subtract, 1, 2))
}

func sayHello(name string) string {
	return "hello " + name
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func cal(operator func(int, int) int, a int, b int) int {
	result := operator(a, b)
	fmt.Println("result =", result)
	return result
}
