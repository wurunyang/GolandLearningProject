package main

// 150个线程同时往帐户里存入1块钱

import (
	"fmt"
	"sync"
)

// Account 定义一个结构体，模拟帐户，里边包含一个互斥锁
type Account struct {
	count int
	mu    sync.Mutex
}

// Add 新增一个给帐户存钱的方法
func (a *Account) Add() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.count++
}

func main() {
	account := &Account{count: 0}
	// 增加协程计数器，防止主线程在子线程结束之前退出
	wg := sync.WaitGroup{}
	wg.Add(150)

	for i := 0; i < 150; i++ {
		go func() {
			defer wg.Done() // 协程执行完毕后，计数器的值减一
			account.Add()
		}()
	}
	wg.Wait() // 等待所有协程结束
	fmt.Println(account.count)
}
