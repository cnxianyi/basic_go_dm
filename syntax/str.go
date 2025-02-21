package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func DoStr() {
	var str string = "你好"
	fmt.Println(len(str))                    // 6
	fmt.Println(utf8.RuneCountInString(str)) // 2 uft8字符长度

	// strings 包
	fmt.Println(strings.Split(str, ""))
}
