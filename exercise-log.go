package main

import "log"

func main() {
	//使用log包预定义的Logger来打印日志
	log.Println("我是一条日志信息")
	log.Printf("=========> %d, %s", 1, "tifa")

	//log.Fatal[f|ln]函数在写入日志后，会调用os.Exit(1)退出进程
	//log.Fatalln("警告警告")
	//Panic[f|ln]相当于Print后调用了panic
	//log.Panicln()

	//使用自定义的Logger
}
