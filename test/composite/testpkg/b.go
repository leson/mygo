package testpkg

import (
	"fmt"
	"mygo/pkg/utils"
)

func TestUtilInB() {
	fmt.Println("call test util in test/composite/testpkg/b.go")
	utils.TestUtil()
}
