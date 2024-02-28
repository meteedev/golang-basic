package main

import "fmt"

func main() {

	errors_map := map[int]string{}

	errors_map[200] = "Ok"
	errors_map[400] = "Bad Request"
	errors_map[500] = "Internal Server Error"

	fmt.Printf("% #v\n", errors_map)

}
