package composite

import "fmt"

var (
	a string = "hello"
	b string = "world"
)

func TestStr() {
	url := "abc=%v&def=%v\n"
	str_url := fmt.Sprintf(url, a, b)
	fmt.Printf("%s", str_url)
}
