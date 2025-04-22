// 变量和基本数据类型示例
//
// 这个示例展示了Go语言中变量声明和基本数据类型的使用，
// 包括数字、字符串、布尔值、复数等类型及其操作。
package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	// 0. 类型大小和范围说明
	fmt.Println("=== 类型大小和范围 ===")
	fmt.Printf("int 大小: %d 字节\n", unsafe.Sizeof(int(0)))
	fmt.Printf("int8 范围: %d 到 %d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("uint8 范围: 0 到 %d\n", math.MaxUint8)
	fmt.Printf("float32 精度: 约6-7位小数\n")
	fmt.Printf("float64 精度: 约15-16位小数\n\n")

	// 0.1 类型零值示例
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool
	fmt.Println("=== 类型零值示例 ===")
	fmt.Printf("int 零值: %d\n", zeroInt)
	fmt.Printf("float64 零值: %f\n", zeroFloat)
	fmt.Printf("string 零值: %q\n", zeroString)
	fmt.Printf("bool 零值: %t\n\n", zeroBool)
	// 1. 变量声明的多种方式
	// 1.1 标准声明
	var name string = "张三"
	var age int = 25

	// 1.2 类型推导（省略类型）
	var country = "中国"

	// 1.3 简短声明（只能在函数内部使用）
	score := 95.5

	// 1.4 多变量声明
	var (
		height  float64 = 175.5
		weight  float64 = 70.2
		isAdult bool    = true
	)

	// 2. 基本数据类型示例
	// 2.1 整数类型
	var (
		intNum    int    = 42         // 根据系统架构可能是32或64位
		int8Num   int8   = 127        // -128 到 127
		int16Num  int16  = 32767      // -32768 到 32767
		int32Num  int32  = 2147483647 // -2147483648 到 2147483647
		int64Num  int64  = 9223372036854775807
		uintNum   uint   = 42    // 无符号整数
		uint8Num  uint8  = 255   // 0 到 255
		uint16Num uint16 = 65535 // 0 到 65535
	)

	// 2.2 浮点数类型
	var (
		float32Num float32 = 3.14159 // 32位浮点数
		float64Num float64 = math.Pi // 64位浮点数
	)

	// 2.3 复数类型
	var (
		complex64Num  complex64  = 3.2 + 12i  // 32位实数和虚数
		complex128Num complex128 = 5.0 + 2.4i // 64位实数和虚数
	)

	// 2.4 布尔类型
	var (
		_ bool = true
		_ bool = false
	)

	// 2.5 字符串类型
	var (
		str1 string = "Hello, 世界" // 可以包含Unicode字符
		str2 string = `这是一个
多行字符串` // 原始字符串
	)

	// 2.6 字符类型（rune，实际上是int32的别名）
	var (
		char1 rune = 'A'
		char2 rune = '中'
	)

	// 3. 常量声明
	const (
		Pi       = 3.14159
		MaxValue = 100
		Prefix   = "GO_"
	)

	// 4. iota 的使用
	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)

	// 5. 类型转换示例
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	// 5.1 更复杂的类型转换
	// 浮点数转整数会截断小数部分
	pi := 3.14159
	piInt := int(pi)
	fmt.Println("\n=== 浮点数转整数 ===")
	fmt.Printf("%.2f -> %d\n", pi, piInt)

	// 5.2 字符串和字节切片的转换
	str := "Hello, 世界"
	bytes := []byte(str)          // 字符串转字节切片
	strFromBytes := string(bytes) // 字节切片转字符串
	fmt.Println("\n=== 字符串和字节切片转换 ===")
	fmt.Printf("原始字符串: %s\n", str)
	fmt.Printf("字节切片: %v\n", bytes)
	fmt.Printf("转换回的字符串: %s\n", strFromBytes)

	// 5.3 类型断言示例
	var any interface{} = "这是一个字符串"
	if str, ok := any.(string); ok {
		fmt.Println("\n=== 类型断言成功 ===")
		fmt.Printf("断言得到的字符串: %s\n", str)
	}

	// 5.4 潜在危险的类型转换
	bigInt := int64(math.MaxInt64)
	smallInt := int32(bigInt) // 可能丢失精度
	fmt.Println("\n=== 潜在危险的类型转换 ===")
	fmt.Printf("int64最大值: %d\n", bigInt)
	fmt.Printf("转换为int32: %d (可能不正确)\n", smallInt)

	// 打印变量值，验证上述代码
	fmt.Println("=== 基本变量示例 ===")
	fmt.Printf("姓名: %s\n", name)
	fmt.Printf("年龄: %d\n", age)
	fmt.Printf("国家: %s\n", country)
	fmt.Printf("分数: %.1f\n", score)
	fmt.Printf("身高: %.1f cm\n", height)
	fmt.Printf("体重: %.1f kg\n", weight)
	fmt.Printf("是否成年: %t\n", isAdult)

	fmt.Println("\n=== 数值类型示例 ===")
	fmt.Printf("int: %d\n", intNum)
	fmt.Printf("int8: %d\n", int8Num)
	fmt.Printf("int16: %d\n", int16Num)
	fmt.Printf("int32: %d\n", int32Num)
	fmt.Printf("int64: %d\n", int64Num)
	fmt.Printf("uint: %d\n", uintNum)
	fmt.Printf("uint8: %d\n", uint8Num)
	fmt.Printf("uint16: %d\n", uint16Num)

	fmt.Println("\n=== 浮点数示例 ===")
	fmt.Printf("float32: %f\n", float32Num)
	fmt.Printf("float64: %f\n", float64Num)

	fmt.Println("\n=== 复数示例 ===")
	fmt.Printf("complex64: %v\n", complex64Num)
	fmt.Printf("complex128: %v\n", complex128Num)

	fmt.Println("\n=== 字符串和字符示例 ===")
	fmt.Printf("字符串1: %s\n", str1)
	fmt.Printf("字符串2: %s\n", str2)
	fmt.Printf("字符1: %c (Unicode: %d)\n", char1, char1)
	fmt.Printf("字符2: %c (Unicode: %d)\n", char2, char2)

	fmt.Println("\n=== 类型转换示例 ===")
	fmt.Printf("int(%d) -> float64(%.1f) -> uint(%d)\n", i, f, u)
}
