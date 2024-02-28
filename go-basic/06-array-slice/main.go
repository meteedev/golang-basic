package main

import "fmt"

func main() {
	display_array()
	display_slice()
}

func display_array() {
	var langs = [4]string{"java", "go", "python", "js"}

	fmt.Println(langs[0])
	fmt.Println(langs[len(langs)-1])
	fmt.Println(langs[0:2])
	fmt.Printf("array desc : %#v\n", langs)

}

func display_slice() {
	var dbs = []string{"oracle", "mongo", "sql", "postgres"}

	fmt.Println(dbs[0])
	fmt.Println(dbs[len(dbs)-1])
	fmt.Println(dbs[0:2])
	fmt.Printf("array desc : %v\n", dbs)

	fmt.Println(dbs)
	dbs = append(dbs, "h2")
	fmt.Println(dbs)

}
