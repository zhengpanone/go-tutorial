package main

//import "strings"

// 业务操作层代码模拟

func UserName(user UserInterface) (string, error) {
	// 数据层获取用户名字
	name := user.Name()
	// 后续操作
	// name = strings.ToUpper(name)
	// 数据反馈
	return name, nil
}
