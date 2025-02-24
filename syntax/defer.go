package main

import "fmt"

func deferCloseTestA() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("A: %v &: %p\n", i, &i)
		}()
	}
}

func deferCloseTestB() {
	for i := 0; i < 10; i++ {
		defer func(v int) {
			fmt.Printf("C: %v &: %p\n", v, &v)
		}(i)
	}
}

func deferCloseTestC() {
	for i := 0; i < 10; i++ {
		defer func() {
			j := i
			fmt.Printf("B: %v &: %p\n", j, &j)
		}()
	}
}
