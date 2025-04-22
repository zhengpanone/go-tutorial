package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

// 1. 自定义错误类型
type MyError struct {
	When time.Time
	What string
}

// 实现error接口
func (e *MyError) Error() string {
	return fmt.Sprintf("错误发生时间：%v，错误信息：%s", e.When, e.What)
}

// 2. 返回自定义错误的函数
func doSomething() error {
	return &MyError{
		When: time.Now(),
		What: "发生了一些错误",
	}
}

// 3. 包装错误的函数
func readFile(path string) error {
	_, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("无法打开文件 %s: %w", path, err)
	}
	return nil
}

// 4. 多值返回中的错误处理
func divide2(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

// 5. 使用类型断言进行错误检查
func checkFileError(err error) {
	if err == nil {
		return
	}

	switch e := err.(type) {
	case *os.PathError:
		fmt.Printf("文件路径错误: %v\n", e)
	case *MyError:
		fmt.Printf("自定义错误: %v\n", e)
	default:
		fmt.Printf("其他错误: %v\n", e)
	}
}

// 6. 资源清理的延迟调用
func processFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close() // 确保文件最终被关闭

	// 处理文件...
	return nil
}

// 7. panic和recover的示例
func mayPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("从panic中恢复: %v\n", r)
		}
	}()

	panic("发生了一些可怕的事情")
}

// 8. 多重错误处理示例
type MultiError struct {
	Errors []error
}

func (m *MultiError) Error() string {
	var result string
	for i, err := range m.Errors {
		if i > 0 {
			result += "; "
		}
		result += err.Error()
	}
	return result
}

func validateData(data string) error {
	var errors []error

	if len(data) == 0 {
		errors = append(errors, fmt.Errorf("数据不能为空"))
	}
	if len(data) > 100 {
		errors = append(errors, fmt.Errorf("数据长度不能超过100"))
	}

	if len(errors) > 0 {
		return &MultiError{Errors: errors}
	}
	return nil
}

// 9. 错误重试函数
func retryOperation(attempts int, sleep time.Duration, f func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = f()
		if err == nil {
			return nil
		}
		if i < attempts-1 {
			time.Sleep(sleep)
			fmt.Printf("重试操作，尝试次数：%d\n", i+1)
		}
	}
	return fmt.Errorf("在%d次尝试后失败: %w", attempts, err)
}

func main() {
	fmt.Println("=== 错误处理示例 ===")

	// 1. 基本错误处理
	fmt.Println("\n--- 基本错误处理 ---")
	if err := doSomething(); err != nil {
		fmt.Printf("发生错误: %v\n", err)
	}

	// 2. 文件操作错误处理
	fmt.Println("\n--- 文件操作错误处理 ---")
	err := readFile("不存在的文件.txt")
	checkFileError(err)

	// 3. 错误包装和解包
	fmt.Println("\n--- 错误包装和解包 ---")
	if err := readFile("test.txt"); err != nil {
		// 解包装错误
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Printf("文件操作错误: %v\n", pathError)
		}
		// 检查是否包含特定错误
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("文件不存在")
		}
	}

	// 4. 除法错误处理
	fmt.Println("\n--- 除法错误处理 ---")
	if result, err := divide2(10, 0); err != nil {
		fmt.Printf("除法错误: %v\n", err)
	} else {
		fmt.Printf("结果: %f\n", result)
	}

	// 5. panic和recover
	fmt.Println("\n--- panic和recover ---")
	mayPanic()
	fmt.Println("继续执行...")

	// 6. defer的多种用法
	fmt.Println("\n--- defer的多种用法 ---")
	func() {
		defer fmt.Println("第一个defer")
		defer fmt.Println("第二个defer")
		fmt.Println("函数执行")
	}()

	// 7. 多重错误处理
	fmt.Println("\n--- 多重错误处理 ---")
	if err := validateData(""); err != nil {
		fmt.Printf("验证错误: %v\n", err)
	}

	// 8. 错误重试示例
	fmt.Println("\n--- 错误重试示例 ---")
	err = retryOperation(3, time.Second, func() error {
		return fmt.Errorf("操作失败")
	})
	if err != nil {
		fmt.Printf("最终错误: %v\n", err)
	}

	// 9. 资源清理示例
	fmt.Println("\n--- 资源清理示例 ---")
	func() {
		resource := "一些资源"
		defer func() {
			fmt.Printf("清理资源: %s\n", resource)
		}()

		// 使用资源...
		fmt.Printf("使用资源: %s\n", resource)
	}()

	// 10. 错误处理最佳实践示例
	fmt.Println("\n--- 错误处理最佳实践 ---")

	// 创建自定义错误
	var ErrInvalidInput = errors.New("无效的输入")

	// 使用自定义错误
	validate := func(input string) error {
		if input == "" {
			return fmt.Errorf("验证失败: %w", ErrInvalidInput)
		}
		return nil
	}

	if err := validate(""); err != nil {
		if errors.Is(err, ErrInvalidInput) {
			fmt.Println("检测到无效输入错误")
		}
		fmt.Printf("错误: %v\n", err)
	}
}

// 11. 实现io.Reader接口的错误处理示例
type ErrorReader struct{}

func (er ErrorReader) Read(p []byte) (n int, err error) {
	return 0, io.ErrUnexpectedEOF
}
