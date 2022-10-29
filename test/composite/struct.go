package composite

import "fmt"

type User struct {
	name   string
	gender string
}

func TestStruct() {
	user := User{"Leson", "Male"}
	fmt.Printf("%v\n", user)
}
