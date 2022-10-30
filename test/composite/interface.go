package composite

import "fmt"

type Animal interface {
	eat()
}

type Pig struct {
	name string
}

type Sheep struct {
	name string
}

func (p Pig) eat() {
	fmt.Printf("pig's name: %v is eating\n", p.name)
}

func (s Sheep) eat() {
	fmt.Printf("sheep's name: %v is eating\n", s.name)
}

func TestInterface() {
	/** testing polymorphism **/
	var an Animal
	an = Pig{"哼哼"}
	an.eat()
	an = Sheep{"咩咩"}
	an.eat()

}
