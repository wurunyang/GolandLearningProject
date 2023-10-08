package main

import (
	"context"
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
	context.Background()
}
