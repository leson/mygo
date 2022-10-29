package composite

import "fmt"

func TestFunc() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	fmt.Println()

	testAnonymousFunc()
	testSelfInvocation()
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func testAnonymousFunc() {
	max := func(x int, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	r := max(44, 91)
	fmt.Printf("r: %v\n", r)
}

func testSelfInvocation() {
	r := func(x int, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}(66, 55)
	fmt.Printf("r: %v\n", r)
}
