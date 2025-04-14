package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "1234567890"

	fmt.Printf("第一次加密后的密码：%s\n", encryptPassword(password))
	fmt.Printf("第二次加密后的密码：%s\n", encryptPassword(password))
	cost, _ := bcrypt.Cost([]byte(encryptPassword(password)))
	fmt.Printf("Cost次数：%v\n", cost)

	fmt.Printf("密码对比结果：%v\n", comparePasssword(password, encryptPassword(password)))
	fmt.Printf("密码对比结果：%v\n", comparePasssword("123", encryptPassword(password)))
}

func encryptPassword(password string) string {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashPassword)
}

func comparePasssword(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
