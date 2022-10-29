package composite

import "fmt"

func TestFunc() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	fmt.Println()

	testAnonymousFunc()
	testSelfInvocation()
	testFuncCaculateInternal()
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

/**
 * simulate closure
 */
func testFuncCaculateInternal() {
	name := "Alice "
	age := "18"

	s1 := func() string {
		return name + age
	}()
	fmt.Printf("s1: %v\n", s1)
}
