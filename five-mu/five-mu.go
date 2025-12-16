package five_mu

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 核心原则：锁的粒度要尽可能小，只在真正需要保护共享资源时才加锁。
// 总结：
// 加锁位置	   并发性	       性能	     适用场景
// 循环外	   无并发（串行）	差	      需要保证操作序列的原子性
// 循环内	   高并发	       中	     每个操作独立且简单
// 本地计数（最佳）	   高并发	       优	     需要聚合结果的独立操作
// 原子操作（更简单）	   高并发	       优	      简单的计数或标志操作

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

			// for循环外使用：2、defer在循环外
			mu.Lock()
			defer mu.Unlock()

			// 在 for 循环外加锁效果：
			// 1、串行执行：每个goroutine必须等待前一个goroutine完全执行完整个循环（10000次）后才能获取锁
			// 2、性能差：相当于串行执行10×10000次操作
			// 3、无并发优势：虽然有10个goroutine，但同一时间只有一个在工作

			for j := 1; j <= 10000; j++ {
				// for循环内使用锁处理，输出结果为1的原因说明如下：
				// 关键点：defer 语句不是立即执行，而是将函数调用推迟到当前函数返回时才执行。
				// mu.Lock()         // 第1次循环：加锁
				// defer mu.Unlock() // 第1次循环：安排解锁（在函数返回时执行）
				// count++           // 执行 count = 1

				// 第2次循环：试图再次加锁，但锁已经被第1次的defer持有！
				// 因为goroutine函数还没返回，defer不会执行
				// 所以这里会死锁！（或超时后会输出1）

				// 所以正确写法：1、立即解锁 2、defer在循环外 3、本地计数（推荐）
				// 4、如果要在循环中使用defer，应该为每次迭代创建新函数（如定义线程安全计数器SafeCounter、定义方法，在此方法中使用defer）
				// 1、立即解锁
				// mu.Lock()
				count++
				// mu.Unlock()

				// 黄金规则：
				// 1、不要在循环内对同一把锁多次使用defer mu.Unlock()
				// 2、要么在循环内立即解锁：mu.Lock(); ...; mu.Unlock()
				// 3、要么在循环外加锁：mu.Lock(); defer mu.Unlock(); for {...}

				// 在 for 循环内加锁效果：
				// 1、高并发：所有goroutine可以同时进行，只在count++时竞争锁
				// 2、性能好：充分利用并发优势
				// 3、锁竞争激烈：10个goroutine×10000次=100000次锁竞争

				// 加锁位置的黄金法则：
				// 1、锁的范围应该最小化（只在必要时加锁）- 即场景2：独立计数器（锁在内或本地计数）
				// 2、区分是否需要保护整个操作序列 - 即场景1：需要原子性的批量操作（锁在外）
			}
		}()
	}

	wg.Wait()
	duration := time.Since(startTime)
	fmt.Println("count:", count, "Time:", duration)
}

func GetTwo() { // 版本2：优化性能 - 减少锁竞争 - 本地计数（最佳）
	fmt.Println("=== 版本2：优化性能（减少锁竞争）===")

	count := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 本地计数，减少锁竞争：3、本地计数（推荐）
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
func GetThree() { // 版本3：使用原子操作 - 原子操作（更简单）
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
			atomic.AddInt64(&count, localCount) // 本地计数+原子操作
		}() // 所以这块容易出现最常见的闭包问题：循环中变量捕获错误/意外的变量共享
	}

	wg.Wait()
	duration := time.Since(startTime)
	fmt.Println("count:", count, "Time:", duration)
}

func GetFour() { // 版本4：使用sync/atomic包 - 原子操作（更简单）
	fmt.Println("=== 版本4：直接使用原子递增 ===")

	var count int64
	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 1; j <= 10000; j++ {
				atomic.AddInt64(&count, 1) // 直接使用原子操作
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
			resultCh <- localCount // 本地计数+通道
		}()
		count += <-resultCh
	}

	duration := time.Since(startTime)
	fmt.Println("count:", count, "Time:", duration)

}

// 一、原子操作
// sync/atomic 包提供了底层的原子级内存操作，用于实现无锁的并发安全操作。
// 原子操作底层原理（简单说明）：原子操作是CPU级别的操作
// atomic.AddInt64 的原子性是通过硬件指令保证的，不是软件实现的：
// 伪代码说明原理
// func AddInt64(addr *int64, delta int64) int64 {
//     // CPU级别保证这些操作是原子的：
//     // 1. 从内存读取当前值到寄存器
//     // 2. 在寄存器中加上delta
//     // 3. 写回内存
//
//     // 实际上使用的是类似LOCK XADD的CPU指令
//     // 这个指令会锁定内存总线，防止其他CPU访问
// }
// 常见使用模式：1、累加本地计算结果 2、递减计数器
// 性能：比使用互斥锁更快；适用：简单的数值累加、递减操作；不适用：需要保护复杂操作序列的场景
// 原子操作是 Go 并发编程中的重要工具，特别适合高并发下的简单数值操作。

// 二、Go中锁：
// 1. 互斥锁（Mutex） - 最常用的锁 - 示例：保护共享资源
// 2. 读写锁（RWMutex） - 读多写少的场景 - 示例：缓存系统
// 3. 原子锁（atomic） - 轻量级操作 - 示例：高性能计数器
// 4. 条件变量（Cond） - 等待/通知机制 - 示例：生产者-消费者
// 5. Once（一次性锁） - 确保只执行一次 - 示例：单例模式
// 6. WaitGroup（等待组） - 协程同步
// 7. Map（并发安全映射） - Go 1.9+
// 8. Pool（对象池） - 复用对象
// 9. 自旋锁（Spinlock） - 高性能场景
//
// 选择锁的建议
// 简单共享变量 → sync.Mutex------------------------------------排序3
// 读多写少的数据 → sync.RWMutex--------------------------------------排序2
// 计数器/标志位 → sync/atomic---------------------------------------------排序1
// 等待条件满足 → sync.Cond
// 单次初始化 → sync.Once
// 等待多个goroutine → sync.WaitGroup
// 并发map → sync.Map
// 对象复用 → sync.Pool
//
// 记住：能用原子操作（排序1）就不用锁，能用读写锁（排序2）就不用互斥锁（排序3），尽量减少锁的粒度和持有时间。
//
// 对比表格
// 锁类型	     用途	       特点					性能		适用场景
// Mutex	    互斥访问		简单、通用			中等		一般共享资源保护
// RWMutex	    读写分离		读共享，写互斥		读多时高	读多写少的缓存
// atomic	    原子操作		无锁，CPU指令		非常高		计数器、标志位
// Cond	       条件等待		等待/通知机制			中等		生产者-消费者
// Once	       一次性执行		确保只执行一次		高			单例、初始化
// WaitGroup	协程同步		等待一组goroutine	高			并行任务
// Map	       并发安全map		内置并发安全		高			并发map操作
// Pool	       对象复用		减少GC压力				高			频繁创建的对象

// 三、闭包概念、闭包问题
// ！！！！！闭包==匿名函数
// ！！！！！闭包 = 函数 + 其引用环境（捕获的外部变量）；闭包捕获变量是引用/指针（浅拷贝），不是值拷贝（深拷贝）
// 核心理解：闭包 = 函数 + 函数能够访问的自由变量（定义在函数外部，但在函数内部被引用的变量）
// ！！！！！闭包/匿名函数的特性：1、捕获外部变量（闭包） 2、访问外部作用域变量
//
// 总结：
// 1、闭包是什么：一个能访问其词法作用域的函数。
// ！！！！！2、“闭包问题”：通常指由闭包引起的常见陷阱，如循环中变量捕获错误、意外的变量共享、潜在的内存泄漏等。
// 3、关键：理解闭包捕获的是变量的引用，并且其生命周期会延长被捕获变量的生命周期。
// 掌握闭包的关键在于清晰地理解作用域链和变量生命周期。
//
// 闭包的优点：尽管有陷阱，闭包非常强大：
// 1、数据封装/私有化：模拟私有变量（如模块模式）。
// 2、函数工厂：创建具有特定配置的函数。
// 3、状态保持：用于回调、事件处理、异步编程中保持状态。
// 4、函数柯里化与部分应用：函数式编程的基础。
//
// “闭包问题”通常指什么？
// 在编程实践中，提到“闭包问题”通常不是指闭包的概念本身，而是指与闭包相关的、容易导致错误的常见陷阱。主要有以下几类：
// 1. 循环中的闭包陷阱（最常见）- 循环中变量捕获错误
// for (var i = 0; i < 3; i++) {
//     setTimeout(function() {
//         console.log(i); // 你以为会输出 0, 1, 2？
//     }, 100);
// }
// // 实际输出：3, 3, 3
// 原因：三个闭包（setTimeout 的回调函数）共享同一个变量 i。当回调执行时，循环早已结束，i 的值已经是 3。
// 解决方案：
// （1）使用 let（块级作用域）：for (let i = 0; ...)，每次迭代都会创建一个新的 i 绑定。
// （2）使用 IIFE 创建新的作用域：(function(j) { setTimeout(... j ...) })(i)。
// 2. 意外的变量捕获
// 闭包捕获的是变量的引用，而不是创建闭包时的值。如果捕获的变量后续被修改，闭包看到的是修改后的值。
// let arr = [];
// for (var i = 0; i < 3; i++) {
//     arr.push(function() { return i; }); // 所有函数都捕获同一个 i
// }
// console.log(arr[0]()); // 3，不是0
// 3. 内存泄漏
// 由于闭包会持续引用其外部变量，如果闭包本身（例如一个事件监听器）长期存在（如绑定在全局对象上），那么它引用的整个作用域链都无法被垃圾回收，可能导致内存泄漏。
// function attachEvent() {
//     let hugeData = getHugeData(); // 一个很大的数据
//     document.getElementById('myButton').onclick = function() {
//         // 这个闭包捕获了 hugeData，即使点击事件里没用到它！
//         console.log('clicked');
//     };
//     // 即使函数结束，由于闭包引用，hugeData 无法被释放。
// }
