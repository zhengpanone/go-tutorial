package main

import (
	"fmt"
	"math"
	"testing"
)

// 1. 基本单元测试
func TestCalculatorAdd(t *testing.T) {
	// 创建测试用例结构
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"正数相加", 2, 3, 5},
		{"负数相加", -2, -3, -5},
		{"零相加", 0, 0, 0},
		{"正负数相加", 2, -3, -1},
	}

	// 运行所有测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := &Calculator{}
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%v, %v) = %v; 期望值 %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// 2. 测试除法错误处理
func TestCalculatorDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        float64
		expected    float64
		expectError bool
	}{
		{"正常除法", 6, 2, 3, false},
		{"除以零", 1, 0, 0, true},
		{"负数除法", -6, 2, -3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := &Calculator{}
			result, err := calc.Divide(tt.a, tt.b)

			// 检查错误情况
			if tt.expectError && err == nil {
				t.Error("期望得到一个错误，但没有得到")
			}
			if !tt.expectError && err != nil {
				t.Errorf("不期望得到错误，但得到了: %v", err)
			}

			// 如果不期望有错误，检查结果
			if !tt.expectError && result != tt.expected {
				t.Errorf("Divide(%v, %v) = %v; 期望值 %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// 3. 基准测试
func BenchmarkCalculatorAdd(b *testing.B) {
	calc := &Calculator{}
	for i := 0; i < b.N; i++ {
		calc.Add(2, 3)
	}
}

func BenchmarkCalculatorPower(b *testing.B) {
	calc := &Calculator{}
	for i := 0; i < b.N; i++ {
		calc.Power(2, 3)
	}
}

// 4. 示例测试
func ExampleCalculator_Add() {
	calc := &Calculator{}
	result := calc.Add(2, 3)
	fmt.Printf("2 + 3 = %.0f\n", result)
	// Output: 2 + 3 = 5
}

func ExampleCalculator_Multiply() {
	calc := &Calculator{}
	result := calc.Multiply(4, 5)
	fmt.Printf("4 * 5 = %.0f\n", result)
	// Output: 4 * 5 = 20
}

// 5. 测试辅助函数
func assertFloat64Equal(t *testing.T, expected, actual float64, epsilon float64) {
	t.Helper() // 标记这是一个辅助函数
	if math.Abs(expected-actual) > epsilon {
		t.Errorf("期望值 %v，实际值 %v，差异超过 %v", expected, actual, epsilon)
	}
}

// 6. 使用测试辅助函数的测试
func TestCalculatorSquareRoot(t *testing.T) {
	tests := []struct {
		name        string
		input       float64
		expected    float64
		expectError bool
	}{
		{"正数平方根", 16, 4, false},
		{"零的平方根", 0, 0, false},
		{"负数平方根", -1, 0, true},
	}

	calc := &Calculator{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.SquareRoot(tt.input)

			if tt.expectError && err == nil {
				t.Error("期望得到一个错误，但没有得到")
				return
			}
			if !tt.expectError && err != nil {
				t.Errorf("不期望得到错误，但得到了: %v", err)
				return
			}

			if !tt.expectError {
				assertFloat64Equal(t, tt.expected, result, 0.0001)
			}
		})
	}
}

// 7. 表格驱动测试
func TestCalculatorOperations(t *testing.T) {
	// 定义测试表格
	tests := []struct {
		name     string
		op       string
		a, b     float64
		expected float64
		hasError bool
	}{
		{"加法-正数", "add", 2, 3, 5, false},
		{"减法-正数", "subtract", 5, 3, 2, false},
		{"乘法-正数", "multiply", 4, 3, 12, false},
		{"除法-正数", "divide", 6, 2, 3, false},
		{"除法-除以零", "divide", 6, 0, 0, true},
	}

	// 执行测试表格
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := &Calculator{}
			var result float64
			var err error

			switch tt.op {
			case "add":
				result = calc.Add(tt.a, tt.b)
			case "subtract":
				result = calc.Subtract(tt.a, tt.b)
			case "multiply":
				result = calc.Multiply(tt.a, tt.b)
			case "divide":
				result, err = calc.Divide(tt.a, tt.b)
			}

			// 检查错误
			if tt.hasError && err == nil {
				t.Error("期望得到一个错误，但没有得到")
				return
			}
			if !tt.hasError && err != nil {
				t.Errorf("不期望得到错误，但得到了: %v", err)
				return
			}

			// 检查结果
			if !tt.hasError && result != tt.expected {
				t.Errorf("%s(%v, %v) = %v; 期望值 %v",
					tt.op, tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// 8. 测试Setup和Teardown
func setupTest(t *testing.T) (*Calculator, func()) {
	t.Helper()
	calc := &Calculator{}
	// 返回清理函数
	return calc, func() {
		// 这里可以进行清理工作
		calc.lastResult = 0
	}
}

func TestCalculatorWithSetup(t *testing.T) {
	calc, teardown := setupTest(t)
	defer teardown()

	result := calc.Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %v; 期望值 5", result)
	}
}

// 9. 并行测试
func TestCalculatorParallel(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"test1", 2, 3, 5},
		{"test2", 4, 5, 9},
		{"test3", 6, 7, 13},
	}

	for _, tt := range tests {
		tt := tt // 捕获范围变量
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // 标记测试可以并行运行
			calc := &Calculator{}
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%v, %v) = %v; 期望值 %v",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

/*
运行测试的命令：

1. 运行所有测试
   go test

2. 运行特定测试
   go test -run TestCalculatorAdd

3. 运行基准测试
   go test -bench=.

4. 生成测试覆盖率报告
   go test -cover
   go test -coverprofile=coverage.out
   go tool cover -html=coverage.out

5. 运行示例测试
   go test -v

6. 并行测试
   go test -parallel 4

7. 详细输出
   go test -v

8. 运行基准测试并输出内存分配统计
   go test -bench=. -benchmem
*/
