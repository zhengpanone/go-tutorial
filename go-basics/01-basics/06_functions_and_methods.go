package main

import (
	"fmt"
)

// 1. 基本函数定义
func add(a, b int) int {
	return a + b
}

// 2. 多返回值函数
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}
	return a / b, nil
}

// 3. 命名返回值
func rectangle(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // 裸返回，会返回命名的返回值
}

// 4. 可变参数函数
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 5. 定义结构体用于演示方法
type Rectangle struct {
	Width  float64
	Height float64
}

// 6. 为Rectangle定义方法（值接收者）
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 7. 为Rectangle定义方法（指针接收者）
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// 8. 函数类型定义
type MathFunc func(int, int) int

// 9. 高阶函数：接收函数作为参数
func calculate(a, b int, op MathFunc) int {
	return op(a, b)
}

// 10. 用于演示defer的函数
func deferDemo() {
	fmt.Println("开始执行函数")
	defer fmt.Println("第一个defer")
	defer fmt.Println("第二个defer")
	fmt.Println("函数执行结束")
}

// 11. 用于演示panic和recover的函数
func panicDemo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到panic: %v\n", r)
		}
	}()

	fmt.Println("开始执行panicDemo")
	panic("发生严重错误")
	fmt.Println("这行不会执行") // 这行代码不会执行
}

func main() {
	fmt.Println("=== 函数和方法示例 ===")

	// 1. 基本函数调用
	fmt.Println("\n--- 基本函数调用 ---")
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	// 2. 多返回值函数调用
	fmt.Println("\n--- 多返回值函数 ---")
	if result, err := divide(10, 2); err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	if result, err := divide(10, 0); err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result)
	}

	// 3. 命名返回值函数调用
	fmt.Println("\n--- 命名返回值函数 ---")
	area, perimeter := rectangle(5, 3)
	fmt.Printf("矩形面积: %.2f, 周长: %.2f\n", area, perimeter)

	// 4. 可变参数函数调用
	fmt.Println("\n--- 可变参数函数 ---")
	fmt.Printf("求和结果: %d\n", sum(1, 2, 3, 4, 5))
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("切片求和结果: %d\n", sum(numbers...)) // 展开切片作为可变参数

	// 5. 方法调用
	fmt.Println("\n--- 方法调用 ---")
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Printf("矩形面积: %.2f\n", rect.Area())

	rect.Scale(2)
	fmt.Printf("放大后的矩形: %+v\n", rect)

	// 6. 匿名函数和闭包
	fmt.Println("\n--- 匿名函数和闭包 ---")
	// 定义并立即调用匿名函数
	func() {
		fmt.Println("这是一个匿名函数")
	}()

	// 闭包示例
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	fmt.Printf("计数器: %d\n", counter())
	fmt.Printf("计数器: %d\n", counter())
	fmt.Printf("计数器: %d\n", counter())

	// 7. 函数作为值
	fmt.Println("\n--- 函数作为值 ---")
	multiply := func(a, b int) int {
		return a * b
	}

	result = calculate(5, 3, multiply)
	fmt.Printf("5 * 3 = %d\n", result)

	// 使用匿名函数作为回调
	result = calculate(5, 3, func(a, b int) int {
		return a - b
	})
	fmt.Printf("5 - 3 = %d\n", result)

	// 8. defer示例
	fmt.Println("\n--- defer示例 ---")
	deferDemo()

	// 9. panic和recover示例
	fmt.Println("\n--- panic和recover示例 ---")
	panicDemo()
	fmt.Println("程序继续执行") // 这行会执行，因为panic被recover捕获了

	// 10. 方法值和方法表达式
	fmt.Println("\n--- 方法值和方法表达式 ---")
	r := Rectangle{Width: 3, Height: 4}

	// 方法值
	area1 := r.Area
	fmt.Printf("通过方法值调用: %.2f\n", area1())

	// 方法表达式
	area2 := Rectangle.Area
	fmt.Printf("通过方法表达式调用: %.2f\n", area2(r))

	// 11. 函数作为结构体字段
	fmt.Println("\n--- 函数作为结构体字段 ---")
	type Calculator struct {
		Operation func(int, int) int
	}

	calc := Calculator{
		Operation: func(a, b int) int {
			return a + b
		},
	}
	fmt.Printf("计算结果: %d\n", calc.Operation(5, 3))
}
