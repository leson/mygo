package main

import (
	"mygo/test/composite"

	"mygo/test/base"
	"mygo/test/composite/testpkg"
)

func main() {
	base.TestType()
	// base.TestPointer()
	// composite.TestFunc()
	composite.TestStruct()
	// composite.TestStr()
	// composite.TestSlice()
	// composite.TestMap()
	// composite.TestDefer()
	// composite.TestInterface()
	// composite.TestInterfaceNest()
	// composite.TestInterfaceOpenClosePrinciple()
	testpkg.TestUtilInB()
}
