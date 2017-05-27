package main //select 发送
//空的Select完全阻塞。
//可以设置超时

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}
}
