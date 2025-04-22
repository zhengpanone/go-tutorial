// 高级设计模式示例
//
// 这个示例展示了Go语言中常用的高级设计模式和最佳实践，
// 包括创建型、结构型、行为型模式等。
package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. 创建型模式

// 1.1 单例模式
type Singleton struct {
	data string
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: "单例数据"}
	})
	return instance
}

// 1.2 工厂模式
type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d *Dog) Speak() string { return "汪汪!" }
func (c *Cat) Speak() string { return "喵喵!" }

// 工厂函数
func CreateAnimal(animalType string) Animal {
	switch animalType {
	case "dog":
		return &Dog{}
	case "cat":
		return &Cat{}
	default:
		return nil
	}
}

// 1.3 构建器模式
type Server struct {
	Host     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
}

type ServerBuilder struct {
	server *Server
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{server: &Server{}}
}

func (b *ServerBuilder) SetHost(host string) *ServerBuilder {
	b.server.Host = host
	return b
}

func (b *ServerBuilder) SetPort(port int) *ServerBuilder {
	b.server.Port = port
	return b
}

func (b *ServerBuilder) SetProtocol(protocol string) *ServerBuilder {
	b.server.Protocol = protocol
	return b
}

func (b *ServerBuilder) SetTimeout(timeout time.Duration) *ServerBuilder {
	b.server.Timeout = timeout
	return b
}

func (b *ServerBuilder) SetMaxConns(maxConns int) *ServerBuilder {
	b.server.MaxConns = maxConns
	return b
}

func (b *ServerBuilder) Build() *Server {
	return b.server
}

// 2. 结构型模式

// 2.1 适配器模式
type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (p *MyLegacyPrinter) Print(s string) string {
	return fmt.Sprintf("Legacy: %s", s)
}

type ModernPrinter interface {
	PrintModern(s string) string
}

type PrinterAdapter struct {
	LegacyPrinter
}

func (p *PrinterAdapter) PrintModern(s string) string {
	return p.Print(s)
}

// 2.2 装饰器模式
type Component interface {
	Operation() string
}

type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

type Decorator struct {
	component Component
}

func (d *Decorator) Operation() string {
	return d.component.Operation() + " + Decorator"
}

// 3. 行为型模式

// 3.1 观察者模式
type Observer interface {
	Update(string)
}

type Subject interface {
	Register(Observer)
	Deregister(Observer)
	NotifyAll()
}

type ConcreteObserver struct {
	id string
}

func (o *ConcreteObserver) Update(data string) {
	fmt.Printf("观察者 %s 收到更新: %s\n", o.id, data)
}

type ConcreteSubject struct {
	observers []Observer
	data      string
}

func (s *ConcreteSubject) Register(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *ConcreteSubject) Deregister(o Observer) {
	for i, obs := range s.observers {
		if obs == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) NotifyAll() {
	for _, obs := range s.observers {
		obs.Update(s.data)
	}
}

// 3.2 策略模式
type Strategy interface {
	Execute(int, int) int
}

type AddStrategy struct{}
type SubtractStrategy struct{}

func (s *AddStrategy) Execute(a, b int) int      { return a + b }
func (s *SubtractStrategy) Execute(a, b int) int { return a - b }

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

// 4. 并发模式

// 4.1 工作池模式
type Job struct {
	ID   int
	Data string
}

type Worker struct {
	ID      int
	JobChan chan Job
	Quit    chan bool
}

func NewWorker(id int, jobChan chan Job) *Worker {
	return &Worker{
		ID:      id,
		JobChan: jobChan,
		Quit:    make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.JobChan:
				fmt.Printf("工作者 %d 处理任务 %d: %s\n", w.ID, job.ID, job.Data)
			case <-w.Quit:
				return
			}
		}
	}()
}

// 4.2 发布订阅模式
type PubSub struct {
	mu     sync.RWMutex
	subs   map[string][]chan string
	closed bool
}

func NewPubSub() *PubSub {
	return &PubSub{
		subs: make(map[string][]chan string),
	}
}

func (ps *PubSub) Subscribe(topic string) chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan string, 1)
	ps.subs[topic] = append(ps.subs[topic], ch)
	return ch
}

func (ps *PubSub) Publish(topic string, msg string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if ps.closed {
		return
	}

	for _, ch := range ps.subs[topic] {
		ch <- msg
	}
}

// 5. 函数式编程模式

// 5.1 函数类型和高阶函数
type Operation func(int) int

func MapSlice(slice []int, op Operation) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = op(v)
	}
	return result
}

func Double(x int) int { return x * 2 }
func Square(x int) int { return x * x }

// 5.2 管道模式
type IntPipe func([]int) []int

func Pipeline(pipes ...IntPipe) IntPipe {
	return func(integers []int) []int {
		result := integers
		for _, pipe := range pipes {
			result = pipe(result)
		}
		return result
	}
}

// 6. 错误处理模式

// 6.1 自定义错误类型
type ValidationError struct {
	Field string
	Error string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("验证错误 - %s: %s", v.Field, v.Error)
}

// 6.2 错误包装
type stackTracer interface {
	StackTrace() string
}

type wrappedError struct {
	msg   string
	err   error
	stack string
}

func (w *wrappedError) Error() string {
	return fmt.Sprintf("%s: %v", w.msg, w.err)
}

func (w *wrappedError) StackTrace() string {
	return w.stack
}

func WrapError(err error, msg string) error {
	if err == nil {
		return nil
	}
	return &wrappedError{
		msg:   msg,
		err:   err,
		stack: "stack trace here", // 实际应用中应该获取真实的堆栈信息
	}
}

// 7. 依赖注入模式

// 7.1 简单的依赖注入
type Service interface {
	DoSomething() string
}

type ServiceImpl struct{}

func (s *ServiceImpl) DoSomething() string {
	return "服务执行操作"
}

type Client struct {
	service Service
}

func NewClient(service Service) *Client {
	return &Client{service: service}
}

func main() {
	// 1. 创建型模式示例
	fmt.Println("=== 创建型模式示例 ===")

	// 单例模式
	s1 := GetInstance()
	s2 := GetInstance()
	fmt.Printf("单例相同?: %v\n", s1 == s2)

	// 工厂模式
	dog := CreateAnimal("dog")
	cat := CreateAnimal("cat")
	fmt.Printf("狗说: %s\n", dog.Speak())
	fmt.Printf("猫说: %s\n", cat.Speak())

	// 构建器模式
	server := NewServerBuilder().
		SetHost("localhost").
		SetPort(8080).
		SetProtocol("http").
		SetTimeout(time.Second * 30).
		SetMaxConns(100).
		Build()
	fmt.Printf("服务器配置: %+v\n", server)

	// 2. 结构型模式示例
	fmt.Println("\n=== 结构型模式示例 ===")

	// 适配器模式
	legacy := &MyLegacyPrinter{}
	adapter := &PrinterAdapter{legacy}
	fmt.Println(adapter.PrintModern("Hello"))

	// 装饰器模式
	component := &ConcreteComponent{}
	decorator := &Decorator{component}
	fmt.Println(decorator.Operation())

	// 3. 行为型模式示例
	fmt.Println("\n=== 行为型模式示例 ===")

	// 观察者模式
	subject := &ConcreteSubject{}
	observer1 := &ConcreteObserver{id: "1"}
	observer2 := &ConcreteObserver{id: "2"}
	subject.Register(observer1)
	subject.Register(observer2)
	subject.data = "新数据"
	subject.NotifyAll()

	// 策略模式
	context := &Context{}
	context.SetStrategy(&AddStrategy{})
	fmt.Printf("10 + 5 = %d\n", context.ExecuteStrategy(10, 5))
	context.SetStrategy(&SubtractStrategy{})
	fmt.Printf("10 - 5 = %d\n", context.ExecuteStrategy(10, 5))

	// 4. 并发模式示例
	fmt.Println("\n=== 并发模式示例 ===")

	// 工作池
	jobChan := make(chan Job, 10)
	worker := NewWorker(1, jobChan)
	worker.Start()
	jobChan <- Job{ID: 1, Data: "测试数据"}
	time.Sleep(time.Second)

	// 发布订阅
	ps := NewPubSub()
	ch := ps.Subscribe("news")
	go func() {
		msg := <-ch
		fmt.Printf("收到消息: %s\n", msg)
	}()
	ps.Publish("news", "突发新闻!")
	time.Sleep(time.Second)

	// 5. 函数式编程示例
	fmt.Println("\n=== 函数式编程示例 ===")

	numbers := []int{1, 2, 3, 4, 5}
	doubled := MapSlice(numbers, Double)
	squared := MapSlice(numbers, Square)
	fmt.Printf("加倍: %v\n", doubled)
	fmt.Printf("平方: %v\n", squared)

	// 管道
	pipeline := Pipeline(
		func(nums []int) []int { return MapSlice(nums, Double) },
		func(nums []int) []int { return MapSlice(nums, Square) },
	)
	result := pipeline(numbers)
	fmt.Printf("管道结果: %v\n", result)

	// 6. 错误处理示例
	fmt.Println("\n=== 错误处理示例 ===")

	err := &ValidationError{
		Field: "用户名",
		Error: "不能为空",
	}
	wrappedErr := WrapError(err, "验证失败")
	fmt.Println(wrappedErr)

	// 7. 依赖注入示例
	fmt.Println("\n=== 依赖注入示例 ===")

	service := &ServiceImpl{}
	client := NewClient(service)
	fmt.Println(client.service.DoSomething())
}

/* 设计模式最佳实践

1. 创建型模式
   - 使用工厂方法封装对象创建
   - 使用构建器处理复杂对象构造
   - 确保单例线程安全
   - 使用原型模式优化对象创建

2. 结构型模式
   - 使用适配器处理接口不兼容
   - 使用装饰器动态扩展功能
   - 使用代理控制对象访问
   - 使用组合表示层次结构

3. 行为型模式
   - 使用观察者处理一对多关系
   - 使用策略封装算法
   - 使用命令解耦调用者和接收者
   - 使用状态处理状态转换

4. 并发模式
   - 使用工作池处理并发任务
   - 使用发布订阅处理事件
   - 正确处理goroutine生命周期
   - 避免goroutine泄漏

5. 函数式编程
   - 使用高阶函数增加代码复用
   - 使用闭包保持状态
   - 使用管道处理数据流
   - 保持函数纯净

6. 错误处理
   - 使用自定义错误类型
   - 正确包装和传播错误
   - 提供有意义的错误信息
   - 在适当层级处理错误

7. 依赖注入
   - 面向接口编程
   - 使用构造注入
   - 避免服务定位器
   - 保持依赖简单和显式
*/
