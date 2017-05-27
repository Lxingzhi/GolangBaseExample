package main //高并发，有十个缓存，多核一起运行 第二种解决方案 是同步包

import (
	"fmt"
	"runtime" //获取电脑核数
	"sync"    //同步包
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //高并发
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}
	wg.Wait()
}

func Go(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()

}
