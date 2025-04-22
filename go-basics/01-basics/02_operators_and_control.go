// 运算符和控制流程示例
//
// 这个示例展示了Go语言中的各种运算符和控制流程语句，
// 包括算术、比较、逻辑运算符以及if/switch/for等控制语句。
package main

import (
	"fmt"
)

func main() {
	// 1. 运算符示例
	fmt.Println("=== 运算符示例 ===")

	// 1.1 算术运算符
	fmt.Println("\n--- 算术运算符 ---")
	a, b := 10, 3
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("加法: a + b = %d\n", a+b)
	fmt.Printf("减法: a - b = %d\n", a-b)
	fmt.Printf("乘法: a * b = %d\n", a*b)
	fmt.Printf("除法: a / b = %d\n", a/b)
	fmt.Printf("取余: a %% b = %d\n", a%b)

	// 自增和自减
	x := 5
	x++ // 自增
	fmt.Printf("x++ 后: %d\n", x)
	x-- // 自减
	fmt.Printf("x-- 后: %d\n", x)

	// 1.2 比较运算符
	fmt.Println("\n--- 比较运算符 ---")
	fmt.Printf("等于: a == b: %t\n", a == b)
	fmt.Printf("不等于: a != b: %t\n", a != b)
	fmt.Printf("大于: a > b: %t\n", a > b)
	fmt.Printf("小于: a < b: %t\n", a < b)
	fmt.Printf("大于等于: a >= b: %t\n", a >= b)
	fmt.Printf("小于等于: a <= b: %t\n", a <= b)

	// 1.3 逻辑运算符
	fmt.Println("\n--- 逻辑运算符 ---")
	t, f := true, false
	fmt.Printf("与: t && f = %t\n", t && f)
	fmt.Printf("或: t || f = %t\n", t || f)
	fmt.Printf("非: !t = %t\n", !t)

	// 1.4 位运算符
	fmt.Println("\n--- 位运算符 ---")
	x, y := 5, 3 // 二进制: 0101, 0011
	fmt.Printf("按位与: x & y = %04b (%d)\n", x&y, x&y)
	fmt.Printf("按位或: x | y = %04b (%d)\n", x|y, x|y)
	fmt.Printf("按位异或: x ^ y = %04b (%d)\n", x^y, x^y)
	fmt.Printf("左移: x << 1 = %04b (%d)\n", x<<1, x<<1)
	fmt.Printf("右移: x >> 1 = %04b (%d)\n", x>>1, x>>1)

	// 2. 控制流程
	fmt.Println("\n=== 控制流程示例 ===")

	// 2.1 if-else 语句
	fmt.Println("\n--- if-else 示例 ---")
	score := 85
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	// if 初始化语句
	if num := 20; num > 0 {
		fmt.Printf("%d 是正数\n", num)
	}

	// 2.2 switch 语句
	fmt.Println("\n--- switch 示例 ---")
	day := 3
	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	default:
		fmt.Println("周末")
	}

	// switch 使用表达式
	score = 85
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	// 2.3 for 循环
	fmt.Println("\n--- for 循环示例 ---")

	// 标准 for 循环
	fmt.Println("标准 for 循环:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 类似 while 的 for 循环
	fmt.Println("类似 while 的 for 循环:")
	n := 0
	for n < 5 {
		fmt.Printf("%d ", n)
		n++
	}
	fmt.Println()

	// 无限循环（使用 break）
	fmt.Println("break 示例:")
	sum := 0
	for {
		sum++
		if sum > 5 {
			break
		}
		fmt.Printf("%d ", sum)
	}
	fmt.Println()

	// continue 示例
	fmt.Println("continue 示例:")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // 跳过 2
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// for range 循环
	fmt.Println("\n--- for range 示例 ---")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("遍历切片:")
	for index, value := range numbers {
		fmt.Printf("索引:%d 值:%d\n", index, value)
	}

	str := "Hello, 世界"
	fmt.Println("\n遍历字符串:")
	for index, char := range str {
		fmt.Printf("位置:%d 字符:%c\n", index, char)
	}

	// 2.4 高级控制流示例
	fmt.Println("\n=== 高级控制流示例 ===")

	// 带标签的break/continue
	fmt.Println("\n--- 带标签的break/continue ---")
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("跳过i=1,j=1的情况")
				continue OuterLoop
			}
			fmt.Printf("i=%d,j=%d ", i, j)
		}
	}
	fmt.Println()

	// goto示例(不推荐但需要了解)
	fmt.Println("\n--- goto示例 ---")
	i := 0
Start:
	if i < 5 {
		fmt.Printf("%d ", i)
		i++
		goto Start
	}
	fmt.Println("\n警告: goto语句应谨慎使用，通常有更好的替代方案")

	// 更复杂的switch case
	fmt.Println("\n--- 更复杂的switch case ---")
	value := 42
	switch {
	case value < 0:
		fmt.Println("负数")
	case value == 0:
		fmt.Println("零")
	case value > 0 && value < 100:
		fmt.Println("1到99之间的正数")
		fallthrough // 继续执行下一个case
	case value >= 100:
		fmt.Println("大于等于100的数")
	default:
		fmt.Println("不应该到达这里")
	}

	// 带条件的for range
	fmt.Println("\n--- 带条件的for range ---")
	data := []int{1, 2, 3, 4, 5}
	for i, v := range data {
		if i == 0 {
			continue // 跳过第一个元素
		}
		if v > 3 {
			break // 遇到大于3的元素就停止
		}
		fmt.Printf("data[%d] = %d\n", i, v)
	}
}
