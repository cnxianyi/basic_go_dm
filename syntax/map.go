package main

import "fmt"

func Map() {
	m1 := map[string]int{}
	fmt.Println(len(m1))

	m2 := make(map[string]int)
	fmt.Println(len(m2))

	// 通过键访问. map会额外返回一个是否存在的值
	v, ok := m2["someKey"]
	if ok {
		fmt.Println("someKey 存在", v)
	}

	// 删除.
	// 并不会返回任何内容
	delete(m2, "someKey")

	// comparable 可比较的
	// map的键和switch的case 都必须要是comparable类型
}
