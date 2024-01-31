package test

import (
	"fmt"
	"testing"
)

type operation func(int, int) int

func calculate(a, b int, op operation) int {
	return op(a, b)
}

func TestFun(t *testing.T) {
	//匿名函数
	f := func(i, j int) int {
		return j + i
	}
	f(1, 1)

	// lambda函数示例
	add := func(a, b int) int {
		return a + b
	}

	subtract := func(a, b int) int {
		return a - b
	}

	result1 := calculate(5, 3, add)
	fmt.Println(result1) // 输出：8

	result2 := calculate(5, 3, subtract)
	fmt.Println(result2) // 输出：2

	// defer func() {
	// 	fmt.Println("defer func ...")
	// }()
	// fmt.Println("no defer ...")

	f := func() {
		panic("has err ...")
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	f()
}
