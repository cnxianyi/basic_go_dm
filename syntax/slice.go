package main

import "fmt"

func Arr() {
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(cap(arr)) // 容量
}

func Slice() {
	// make 初始化可以 指定长度len和容量cap 无法初始化值
	// 只穿一个 如 []int, 5. 则len和cap都为5
	var s1 = make([]int, 3, 5)
	fmt.Printf("s2: %v len:%d cap:%d\n", s1, len(s1), cap(s1))

	var s2 = make([]int, 5)
	fmt.Printf("s2: %v len:%d cap:%d\n", s2, len(s2), cap(s2))

	// Go会自动扩容 cap.
	// 一般是直接翻倍 cap*2
	// 如果翻倍后仍不足. 最终cap会是扩容后的长度+1 如len=13 cap=14
	s2 = append(s2, []int{6, 7, 8, 9, 10, 11, 12, 13}...) // append 追加一个元素
	fmt.Printf("s2: %v len:%d cap:%d\n", s2, len(s2), cap(s2))

	// 不指定容量时
	// cap 与 len一致
	s3 := []int{1, 2, 3}
	fmt.Printf("s3: %v len:%d cap:%d\n", s3, len(s3), cap(s3))

	// 子切片的 cap 取决于start. 是start到原切片的末尾
	// 切的是index 从start开始到end前结束. 包括start不包括end.
	// 左闭右开. 左边start使用. 右边end不使用
	s4 := s2[5:10] // 原切片cap - 子切片start -> 14 - 5 = 9
	fmt.Printf("s4: %v len:%d cap:%d\n", s4, len(s4), cap(s4))

	// 只有end时. end默认为0. 即cap等于原切片cap
	s5 := s2[:5] // 原切片cap - 子切片start -> 14 - 5 = 9
	fmt.Printf("s5: %v len:%d cap:%d\n", s5, len(s5), cap(s5))

	// 扩容操作
	// 创建一个区域. 将原有内容复制进去.
	// 直接复制. 所以 原切片和子切片任意一个触发扩容后
	// 两者内存都不会在共享
	s6 := []int{1, 2, 3, 4, 5}
	s7 := s6[2:]

	fmt.Printf("s6: %v len:%d cap:%d\n", s6, len(s6), cap(s6))
	fmt.Printf("s7: %v len:%d cap:%d\n", s7, len(s7), cap(s7))
	// 共享内存. 如子切片还有容量、直接修改已有元素等
	s6[1] = 100
	fmt.Printf("s6: %v len:%d cap:%d\n", s6, len(s6), cap(s6))
	fmt.Printf("s7: %v len:%d cap:%d\n", s7, len(s7), cap(s7))

	// 子切片触发扩容后
	s7 = append(s7, 101)
	fmt.Printf("s6: %v len:%d cap:%d\n", s6, len(s6), cap(s6))
	fmt.Printf("s7: %v len:%d cap:%d\n", s7, len(s7), cap(s7))
}
