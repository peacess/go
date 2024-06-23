package main

import (
	"fmt"
	"sync"
)

// 测试当main退时，goroutine是否会自动退
// 结果是会退出，且退出的时间点是不可控制的
func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Println(i)
			if i == 0 {
				wg.Done()
			}
		}
	}()
	wg.Wait()
}
