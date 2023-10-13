package main

import (
	"encoding/json"
	"fmt"
)

// 本代码用于记录go语言中对json的使用
/*
`` 反引号的用法和python的使用方式是一致的，用于表示原始字符串字面量，其中的内容会保持原样，包括换行符和特殊字符，不会进行转义。
同时，可以用来表示多行字符串
*/

// Person 先定义一个结构体类型，将结构体的字段和JSON数据的键对应
// 在结构体的定义中，类型后面的字符串是用于指定结构体字段的标签（Tag），用于存储和传递额外信息。一般使用反引号``包围的字符串。
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 编码 JSON
	person := Person{
		Name: "Tifa",
		Age:  18,
	}
	// 返回的是byte[]，打印的话需要转成string
	jsonData, _ := json.Marshal(person)
	fmt.Println(string(jsonData))

	// 解码 JSON
	jsonData = []byte(`{"name": "Alice", "age": 1}`)
	var decodedPerson Person
	_ = json.Unmarshal(jsonData, &decodedPerson)
	fmt.Printf("name: %s, age: %d", decodedPerson.Name, decodedPerson.Age)
}
