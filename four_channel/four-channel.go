package four_channel

import (
	"fmt"
	"sync"
	"time"
)

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
func GetOne() {
	// 创建无缓冲通道
	// 无缓冲即缓冲=0，所有从发送0一开始就阻塞了。类似：ch := make(chan int, 0)
	ch := make(chan int) // ！！！1、无缓冲通道（同步通道）// 发送和接收必须同时准备好
	// 有缓冲时，当通道被写满后，通道会阻塞
	// ch := make(chan int, 2) // ！！！2、带缓冲通道（异步通道）// 可以缓冲2个元素
	// ！！！3、单向通道

	// 第一个goroutine：生成从1到10的整数，并将这些整数发送到通道中
	go func() {
		defer close(ch)
		for i := 1; i <= 10; i++ {
			fmt.Printf("生产者: 发送前 %d\n", i)
			ch <- i // 发送数据到通道
			fmt.Printf("生产者: 发送后 %d\n", i)
			// time.Sleep(1 * time.Second) // 模拟处理时间
		}

	}()

	// 第二个goroutine：从通道中接收这些整数并打印出来
	go func() {
		for v := range ch { // range 会在通道关闭后自动退出
			fmt.Printf("消费者: 接收 %d\n", v)
			// time.Sleep(1 * time.Second) // 模拟处理时间
		}
	}()

	time.Sleep(3 * time.Second)
}

func sendOnly(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	for i := 1; i <= 100; i++ {
		fmt.Println("生产前：", i)
		ch <- i
		fmt.Println("生产后：", i)
	}

}

func receiveOnly(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("消费：", v)
	}

}

// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：指针运算、切片操作。
func GetTwo() {
	// 单向通道（更安全的设计）
	// 创建一个双向通道
	ch := make(chan int, 100)
	var wg sync.WaitGroup
	wg.Add(2)

	// 仅发送/生产
	go sendOnly(ch, &wg)

	time.Sleep(1 * time.Microsecond)

	// 仅接收/消费
	go receiveOnly(ch, &wg)

	wg.Wait()

}
