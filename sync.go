package main

import (
	"sync"
	"time"
	"fmt"
)

func sync1() {
	fmt.Println("start")

	// 一个[]byte的对象池，每个对象为一个[]byte
	var bytePool = sync.Pool{
		New: func() interface{} {
			b := make([]byte, 1024)
			return &b
		},
	}

	// 不使用对象池	
	a := time.Now().Unix()
	for i := 0; i < 100000000; i++{
		obj := make([]byte,1024)
		_ = obj
	}
	b := time.Now().Unix()

	// 使用对象池
	for i := 0; i < 100000000; i++{
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
	c := time.Now().Unix()


	fmt.Println("without pool ", b - a, "s")
	fmt.Println("with    pool ", c - b, "s")
}

func sync2() {
	fmt.Println("aaa")
}