package main

import (
	"encoding/json"
	"fmt"
)

// 1. 定义结构体
// Person2 结构体，首字母大写表示可以被其他包访问
type Person2 struct {
	Name   string // 可导出的字段，首字母大写
	Age    int
	gender string // 不可导出的字段，首字母小写
}

// 2. 带标签的结构体
type Student struct {
	Name    string `json:"name" validate:"required"`
	Age     int    `json:"age" validate:"gte=0,lte=130"`
	Grade   string `json:"grade"`
	Address string `json:"address,omitempty"` // omitempty表示如果字段为空则不输出
}

// 3. 嵌套结构体
type Address struct {
	City    string
	Country string
}

type Employee struct {
	Person2         // 匿名嵌套，获得Person的所有字段
	Address Address // 命名嵌套
	Salary  float64
}

// 4. 为Person添加方法
// 值接收者的方法
func (p Person2) Introduce() string {
	return fmt.Sprintf("我是%s，今年%d岁", p.Name, p.Age)
}

// 指针接收者的方法
func (p *Person) SetAge(age int) {
	p.Age = age
}

// 5. 构造函数（通常返回指针）
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	// 1. 结构体的初始化
	fmt.Println("=== 结构体的初始化 ===")

	// 1.1 使用字段名初始化
	p1 := Person{
		Name: "张三",
		Age:  25,
	}
	fmt.Printf("p1: %+v\n", p1)

	// 1.2 按字段顺序初始化
	p2 := Person2{"李四", 30, "男"}
	fmt.Printf("p2: %+v\n", p2)

	// 1.3 使用new关键字（返回指针）
	p3 := new(Person)
	p3.Name = "王五"
	p3.Age = 28
	fmt.Printf("p3: %+v\n", p3)

	// 2. 结构体方法的使用
	fmt.Println("\n=== 结构体方法的使用 ===")

	p1.SetAge(26)
	fmt.Printf("修改年龄后: %+v\n", p1)

	// 3. 结构体标签的使用
	fmt.Println("\n=== 结构体标签的使用 ===")
	student := Student{
		Name:    "赵六",
		Age:     18,
		Grade:   "高三",
		Address: "北京",
	}

	// 将结构体转换为JSON
	jsonData, err := json.Marshal(student)
	if err != nil {
		fmt.Println("JSON转换错误:", err)
	}
	fmt.Printf("JSON格式: %s\n", jsonData)

	// 从JSON转换回结构体
	jsonStr := `{"name":"钱七","age":20,"grade":"大一"}`
	var newStudent Student
	err = json.Unmarshal([]byte(jsonStr), &newStudent)
	if err != nil {
		fmt.Println("JSON解析错误:", err)
	}
	fmt.Printf("解析后的结构体: %+v\n", newStudent)

	// 4. 嵌套结构体的使用
	fmt.Println("\n=== 嵌套结构体的使用 ===")
	emp := Employee{
		Person2: Person2{
			Name: "孙八",
			Age:  35,
		},
		Address: Address{
			City:    "上海",
			Country: "中国",
		},
		Salary: 20000,
	}
	fmt.Printf("员工信息: %+v\n", emp)

	// 访问嵌套字段
	fmt.Printf("姓名: %s\n", emp.Name)         // 直接访问Person的字段
	fmt.Printf("城市: %s\n", emp.Address.City) // 访问Address的字段

	// 5. 结构体指针的使用
	fmt.Println("\n=== 结构体指针的使用 ===")
	pPtr := &Person{
		Name: "周九",
		Age:  40,
	}
	fmt.Printf("通过指针访问: %+v\n", pPtr)

	// 通过指针修改结构体字段
	pPtr.Age = 41 // Go语言会自动解引用，等同于 (*pPtr).Age = 41
	fmt.Printf("修改后: %+v\n", pPtr)

	// 6. 使用构造函数
	fmt.Println("\n=== 使用构造函数 ===")
	p4 := NewPerson("吴十", 45)
	fmt.Printf("通过构造函数创建: %+v\n", p4)

	// 7. 结构体的比较
	fmt.Println("\n=== 结构体的比较 ===")
	p5 := Person{Name: "张三", Age: 25}
	p6 := Person{Name: "张三", Age: 25}
	fmt.Printf("p5 == p6: %v\n", p5 == p6)

	// 8. 结构体切片
	fmt.Println("\n=== 结构体切片 ===")
	people := []Person{
		{Name: "张三", Age: 25},
		{Name: "李四", Age: 30},
		{Name: "王五", Age: 28},
	}
	fmt.Println("人员列表:")
	for _, person := range people {
		fmt.Printf("  %s: %d岁\n", person.Name, person.Age)
	}

	// 9. 匿名结构体
	fmt.Println("\n=== 匿名结构体 ===")
	anonymous := struct {
		Name string
		Age  int
	}{
		Name: "匿名",
		Age:  20,
	}
	fmt.Printf("匿名结构体: %+v\n", anonymous)
}
