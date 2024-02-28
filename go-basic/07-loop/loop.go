package main

import "fmt"

var LANGS = []string{"java", "go", "js", "react"}

func main() {
	// display_normal_loop()
	// display_range_index_value_loop()
	// display_range_value_loop()
	// display_range_index_loop()
	display_for_loop()
	display_while_loop()
}

func display_normal_loop() {
	fmt.Println("\nfor basic")
	for i := 0; i < len(LANGS); i++ {
		value := LANGS[i]
		fmt.Println(i, ":", value)
	}
}

func display_range_index_value_loop() {
	fmt.Println("\nfor range index value")
	for index, value := range LANGS {
		fmt.Println(index, ":", value)
	}
}

func display_range_value_loop() {
	fmt.Println("\nfor range  value")
	for _, value := range LANGS {
		fmt.Println(value)
	}
}

func display_range_index_loop() {
	fmt.Println("\nfor range index ")
	for index := range LANGS {
		fmt.Println(index)
	}

}

func display_for_loop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func display_while_loop() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
