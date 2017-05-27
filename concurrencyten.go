package main //高并发，有十个缓存，多核一起运行

import (
	"fmt"
	"runtime" //获取电脑核数
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //高并发
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	for i := 0; i < 10; i++ { //多核处理需要一个判断
		<-c
	}

}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	c <- true

}
