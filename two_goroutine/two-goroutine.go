package two_goroutine

import (
	"fmt"
	"sync"
	"time"
)

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
// 更好的解决方案（使用 WaitGroup）
func GetOne() {
	var wg sync.WaitGroup

	a := []int{1, 3, 5, 7, 9}
	wg.Add(1)
	go func() { // 打印奇数
		defer wg.Done()
		for _, v := range a {
			time.Sleep(50 * time.Millisecond)
			fmt.Println("奇数：", v)
		}
	}()

	b := []int{2, 4, 6, 8, 10}
	wg.Add(1)
	go func() { // 打印偶数
		defer wg.Done()
		for _, v := range b {
			time.Sleep(50 * time.Millisecond)
			fmt.Println("偶数：", v)
		}
	}()

	// time.Sleep(3 * time.Second)   // time.Sleep
	// <-time.After(3 * time.Second) // 从 channel 接收值实现等待
	wg.Wait() // 等待所有 goroutine 完成
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

// Task 定义任务类型
type Task struct { // 使用AI
	Name string
	Fn   func() error
}

// TaskResult 存储任务执行结果
type TaskResult struct {
	Name     string
	Duration time.Duration
	Error    error
}

// TaskScheduler 任务调度器
type TaskScheduler struct {
	tasks   []Task
	results []TaskResult
	mu      sync.Mutex
}

// NewTaskScheduler 创建新的任务调度器
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks:   []Task{},
		results: []TaskResult{},
	}
}

// AddTask 添加任务到调度器
func (ts *TaskScheduler) AddTask(name string, fn func() error) {
	ts.tasks = append(ts.tasks, Task{Name: name, Fn: fn})
}

// Execute 并发执行所有任务
func (ts *TaskScheduler) Execute() {
	var wg sync.WaitGroup

	// 为每个任务启动一个协程
	for _, task := range ts.tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()

			// 记录开始时间
			startTime := time.Now()

			// 执行任务
			err := t.Fn()

			// 计算执行时间
			duration := time.Since(startTime)

			// 将结果存储到调度器中（需要使用锁保护）
			ts.mu.Lock()
			ts.results = append(ts.results, TaskResult{
				Name:     t.Name,
				Duration: duration,
				Error:    err,
			})
			ts.mu.Unlock()

			// 打印任务执行信息
			if err != nil {
				fmt.Printf("✗ 任务 '%s' 执行失败，耗时 %v，错误：%v\n", t.Name, duration, err)
			} else {
				fmt.Printf("✓ 任务 '%s' 执行完成，耗时 %v\n", t.Name, duration)
			}
		}(task)
	}

	// 等待所有任务完成
	wg.Wait()
}

// GetResults 获取所有任务的执行结果
func (ts *TaskScheduler) GetResults() []TaskResult {
	return ts.results
}

// PrintSummary 打印执行统计摘要
func (ts *TaskScheduler) PrintSummary() {
	fmt.Println("\n========== 任务执行摘要 ==========")
	var totalDuration time.Duration
	for _, result := range ts.results {
		totalDuration += result.Duration
		status := "成功"
		if result.Error != nil {
			status = "失败"
		}
		fmt.Printf("任务: %-15s | 状态: %-4s | 耗时: %10v\n", result.Name, status, result.Duration)
	}
	fmt.Printf("总耗时: %v\n", totalDuration)
	fmt.Println("==================================")
}

// GetTwo 演示任务调度器的使用
func GetTwo() {
	// 创建任务调度器
	scheduler := NewTaskScheduler()

	// 添加示例任务
	scheduler.AddTask("任务1", func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("  → 任务1 的具体工作内容")
		return nil
	})

	scheduler.AddTask("任务2", func() error {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("  → 任务2 的具体工作内容")
		return nil
	})

	scheduler.AddTask("任务3", func() error {
		time.Sleep(800 * time.Millisecond)
		fmt.Println("  → 任务3 的具体工作内容")
		return nil
	})

	scheduler.AddTask("任务4", func() error {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("  → 任务4 的具体工作内容")
		return nil
	})

	// 执行所有任务（并发）
	fmt.Println("开始执行任务...")
	scheduler.Execute()

	// 打印执行摘要
	scheduler.PrintSummary()
}
