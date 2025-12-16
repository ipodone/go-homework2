package five_mu

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 核心原则：锁的粒度要尽可能小，只在真正需要保护共享资源时才加锁。

// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
func GetOne() { // 版本1
	fmt.Println("=== 版本1 ===")

	count := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock() // 加锁保护共享变量 j、count
			defer mu.Unlock()
			for j := 1; j <= 10000; j++ {
				count++
			}
		}()
	}

	wg.Wait()
	duration := time.Since(startTime)
	fmt.Println("count:", count, "Time:", duration)
}

func GetTwo() { // 版本2：优化性能 - 减少锁竞争
	fmt.Println("=== 版本2：优化性能（减少锁竞争）===")

	count := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 本地计数，减少锁竞争
			localCount := 0
			for j := 1; j <= 10000; j++ {
				localCount++
			}

			// 只锁一次更新全局计数
			mu.Lock()
			count += localCount
			mu.Unlock()
		}()
	}

	wg.Wait()
	duration := time.Since(startTime)
	fmt.Println("count:", count, "Time:", duration)
}

// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。
func GetThree() { // 版本3：使用原子操作
	fmt.Println("=== 版本3：使用原子操作 ===")

	var count int64
	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			localCount := int64(0)
			for j := 1; j <= 10000; j++ {
				localCount++
			}

			// 使用原子操作
			atomic.AddInt64(&count, localCount)
		}()
	}

	wg.Wait()
	duration := time.Since(startTime)
	fmt.Println("count:", count, "Time:", duration)
}

func GetFour() { // 版本4：使用sync/atomic包
	fmt.Println("=== 版本4：直接使用原子递增 ===")

	var count int64
	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 1; j <= 10000; j++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}

	wg.Wait()
	duration := time.Since(startTime)
	fmt.Println("count:", count, "Time:", duration)
}

func GetFive() { // 版本5：通道实现
	fmt.Println("=== 版本5：使用通道 ===")

	count := 0
	resultCh := make(chan int, 10)

	startTime := time.Now()

	for i := 1; i <= 10; i++ {
		go func() {
			localCount := 0
			for j := 1; j <= 10000; j++ {
				localCount++
			}
			resultCh <- localCount
		}()
		count += <-resultCh
	}

	duration := time.Since(startTime)
	fmt.Println("count:", count, "Time:", duration)

}
