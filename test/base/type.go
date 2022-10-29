package base

import (
	"fmt"
	"time"
)

const (
	PI = 3.14159265358979343328
	c1 = "hello"
)

var (
	a = 1
	b = "hello world"
)

// variable alias
type (
	F16 int16
	Str string
)

func TestType() {
	var in int
	var flt F16
	var bl bool
	var str Str
	var arr [2]bool

	fmt.Printf("==> int default value:[%d];\n==> float default value:[%v];\n==> bool default value:[%t];\n==> string default value:[%s];\n==> array default value:[%v];\n", in, flt, bl, str, arr)
	Test_bool()
	Test_time()
}

func Test_bool() {
	b := true
	fmt.Printf("%t", b)
}

func Test_time() {
	currTime := time.Now()
	fmt.Printf("%s\n", currTime)
	y := time.Now().Year()
	m := time.Now().Month()
	d := time.Now().Day()
	h := time.Now().Hour()
	M := time.Now().Minute()
	s := time.Now().Second()
	date_time := time.Date(y, m, d, h, M, s, 0, time.Local)
	fmt.Printf("%s\n", date_time)
}
