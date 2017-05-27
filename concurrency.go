package main //并发。高并发。
import (
	"fmt"
	// "time"
)

func main() {
	//最后一个参数代表有无缓存，有缓存是异步的，无缓存是同步的
	c := make(chan bool, 1) //channel 是goroutine沟通的桥梁，大都是阻塞同步的，用make创建，close关闭。
	go func() {
		fmt.Println("googogogogog")
		c <- true //存入到chan中
		/*eg 2*/
		// close(c)
	}()
	// time.Sleep(2 * time.Second) //暂停两秒钟
	//eg2
	// <-c //如果读出来内容才继续往下走
	//eg3
	// for v := range c { //不断的循环，关闭了才不循环
	// 	fmt.Println(v)
	// }
	//eg4
	<-c //无缓存如果读出来内容才继续往下走
}
