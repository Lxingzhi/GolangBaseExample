package main //goroutine练习 创建一个goroutine，与主线程按顺序相互发送信息若干次并打印。
import (
	"fmt"
)

var c chan string

func Pingpang() { //先接受然后发送
	i := 0
	for {
		fmt.Println(<-c)                             //取主线程的数据
		c <- fmt.Sprintf("From Pingpang: Hi,#%d", i) //存入数据
		i++
	}
}
func main() {
	c = make(chan string)
	go Pingpang()
	for i := 0; i < 10; i++ { //先发送然后接收
		c <- fmt.Sprintf("From Pingpang: Hello,#%d", i) //存入数据
		fmt.Println(<-c)
	}
}
