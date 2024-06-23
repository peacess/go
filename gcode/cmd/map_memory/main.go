package main

import (
	"fmt"
	"runtime"
	"sync"
)

// see https://mp.weixin.qq.com/s/DxLEazUs-1JpdhTfadv_VA
// map的元素全部删除后，仍然留有内存，这是由于map的是由只增加的存储桶实现的原因，元素删除，但是存储桶不会删除
func main() {
	{
		n := 100_000
		m := make(map[int][128]byte)
		printAlloc()

		for i := 0; i < n; i++ {
			m[i] = [128]byte{}
		}
		printAlloc()

		clear(m) //
		// for i := 0; i < n; i++ { // Deletes 1 million elements
		// 	delete(m, i)
		// }

		runtime.GC()
		printAlloc()
		fmt.Println(len(m)) //Keeps a reference to m
	}

	{
		n := 100_000
		m := &sync.Pool{
			New: func() any {
				return [128]byte{}
			},
		}
		printAlloc()

		for i := 0; i < n; i++ {
			m.Put([128]byte{})
		}
		printAlloc()

		// clear(m) //
		for i := 0; i < n; i++ { // Deletes 1 million elements
			m.Get()
		}

		runtime.GC()
		printAlloc()
	}
	// fmt.Println(len(m)) //Keeps a reference to m
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
