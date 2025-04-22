package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 1. 简单的goroutine示例
func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

// 2. 使用channel的生产者
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i // 发送数据到channel
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // 关闭channel
}

// 3. 使用channel的消费者
func consumer(ch <-chan int) {
	for num := range ch { // 从channel接收数据直到它被关闭
		fmt.Printf("接收到: %d\n", num)
	}
}

// 4. 使用互斥锁保护共享资源
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// 5. 使用读写锁的计数器
type RWCounter struct {
	rwmu  sync.RWMutex
	count int
}

func (c *RWCounter) Increment() {
	c.rwmu.Lock()
	defer c.rwmu.Unlock()
	c.count++
}

func (c *RWCounter) Value() int {
	c.rwmu.RLock()
	defer c.rwmu.RUnlock()
	return c.count
}

// 6. 工作池模式
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("工作者 %d 开始处理任务 %d\n", id, j)
		time.Sleep(100 * time.Millisecond) // 模拟工作负载
		results <- j * 2
	}
}

// 7. 使用Context控制goroutine
func controlledWorker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("工作者 %d 收到取消信号\n", id)
			return
		default:
			fmt.Printf("工作者 %d 正在工作\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	fmt.Println("=== 并发编程示例 ===")

	// 1. 基本的goroutine使用
	fmt.Println("\n--- 基本的goroutine使用 ---")
	go printNumbers()
	go printLetters()
	time.Sleep(1 * time.Second) // 等待goroutine完成

	// 2. Channel的基本使用
	fmt.Println("\n--- Channel的基本使用 ---")
	ch := make(chan int) // 创建一个无缓冲的channel
	go producer(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)

	// 3. 带缓冲的channel
	fmt.Println("\n--- 带缓冲的channel ---")
	bufferedCh := make(chan int, 3)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("发送: %d\n", i)
			bufferedCh <- i
		}
		close(bufferedCh)
	}()

	for num := range bufferedCh {
		fmt.Printf("接收: %d\n", num)
	}

	// 4. Select语句
	fmt.Println("\n--- Select语句 ---")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "来自channel 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "来自channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}

	// 5. 互斥锁示例
	fmt.Println("\n--- 互斥锁示例 ---")
	counter := SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	fmt.Printf("计数器最终值: %d\n", counter.Value())

	// 6. 原子操作
	fmt.Println("\n--- 原子操作 ---")
	var atomicCounter uint64
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint64(&atomicCounter, 1)
		}()
	}
	wg.Wait()
	fmt.Printf("原子计数器最终值: %d\n", atomic.LoadUint64(&atomicCounter))

	// 7. 工作池示例
	fmt.Println("\n--- 工作池示例 ---")
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 启动3个工作者
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集结果
	for a := 1; a <= 5; a++ {
		<-results
	}

	// 8. Context示例
	fmt.Println("\n--- Context示例 ---")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for i := 1; i <= 3; i++ {
		go controlledWorker(ctx, i)
	}
	time.Sleep(3 * time.Second)

	// 9. 定时器示例
	fmt.Println("\n--- 定时器示例 ---")
	timer := time.NewTimer(500 * time.Millisecond)
	<-timer.C
	fmt.Println("定时器触发")

	// 10. 打点器示例
	fmt.Println("\n--- 打点器示例 ---")
	ticker := time.NewTicker(200 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Printf("打点器触发 at %v\n", t)
		}
	}()
	time.Sleep(1 * time.Second)
	ticker.Stop()

	// 11. Channel超时模式
	fmt.Println("\n--- Channel超时模式 ---")
	ch = make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
	}()

	select {
	case result := <-ch:
		fmt.Printf("接收到结果: %d\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("操作超时")
	}

	fmt.Println("\n所有并发示例完成")
}
