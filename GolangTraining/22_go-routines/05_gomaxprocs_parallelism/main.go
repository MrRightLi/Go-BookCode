package main

import (
	"fmt"
	"runtime" // 提供了Go 语言运行时配置参数的能力
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	// 给每个可用的物理处理器分配一个逻辑处理器
	// runtime.NumCPU()返回可用使用的物理处理器的数量
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("----------------",runtime.NumCPU()) // mac 17 --->8
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}
