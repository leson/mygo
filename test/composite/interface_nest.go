package composite

import "fmt"

type Swim interface {
	swim()
}
type Fly interface {
	fly()
}

type FlyFish interface {
	Swim
	Fly
}

type Fish struct {
}

func (f Fish) swim() {
	fmt.Println("Fish is swiming")
}

func (f Fish) fly() {
	fmt.Println("Fish is flying")
}

func TestInterfaceNest() {
	/** testing polymorphism **/
	// var ff FlyFish
	ff := Fish{}
	ff.fly()
	ff.swim()
}
