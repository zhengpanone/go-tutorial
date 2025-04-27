package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}

func main() {
	e, err := casbin.NewEnforcer("./model.pml", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}
	check(e, "zhangsan", "/index", "GET")
	check(e, "zhangsan", "/home", "GET")
	check(e, "zhangsan", "/users", "POST")
	check(e, "wangwu", "/users", "POST")

	e.AddPolicy("wangwu", "/users", "POST")
	e.SavePolicy()
	check(e, "wangwu", "/users", "POST")
	e.RemovePolicy("wangwu", "/users", "POST")
	e.SavePolicy()
	check(e, "wangwu", "/users", "POST")

	// 添加角色-权限对应关系
	e.AddPolicy("kuaiji", "/roleusers", "GET")
	e.SavePolicy()

	// 添加用户-角色对应关系
	e.AddRoleForUser("zhangsan", "kuaiji")
	e.SavePolicy()
	check(e, "zhangsan", "/roleusers", "GET")

	// 删除用户-角色对应关系
	e.RemoveGroupingPolicy("zhangsan", "kuaiji")
	e.SavePolicy()
	// 角色-权限对应关系
	e.RemovePolicy("kuaiji", "/roleusers", "GET")
	e.SavePolicy()
	check(e, "zhangsan", "/roleusers", "GET")
}
