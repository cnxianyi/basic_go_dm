package main

import (
	"basic_go_dm/syntax/method"
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

	method.MethodTest()
	u1 := method.Users{
		Name: "Ilya",
		Age:  18,
	}

	u1.ChangeName("Mike")
	u1.ChangeAge(20)

	fmt.Printf("%+v \n", u1)

	method.ExtendTest()
}
