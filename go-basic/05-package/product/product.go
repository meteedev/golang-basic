package product

import "fmt"

func View_products() {
	fmt.Println("view product")
	find_by_sql()
}

func connect_db() {
	fmt.Println("private method connect db")
}

func find_by_sql() {
	fmt.Println("private method find_by_sql")
	connect_db()
	fmt.Println("return data")
}
