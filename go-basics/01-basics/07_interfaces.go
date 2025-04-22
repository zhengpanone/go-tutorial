package main

import (
	"fmt"
	"io"
	"math"
	"strings"
)

// 1. 定义基本接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2. 实现Shape接口的结构体
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 3. 定义接口组合
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// 组合接口
type ReadWriter interface {
	Reader
	Writer
}

// 4. 实现自定义的String接口
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d岁", p.Name, p.Age)
}

// 5. 定义一个通用的打印函数（使用空接口作为参数）
func printAny(v interface{}) {
	fmt.Printf("值: %v, 类型: %T\n", v, v)
}

// 6. 定义一个使用类型断言的函数
func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("整数: %d\n", v)
	case string:
		fmt.Printf("字符串: %s\n", v)
	case Person:
		fmt.Printf("人: %s\n", v)
	default:
		fmt.Printf("未知类型: %T\n", v)
	}
}

// 7. 定义一个简单的数据存储接口
type Storage interface {
	Save(data string) error
	Load() (string, error)
}

// 实现内存存储
type MemoryStorage struct {
	data string
}

func (m *MemoryStorage) Save(data string) error {
	m.data = data
	return nil
}

func (m *MemoryStorage) Load() (string, error) {
	return m.data, nil
}

// 实现文件存储（模拟）
type FileStorage struct {
	filename string
}

func (f *FileStorage) Save(data string) error {
	// 这里简化实现，实际应该写入文件
	f.filename = "data.txt"
	return nil
}

func (f *FileStorage) Load() (string, error) {
	// 这里简化实现，实际应该从文件读取
	return "从文件读取的数据", nil
}

func main() {
	fmt.Println("=== 接口示例 ===")

	// 1. 使用Shape接口
	fmt.Println("\n--- Shape接口示例 ---")
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 3},
	}

	for _, shape := range shapes {
		fmt.Printf("图形面积: %.2f, 周长: %.2f\n",
			shape.Area(), shape.Perimeter())
	}

	// 2. 空接口示例
	fmt.Println("\n--- 空接口示例 ---")
	printAny(42)
	printAny("Hello")
	printAny(true)
	printAny(Person{Name: "张三", Age: 25})

	// 3. 类型断言
	fmt.Println("\n--- 类型断言示例 ---")
	var i interface{} = "Hello, Go"

	// 安全的类型断言
	if str, ok := i.(string); ok {
		fmt.Printf("字符串值: %s\n", str)
	}

	// 使用类型选择
	fmt.Println("\n--- 类型选择示例 ---")
	describe(42)
	describe("Hello")
	describe(Person{Name: "李四", Age: 30})
	describe(3.14)

	// 4. 实现String接口
	fmt.Println("\n--- String接口示例 ---")
	p := Person{Name: "王五", Age: 35}
	fmt.Println(p) // 自动调用String()方法

	// 5. 接口组合示例
	fmt.Println("\n--- 接口组合示例 ---")
	var rw ReadWriter = &strings.Builder{} // strings.Builder实现了ReadWriter接口
	rw.Write([]byte("Hello"))
	fmt.Printf("ReadWriter类型: %T\n", rw)

	// 6. 接口作为函数参数
	fmt.Println("\n--- 接口作为函数参数示例 ---")
	var writer io.Writer = &strings.Builder{}
	fmt.Fprintf(writer, "Hello, %s!", "World")

	// 7. 存储接口示例
	fmt.Println("\n--- 存储接口示例 ---")
	// 使用内存存储
	memStorage := &MemoryStorage{}
	memStorage.Save("测试数据")
	if data, err := memStorage.Load(); err == nil {
		fmt.Printf("从内存读取: %s\n", data)
	}

	// 使用文件存储
	fileStorage := &FileStorage{}
	fileStorage.Save("文件数据")
	if data, err := fileStorage.Load(); err == nil {
		fmt.Printf("从文件读取: %s\n", data)
	}

	// 8. nil接口示例
	fmt.Println("\n--- nil接口示例 ---")
	var s Shape
	fmt.Printf("nil接口: %v, 是否为nil: %t\n", s, s == nil)

	// 9. 接口的动态值和动态类型
	fmt.Println("\n--- 接口的动态值和动态类型 ---")
	var x interface{}
	x = 42
	fmt.Printf("动态类型: %T, 动态值: %v\n", x, x)
	x = "Hello"
	fmt.Printf("动态类型: %T, 动态值: %v\n", x, x)

	// 10. 接口最佳实践示例
	fmt.Println("\n--- 接口最佳实践示例 ---")
	// 定义小接口
	type Reader interface {
		Read(p []byte) (n int, err error)
	}
	type Writer interface {
		Write(p []byte) (n int, err error)
	}
	// 组合小接口
	type ReadWriter interface {
		Reader
		Writer
	}

	// 使用最小所需接口
	var buf strings.Builder
	writeData(&buf) // 只需要Writer接口
}

// 只接收所需的最小接口
func writeData(w io.Writer) {
	w.Write([]byte("数据"))
}
