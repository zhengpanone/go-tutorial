// 数组和切片示例
//
// 这个示例展示了Go语言中数组和切片的使用，
// 包括创建、初始化、操作和性能注意事项。
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 0. 数组和切片的基本概念
	fmt.Println("=== 数组和切片的基本概念 ===")
	fmt.Println("数组是固定长度的值类型，切片是动态长度的引用类型")
	fmt.Println("切片底层引用一个数组，包含指针、长度和容量")
	// 1. 数组示例
	fmt.Println("=== 数组示例 ===")

	// 1.1 数组声明和初始化
	// 声明定长数组
	var arr1 [5]int                     // 声明一个长度为5的整数数组，元素默认为0
	var arr2 = [5]int{1, 2, 3, 4, 5}    // 声明并初始化
	arr3 := [...]int{1, 2, 3, 4, 5}     // 让编译器计算长度
	arr4 := [5]string{"苹果", "香蕉", "橙子"} // 部分初始化，未初始化的元素为空字符串

	fmt.Println("\n--- 数组的基本操作 ---")
	fmt.Printf("arr1: %v\n", arr1)
	fmt.Printf("arr2: %v\n", arr2)
	fmt.Printf("arr3: %v\n", arr3)
	fmt.Printf("arr4: %v\n", arr4)

	// 1.2 访问和修改数组元素
	arr2[0] = 100 // 修改第一个元素
	fmt.Printf("修改后的arr2[0]: %d\n", arr2[0])
	fmt.Printf("数组长度: %d\n", len(arr2))

	// 1.3 数组遍历
	fmt.Println("\n--- 数组遍历方式 ---")

	// 使用 for 循环遍历
	fmt.Println("使用标准for循环遍历:")
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("arr2[%d] = %d\n", i, arr2[i])
	}

	// 使用 for range 遍历
	fmt.Println("\n使用for range遍历:")
	for index, value := range arr2 {
		fmt.Printf("arr2[%d] = %d\n", index, value)
	}

	// 1.4 多维数组
	fmt.Println("\n--- 多维数组 ---")
	matrix := [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	fmt.Println("二维数组:", matrix)

	// 2. 切片示例
	fmt.Println("\n=== 切片示例 ===")

	// 2.1 创建切片的多种方式
	// 通过make创建切片
	slice1 := make([]int, 5)     // 长度为5的切片
	slice2 := make([]int, 5, 10) // 长度为5，容量为10的切片

	// 通过数组切片操作创建
	arr := [5]int{1, 2, 3, 4, 5}
	slice3 := arr[1:4] // 包含arr的第2个到第4个元素

	// 直接声明切片
	slice4 := []int{1, 2, 3, 4, 5}

	fmt.Println("\n--- 切片的基本信息 ---")
	fmt.Printf("slice1: %v, 长度: %d, 容量: %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2: %v, 长度: %d, 容量: %d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3: %v, 长度: %d, 容量: %d\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("slice4: %v, 长度: %d, 容量: %d\n", slice4, len(slice4), cap(slice4))

	// 2.2 切片的动态增长
	fmt.Println("\n--- 切片的动态增长 ---")
	slice5 := []int{}
	fmt.Printf("初始切片: %v, 长度: %d, 容量: %d\n", slice5, len(slice5), cap(slice5))

	// 使用append添加元素
	slice5 = append(slice5, 1)
	fmt.Printf("添加一个元素: %v, 长度: %d, 容量: %d\n", slice5, len(slice5), cap(slice5))

	// 一次添加多个元素
	slice5 = append(slice5, 2, 3, 4)
	fmt.Printf("添加多个元素: %v, 长度: %d, 容量: %d\n", slice5, len(slice5), cap(slice5))

	// 2.3 切片的常用操作
	fmt.Println("\n--- 切片的常用操作 ---")

	// 2.3.1 切片底层数组共享示例
	fmt.Println("\n--- 切片共享底层数组 ---")
	original := []int{1, 2, 3, 4, 5}
	sharedSlice := original[1:4]
	fmt.Printf("原始切片: %v\n", original)
	fmt.Printf("子切片: %v\n", sharedSlice)

	// 修改子切片会影响原始切片
	sharedSlice[0] = 99
	fmt.Printf("修改子切片后:\n原始切片: %v\n子切片: %v\n", original, sharedSlice)

	// 2.3.2 安全复制切片
	fmt.Println("\n--- 安全复制切片 ---")
	slice6 := make([]int, len(slice4))
	copy(slice6, slice4)
	fmt.Printf("复制后的切片: %v\n", slice6)

	// 修改复制后的切片不会影响原切片
	slice6[0] = 100
	fmt.Printf("修改复制切片后:\n原切片: %v\n复制切片: %v\n", slice4, slice6)

	// 2.3.3 查看切片底层信息
	fmt.Println("\n--- 切片底层信息 ---")
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&slice4))
	fmt.Printf("切片数据指针: %x\n", hdr.Data)
	fmt.Printf("切片长度: %d\n", hdr.Len)
	fmt.Printf("切片容量: %d\n", hdr.Cap)

	// 切片截取
	fmt.Println("\n切片截取操作:")
	fmt.Printf("原始切片: %v\n", slice4)
	fmt.Printf("slice4[1:3]: %v\n", slice4[1:3]) // 包含下标1到2的元素
	fmt.Printf("slice4[:3]: %v\n", slice4[:3])   // 从开始到下标2的元素
	fmt.Printf("slice4[2:]: %v\n", slice4[2:])   // 从下标2到结束的元素

	// 2.4 多维切片
	fmt.Println("\n--- 多维切片 ---")
	// 创建二维切片
	matrix2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("二维切片:", matrix2)

	// 2.5 切片的遍历
	fmt.Println("\n--- 切片的遍历 ---")
	numbers := []int{1, 2, 3, 4, 5}

	// 使用 for range 遍历
	fmt.Println("使用 for range 遍历:")
	for index, value := range numbers {
		fmt.Printf("numbers[%d] = %d\n", index, value)
	}

	// 2.6 nil切片
	fmt.Println("\n--- nil切片 ---")
	var nilSlice []int
	fmt.Printf("nil切片: %v, 长度: %d, 容量: %d, 是否为nil: %t\n",
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)

	// 3. 高级主题
	fmt.Println("\n=== 高级主题 ===")

	// 3.1 性能优化建议
	fmt.Println("\n--- 性能优化建议 ---")

	// 预分配切片容量
	preAllocated := make([]int, 0, 100) // 预分配容量
	for i := 0; i < 100; i++ {
		preAllocated = append(preAllocated, i)
	}
	fmt.Println("预分配容量可减少内存分配次数")

	// 避免大切片保留小引用
	bigData := make([]byte, 1<<20) // 1MB
	smallPart := make([]byte, 10)
	copy(smallPart, bigData[:10]) // 复制需要的数据而不是切片
	fmt.Println("复制需要的数据避免内存泄漏")

	// 3.2 字符串和字节切片
	fmt.Println("\n--- 字符串和字节切片 ---")
	str := "Hello, 世界"
	bytes := []byte(str)          // 字符串转字节切片
	strFromBytes := string(bytes) // 字节切片转字符串
	fmt.Printf("字符串: %s\n", str)
	fmt.Printf("字节切片: %v\n", bytes)
	fmt.Printf("转换回字符串: %s\n", strFromBytes)

	// 3.3 切片陷阱
	fmt.Println("\n--- 切片陷阱 ---")

	// 修改切片可能影响原数组
	original1 := []int{1, 2, 3, 4, 5}
	slice := original1[1:4]
	slice[0] = 99
	fmt.Printf("修改切片影响原数组: %v\n", original1)

	// append可能返回新切片
	s := make([]int, 2, 3)
	s1 := append(s, 1)
	s2 := append(s, 2)
	fmt.Printf("s1: %v\ns2: %v\n", s1, s2)
	fmt.Println("注意: append可能返回新切片，共享底层数组")
}
