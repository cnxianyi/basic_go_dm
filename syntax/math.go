package main

import (
	"fmt"
	"math"
)

// 常量声明 iota可以方便的初始化常量

const (
	Zero = iota
	One
	Two
	Fore  = iota + 1
	Eight = iota * 2
)

// 位运算
const (
	Zero1  = iota << 1
	One1   // 0001 -> 0010
	Two1   // 0010 -> 0100
	Three1 // 0011 -> 0110
)

func DoMath() {
	fmt.Println(math.Abs(-1))    // 绝对值
	fmt.Println(math.MaxFloat64) // f64最大值
}
