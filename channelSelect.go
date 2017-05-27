package main //select 随机处理 channel

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string) //接收
	o := make(chan bool, 2)                     //其中任意一个关闭时
	go func() {
		a, b := false, false
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					if !a {
						o <- true
						a = true
					}
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					if !b {
						o <- true
						b = true
					}

					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1 //向c1传值
	c2 <- "hi"
	c2 <- "hello"
	c1 <- 33

	close(c1) //根据实际情况判断关闭几个
	close(c2)

	for i := 0; i < 2; i++ {
		<-o //读值  然后程序退出
	}

}
