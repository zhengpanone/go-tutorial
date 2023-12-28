package main

// 用户模型操作接口
type UserInterface interface {
	// 获取名字
	Name() string
	// 设置名字
	NameSet(string) string
}
