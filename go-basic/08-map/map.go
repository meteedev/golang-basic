package main

import "fmt"

func main() {

	errors_map := map[int]string{}

	errors_map[200] = "Ok"
	errors_map[400] = "Bad Request"
	errors_map[500] = "Internal Server Error"

	display_map(errors_map)

	errors_map[404] = "not found"

	display_map(errors_map)

	delete(errors_map, 200)

	display_map(errors_map)

	loop_display(errors_map)

	is_found_code(101, errors_map)
}

func display_map(error_map map[int]string) {
	fmt.Printf("% #v\n", error_map)
}

func loop_display(error_map map[int]string) {
	for code, value := range error_map {
		//add comment
		fmt.Println(code, value)
	}
}

func is_found_code(code int, error_map map[int]string) {
	if value, ok := error_map[code]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("not found")
	}

}
