package composite

import "fmt"

func TestDefer() {
	fmt.Println("Start")
	defer fmt.Println("step1")
	defer fmt.Println("step2")
	defer fmt.Println("step3")
	fmt.Println("End")
}
