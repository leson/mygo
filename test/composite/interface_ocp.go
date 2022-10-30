package composite

import "fmt"

type Pet interface {
	eat()
	sleep()
}

type Dog struct {
	name string
	age  int
}

func (d Dog) eat() {
	fmt.Printf("dog name : %v; age : %v is eating\n", d.name, d.age)
}
func (d Dog) sleep() {
	fmt.Printf("dog name : %v; age : %v is sleeping\n", d.name, d.age)
}

type Cat struct {
	name string
	age  int
}

func (c Cat) eat() {
	fmt.Printf("cat name : %v; age : %v is eating\n", c.name, c.age)
}
func (c Cat) sleep() {
	fmt.Printf("cat name : %v; age : %v is sleeping\n", c.name, c.age)
}

type Person struct {
	name string
}

func (ps Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func TestInterfaceOpenClosePrinciple() {
	d := Dog{"旺财", 2}
	c := Cat{"花花", 1}
	per := Person{"小明"}
	per.care(d)
	per.care(c)
}
