package composite

import "fmt"

func TestArr() {
	arr := []int{1, 3, 5, 7, 9} // slice ectually
	avg := GetAverage(arr)
	fmt.Printf("%v\n", avg)
}

func GetAverage(arr []int) float32 {
	var sum int
	var avg float32
	aLen := len(arr)
	for i := 0; i < aLen; i++ {
		sum += arr[i]
	}
	avg = float32(sum) / float32(aLen)
	return avg
}
