package composite

import "fmt"

func TestSlice() {
	testSliceDesc()
	testSliceLenCapCut()
	testSliceCRUD()
}

func testSliceDesc() {
	var s1 []string
	var s2 = make([]string, 9)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("len(s2): %v\n", len(s2))
	fmt.Printf("cap(s2): %v\n", cap(s2))
}

func testSliceLenCapCut() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("len(s1): %v\n", len(s1))
	fmt.Printf("cap(s1): %v\n", cap(s1))
	fmt.Printf("s1[0]: %v\n", s1[0])
	fmt.Printf("s1[:3]: %v\n", s1[:3])
	fmt.Printf("s1[3:5]: %v\n", s1[3:5])
}

func testSliceCRUD() {
	testLoopSlice()
	testSliceAdd()
	testSliceUpdate()
	testSliceDelete()
}

func testLoopSlice() {
	s1 := []string{"John", "Lucy", "Julia", "Eric"}
	for i, name := range s1 {
		fmt.Printf("index: %v;name: %v\n", i, name)
	}
}

func testSliceAdd() {
	s1 := []int{}
	s1 = append(s1, 1)
	s1 = append(s1, 10)
	fmt.Printf("s1: %v\n", s1)
}

func testSliceUpdate() {
	s1 := []int{1, 10, 100}
	s1[1] = 50
	fmt.Printf("s1: %v\n", s1)
}

func testSliceDelete() {
	s1 := []int{1, 10, 100, 1000, 10000}
	/* 公式： 想要删除的索引 i ; s1=append(s1[:i],s1[i+1:]...)*/
	s1 = append(s1[:1], s1[2:]...)
	fmt.Printf("s1: %v\n", s1)
}
