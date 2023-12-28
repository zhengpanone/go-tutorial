package main

// 数据层
type User struct {
}

// 接口的操作，全部要依赖于数据库（SQL）实现
func (u *User) Name() string {
	return "test user name"
}

func (u *User) NameSet(n string) string {
	return ""
}
