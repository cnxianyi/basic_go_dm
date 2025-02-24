package method

import "fmt"

type A struct{}

// 组合A
type AA struct {
	A
}

// 组合A
type AB struct {
	A
}

// A 实现Hello 方法
func (v *A) Hello() string {
	return "A"
}

// A 实现 SayHello
func (v *A) SayHello() {
	// 这里的 v是 A
	fmt.Println("hello ", v.Hello())
}

// AA 实现 SayHello
func (v *AA) SayHello() {
	// 这里的 v是AA
	// 因此使用的是 AA的Hello()
	fmt.Println("hello ", v.Hello())
}

func (v *AA) Hello() string {
	return "AA"
}

func ExtendTest() {
	var a A
	a.SayHello()

	var aa AA
	aa.SayHello()

	var ab AB
	ab.SayHello()
	// ab 没有自己的 Hello()
	// 则使用组合的 A 的 Hello
}
