package base

import "fmt"

const MAX int = 3

func TestPointer() {
	itA, itB := 19, 29
	swap(&itA, &itB)
	fmt.Printf("itA = %v\nitB=%v\n", itA, itB)
	var ptr *int = &itB
	fmt.Printf("Address : %v point to %v \n", ptr, *ptr)

	/**/
	testArrayPtr()
	testPtrPtr()
}

func testArrayPtr() {
	arr := []int{10, 100, 1000}
	arrLen := len(arr)
	var ptr [MAX]*int
	for i := 0; i < arrLen; i++ {
		ptr[i] = &arr[i] /* 整数地址赋值给指针数组 */
	}

	for i := 0; i < arrLen; i++ {
		fmt.Printf("arr[%v]=%v\n", i, *ptr[i])
	}
}

func testPtrPtr() {
	val := 10000
	var ptr *int
	var pptr **int
	/*pointer ptr address*/
	ptr = &val

	/*point pointer prt address*/
	pptr = &ptr
	fmt.Printf("%v\n", ptr)
	fmt.Printf("%v\n", *pptr)
	fmt.Printf("%v\n", pptr)
	fmt.Printf("%v\n", val)
	fmt.Printf("%v\n", *ptr)
	fmt.Printf("%v\n", **pptr)
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
