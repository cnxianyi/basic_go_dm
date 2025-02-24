package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello())
	DoMath()
	DoStr()

	fn := Closure()
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())

	Defer()

	fmt.Println(DeferTest1())
	fmt.Println(DeferTest2())

	fmt.Println("_--------")

	DeferTest3()

	DeferTest4()
	DeferTest5()

	ForTest()

	Arr()
	Slice()

	Map()

	deferCloseTestA()
	deferCloseTestB()
	deferCloseTestC()
}
