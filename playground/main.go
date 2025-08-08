package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	f := Employee{"Taro", "Yamada", 1}
	s := Employee{
		firstName: "Hanako",
		lastName:  "Nagasawa",
		id:        2,
	}
	t := Employee{}
	t.firstName = "Saburo"
	t.lastName = "Suzuki"
	t.id = 3

	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(t)
}
