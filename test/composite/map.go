package composite

import "fmt"

func TestMap() {
	var m1 map[string]string
	m1 = make(map[string]string)
	fmt.Printf("m1: %v\n", m1)
	testMapLoop()
	testMapDelete()
}

func testMapLoop() {
	m1 := map[string]string{"France": "Paris", "Italy": "Rome", "China": "BeiJing", "India": "New delhi"}
	fmt.Printf("Country Map: %v\n", m1)

	for ctry, cap := range m1 {
		fmt.Printf("==> Country:[%v]'s Capital:[%v]\n", ctry, cap)
	}

	for ctry := range m1 {
		fmt.Printf("Country:[%v]'s Capital:[%v]\n", ctry, m1[ctry])
	}

}

func testMapDelete() {
	m1 := map[string]string{"name": "Leson", "country": "China", "gender": "male"}
	m2 := make(map[string]string)
	m2["name"] = "Julia"
	m2["gender"] = "Femal"
	for _, v := range m1 {
		fmt.Printf("v: %v\n", v)
	}

	for _, v := range m2 {
		fmt.Printf("v: %v\n", v)
	}
}
