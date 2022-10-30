package composite

import "fmt"

type Pet interface {
	eat()
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (d Dog) eat() {
	fmt.Printf("dog's name: %v is eating\n", d.name)
}

func (c Cat) eat() {
	fmt.Printf("cat's name: %v is eating\n", c.name)
}

func TestInterface() {
	/** testing polymorphism **/
	var p Pet
	p = Dog{"旺财"}
	p.eat()
	p = Cat{"花花"}
	p.eat()

}
