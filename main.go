package main

import (
	"fmt"
)

func test() {
	defer func(a int) {
		if err := recover(); err != nil {
			fmt.Println("我是recover，发生了panic，被我捕获到了")
			fmt.Println(err)
			fmt.Println(a)
		}
	}(111)
	a := 1
	b := 0
	c := a / b
	fmt.Println(c)
}

func main() {
	a := struct {
		name string
	}{name: "aaa"}
	// 把a复制了一份赋给了变量b
	b := a
	b.name = "111"
	fmt.Println(a)
	fmt.Println(b)

	c := []int{1, 2, 3}
	d := c
	d[0] = 20
	fmt.Println(c)
	fmt.Println(d)
}
