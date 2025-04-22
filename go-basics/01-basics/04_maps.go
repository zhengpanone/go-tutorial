// 映射(map)示例
//
// 这个示例展示了Go语言中map的使用，
// 包括创建、初始化、操作和性能注意事项。
package main

import (
	"fmt"
	"sort"
	"sync"
)

// map的底层原理说明
/*
Go中的map是哈希表的实现，具有以下特点：
1. 引用类型，零值为nil
2. 键必须是可比较的类型
3. 非线程安全，并发访问需要加锁
4. 遍历顺序随机
5. 自动扩容，但不会自动缩容
*/

// 声明一个函数，展示map作为参数的使用
func printMap(m map[string]int) {
	for key, value := range m {
		fmt.Printf("键:%s, 值:%d\n", key, value)
	}
}

// 修改map中的值（map是引用类型，函数内的修改会影响原map）
func modifyMap(m map[string]int) {
	m["修改的值"] = 999
}

func main() {
	// 1. map的声明和初始化
	fmt.Println("=== Map的声明和初始化 ===")

	// 1.1 使用make声明
	scores := make(map[string]int)
	fmt.Printf("空map: %v\n", scores)

	// 1.2 直接声明并初始化
	ages := map[string]int{
		"张三": 25,
		"李四": 30,
		"王五": 28,
	}
	fmt.Printf("初始化的map: %v\n", ages)

	// 1.3 声明nil map
	var nilMap map[string]int
	fmt.Printf("nil map: %v, 是否为nil: %t\n", nilMap, nilMap == nil)

	// 2. map的基本操作
	fmt.Println("\n=== Map的基本操作 ===")

	// 2.1 添加和修改元素
	scores["张三"] = 95
	scores["李四"] = 88
	scores["王五"] = 78
	fmt.Printf("添加元素后的map: %v\n", scores)

	// 修改元素
	scores["张三"] = 97
	fmt.Printf("修改后的分数: %v\n", scores)

	// 2.2 获取元素
	score := scores["张三"]
	fmt.Printf("张三的分数: %d\n", score)

	// 2.3 检查键是否存在
	score, exists := scores["赵六"]
	if exists {
		fmt.Printf("赵六的分数: %d\n", score)
	} else {
		fmt.Println("赵六的分数不存在")
	}

	// 2.4 删除元素
	delete(scores, "王五")
	fmt.Printf("删除后的map: %v\n", scores)

	// 3. map的遍历
	fmt.Println("\n=== Map的遍历 ===")

	// 3.1 遍历所有的键值对
	fmt.Println("遍历所有键值对:")
	for name, score := range scores {
		fmt.Printf("姓名:%s, 分数:%d\n", name, score)
	}

	// 3.2 只遍历键
	fmt.Println("\n只遍历键:")
	for name := range scores {
		fmt.Printf("姓名: %s\n", name)
	}

	// 3.3 只遍历值
	fmt.Println("\n只遍历值:")
	for _, score := range scores {
		fmt.Printf("分数: %d\n", score)
	}

	// 4. 嵌套map
	fmt.Println("\n=== 嵌套Map ===")

	// 4.1 创建嵌套map
	students := map[string]map[string]int{
		"张三": {
			"语文": 95,
			"数学": 88,
			"英语": 92,
		},
		"李四": {
			"语文": 85,
			"数学": 94,
			"英语": 87,
		},
	}

	// 4.2 访问嵌套map
	fmt.Println("\n访问嵌套map:")
	for name, subjects := range students {
		fmt.Printf("%s的成绩:\n", name)
		for subject, score := range subjects {
			fmt.Printf("  %s: %d\n", subject, score)
		}
	}

	// 4.3 修改嵌套map
	students["张三"]["语文"] = 98
	fmt.Printf("\n修改后张三的语文成绩: %d\n", students["张三"]["语文"])

	// 5. map作为函数参数
	fmt.Println("\n=== Map作为函数参数 ===")
	testMap := map[string]int{
		"测试1": 100,
		"测试2": 200,
	}

	fmt.Println("原始map:")
	printMap(testMap)

	modifyMap(testMap)
	fmt.Println("\n修改后的map:")
	printMap(testMap)

	// 6. map的常见使用场景
	fmt.Println("\n=== Map的常见使用场景 ===")

	// 6.1 用作计数器
	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	counter := make(map[string]int)

	for _, word := range words {
		counter[word]++
	}
	fmt.Println("单词计数:", counter)

	// 6.2 用作缓存
	cache := make(map[string]string)
	cache["key1"] = "value1"
	cache["key2"] = "value2"

	// 检查缓存中是否存在某个值
	if value, ok := cache["key1"]; ok {
		fmt.Printf("缓存中找到值: %s\n", value)
	}

	// 7. map的注意事项
	fmt.Println("\n=== Map的注意事项 ===")

	// 7.1 map的容量
	m := make(map[string]int, 10) // 创建初始容量为10的map
	fmt.Printf("创建的map: %v\n", m)

	// 7.2 map的键类型
	// map的键必须是可比较的类型
	// 可以使用的类型：布尔、数字、字符串、指针、通道、接口类型、结构体（如果它们的所有字段都是可比较的）
	// 不能使用的类型：切片、映射、函数

	// 正确的键类型示例
	map1 := map[int]string{1: "一"}
	map2 := map[string]int{"一": 1}
	map3 := map[struct{ name string }]int{struct{ name string }{"张三"}: 1}

	fmt.Printf("数字作为键: %v\n", map1)
	fmt.Printf("字符串作为键: %v\n", map2)
	fmt.Printf("结构体作为键: %v\n", map3)

	// 8. 高级主题
	fmt.Println("\n=== 高级主题 ===")

	// 8.1 并发安全的sync.Map
	fmt.Println("\n--- 并发安全的sync.Map ---")
	var sm sync.Map
	sm.Store("key1", "value1")
	sm.Store("key2", "value2")
	if value, ok := sm.Load("key1"); ok {
		fmt.Printf("sync.Map中key1的值: %v\n", value)
	}

	// 8.2 map的排序遍历
	fmt.Println("\n--- map的排序遍历 ---")
	unsorted := map[string]int{
		"张三": 95,
		"李四": 88,
		"王五": 78,
	}

	// 获取排序后的键
	keys := make([]string, 0, len(unsorted))
	for k := range unsorted {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 按排序后的键遍历
	fmt.Println("按姓名排序:")
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, unsorted[k])
	}

	// 8.3 性能优化建议
	fmt.Println("\n--- 性能优化建议 ---")
	fmt.Println("1. 预分配足够容量减少扩容")
	fmt.Println("2. 小map比大map性能更好")
	fmt.Println("3. 频繁访问的键使用短字符串")
	fmt.Println("4. 并发访问使用sync.Map或加锁")

	// 8.4 零值处理
	fmt.Println("\n--- 零值处理 ---")
	var zeroMap map[string]int
	if zeroMap == nil {
		fmt.Println("零值map是nil")
	}
	// 必须初始化后才能使用
	zeroMap = make(map[string]int)
	zeroMap["有效"] = 1
	fmt.Printf("初始化后的零值map: %v\n", zeroMap)

	// 8.5 map的底层实现
	fmt.Println("\n--- map底层实现 ---")
	fmt.Println("Go的map使用哈希表实现:")
	fmt.Println("- 平均O(1)时间复杂度")
	fmt.Println("- 自动处理哈希冲突")
	fmt.Println("- 扩容时渐进式rehash")
	fmt.Println("- 非线程安全设计")
}
