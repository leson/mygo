package composite

import (
	"fmt"
	"strings"
)

func TestFunc() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	fmt.Println()

	testAnonymousFunc()
	testSelfInvocation()
	testFuncCaculateInternal()
	testClosure()
	testClosure1()
	testClosure2()
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

func testClosure() {
	f := add()
	fmt.Printf("f: %v\n", f(10))
	fmt.Printf("f: %v\n", f(10))
	fmt.Println("======================")
	f1 := add()
	fmt.Printf("f1 return: %v\n", f1(100))
	fmt.Printf("f1 return: %v\n", f1(100))
}

func add() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func testClosure1() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Printf("jpgFunc: %v\n", jpgFunc("hello.jpg"))
	fmt.Printf("txtFunc: %v\n", txtFunc("world"))
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}

	}
}

func testClosure2() {
	add, sub := calc(50)
	fmt.Printf("add 10: %v\n", add(10))
	fmt.Printf("sub 20: %v\n", sub(20))
	fmt.Printf("add 30: %v\n", add(30))
	fmt.Printf("sub 100: %v\n", sub(100))
}

func calc(t int) (func(int) int, func(int) int) {
	add := func(i int) int {
		t += i
		return t
	}

	sub := func(i int) int {
		t -= i
		return t
	}

	return add, sub
}
