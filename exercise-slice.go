package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%v, cap=%v\n", s, len(s), cap(s))
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	// 虽然s1的容量为8，长度为3，但是只能打印3个元素
	printSlice(slice)
	printSlice(s1)
	// 访问长度以外的元素会产生运行时panic
	// fmt.Println(s1[3])

	// 对slice进行切片操作，可以理解实际上是对底层数组进行切片操作，不过切片的索引按照slice的索引来
	// 所以超出原slice长度的元素会使用底层数组的值来填充
	s2 := s1[2:6:7]
	printSlice(s2)
	s3 := s2[0:5]
	printSlice(s3)
	// 目前slice、s1、s2三个slice共享同一个底层数组，修改任意一个，都会影响到其他两个
	// 新增一个元素，没有超出s2的容量，没触发扩容操作
	s2 = append(s2, 100)
	// 再次新增元素，触发扩容，s2使用新的底层数组
	s2 = append(s2, 100)

	s1[2] = 20

	printSlice(slice)
	printSlice(s1)
	printSlice(s2)
}
