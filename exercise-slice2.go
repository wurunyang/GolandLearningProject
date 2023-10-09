package main

import "fmt"

func myAppend(s []int) []int {
	// 这里 s 虽然改变了，但并不会影响外层函数的 s
	// 虽然外层函数中s的底层数组也变化了，但是它的len和cap没有变化，所以打印的结果还是和原来一样
	s = append(s, 100)
	return s
}

func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
	return
}

func main() {
	// 这种情况就可以发现s的底层数组被改变了。但是由于len没变，所以无法感知到
	//s := make([]int, 3, 4)
	//newS := myAppend(s)
	//s = append(s, 222)

	//这种情况没有改变s的底层数组，因为发生了扩容，s在函数里的副本指向了新的底层数组
	s := []int{1, 1, 1}
	newS := myAppend(s)

	fmt.Println(s)
	fmt.Println(newS)

	s = newS

	myAppendPtr(&s)
	fmt.Println(s)
}
