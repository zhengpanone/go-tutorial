package main

import (
	"errors"
	"fmt"
	"math"
)

// Calculator 结构体用于演示测试
type Calculator struct {
	lastResult float64
}

// Add 执行加法运算
func (c *Calculator) Add(a, b float64) float64 {
	c.lastResult = a + b
	return c.lastResult
}

// Subtract 执行减法运算
func (c *Calculator) Subtract(a, b float64) float64 {
	c.lastResult = a - b
	return c.lastResult
}

// Multiply 执行乘法运算
func (c *Calculator) Multiply(a, b float64) float64 {
	c.lastResult = a * b
	return c.lastResult
}

// Divide 执行除法运算
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	c.lastResult = a / b
	return c.lastResult, nil
}

// GetLastResult 返回最后一次计算的结果
func (c *Calculator) GetLastResult() float64 {
	return c.lastResult
}

// Square 计算平方
func (c *Calculator) Square(n float64) float64 {
	c.lastResult = n * n
	return c.lastResult
}

// SquareRoot 计算平方根
func (c *Calculator) SquareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("不能对负数求平方根")
	}
	c.lastResult = math.Sqrt(n)
	return c.lastResult, nil
}

// Power 计算幂
func (c *Calculator) Power(base, exponent float64) float64 {
	c.lastResult = math.Pow(base, exponent)
	return c.lastResult
}

// Factorial 计算阶乘
func (c *Calculator) Factorial(n int) (uint64, error) {
	if n < 0 {
		return 0, errors.New("不能计算负数的阶乘")
	}
	if n > 20 {
		return 0, errors.New("数字太大，可能会溢出")
	}
	var result uint64 = 1
	for i := 2; i <= n; i++ {
		result *= uint64(i)
	}
	c.lastResult = float64(result)
	return result, nil
}

func main() {
	// 这个文件主要用于测试，但我们也可以展示一些基本用法
	fmt.Println("=== 计算器测试示例 ===")

	calc := &Calculator{}

	// 基本运算示例
	fmt.Printf("加法: %.2f\n", calc.Add(5, 3))
	fmt.Printf("减法: %.2f\n", calc.Subtract(10, 4))
	fmt.Printf("乘法: %.2f\n", calc.Multiply(6, 7))

	if result, err := calc.Divide(15, 3); err == nil {
		fmt.Printf("除法: %.2f\n", result)
	}

	if result, err := calc.SquareRoot(16); err == nil {
		fmt.Printf("平方根: %.2f\n", result)
	}

	fmt.Printf("幂运算: %.2f\n", calc.Power(2, 3))

	if result, err := calc.Factorial(5); err == nil {
		fmt.Printf("阶乘: %d\n", result)
	}
}
