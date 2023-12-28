package main

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestUserName(t *testing.T) {
	want := "test user name"
	user := &User{}

	value, err := UserName(user)
	if err != nil {
		t.Error(err)
	} else if want != value {
		t.Error("UserName error")
	}
}

// 使用mock的单元测试
func TestUserNameMock(t *testing.T) {
	want := "test user name"
	// 1.获取mock控制器
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	// 2.mock 模拟数据
	// 2.1确定需要模拟的接口
	mockUserInterface := NewMockUserInterface(ctl)
	// 2.2模拟数据
	mockUserInterface.EXPECT().Name().Return("test user name")

	// 3.功能调用
	if value, err := UserName(mockUserInterface); err != nil {
		t.Error(err)
	} else if value != want {
		t.Error("UserName error")
	}
}
