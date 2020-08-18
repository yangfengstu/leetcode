package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 生产者
func Send(ch chan int) {
	x := 0
	defer func() {
		if err := recover(); err != nil && err.(runtime.Error).Error() == "send on closed channel" {
			fmt.Println(err)
			fmt.Println("即将产生的数据：", x)
		} else {
			close(ch) //关闭的目的：不在发送数据
		}
		wg.Done()
	}()
	for i := 0; i < 10; i++ {
		x++
		ch <- x
	}
}

// 消费者
func Receive(ch chan int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			close(ch)         //关闭的目的：不要让生产者继续发送数据
			fmt.Println(<-ch) //继续消费，输出结果为0,说明已经不会生产者已经不会再发送数据了
		}
		wg.Done()
	}()
	for x := range ch {
		time.Sleep(time.Second)
		fmt.Println(x)
		if x == 3 {
			panic("发生意外的错误") //中断主程序,但是协程还是不会关闭的
		}
	}
	fmt.Println("Receive任务结束")
}

func main() {
	fmt.Println("退出程序时，防止channel没有消费完")
	ch := make(chan int)
	wg.Add(2)
	go Send(ch)
	go Receive(ch)
	wg.Wait()
	fmt.Println("任务完成")
	_, ok := <-ch
	fmt.Println(ok)
}
