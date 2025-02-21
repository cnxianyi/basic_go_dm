package main

import "fmt"

// Closure 闭包
func Closure() func() int {
	var t int
	return func() int {
		t++
		return t
	}
}

// defer
// 一个函数最多8个defer
// 后进先出
func Defer() {
	var i int = 0
	defer func() {
		fmt.Println(i) // 2
	}() // 返回最后的i

	i = 1

	defer func(j int) {
		fmt.Println(j) // 1
	}(i) // 此时将1传入了

	i = 2
} // 1 2

func DeferTest1() int { // 0
	i := 0
	defer func() {
		i = 1
	}()
	return i // 此时 i已经被写入内存
}

// 当返回值有命名时. defer中能修改i
func DeferTest2() (i int) { // 1
	i = 0
	defer func() {
		i = 1
	}()
	return i
}

func DeferTest3() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func DeferTest4() {
	for i := 0; i < 10; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func DeferTest5() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			fmt.Println(j)
		}()
	}
}

type Users struct {
	Name string
}

// 不要对可迭代对象取地址
func ForTest() {
	users := []Users{
		{
			Name: "Ilya",
		},
		{
			Name: "Mike",
		},
	}

	m := make(map[string]*Users, 2)
	for _, v := range users {
		m[v.Name] = &v
		fmt.Printf("%p\n", &v)
	}

	for k, v := range m {
		fmt.Printf("name: %s , user:%v\n", k, v)
	}

}
