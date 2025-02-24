package method

import "fmt"

func MethodTest() {
	var t *LinkList
	//fmt.Println(t.next) // panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Println(t)

}
