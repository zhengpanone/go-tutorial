// 数据库操作示例
//
// 这个示例展示了Go语言中数据库操作的基本用法，
// 包括SQL数据库、ORM、NoSQL等。
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // MySQL驱动
	"gorm.io/driver/mysql"             // GORM MySQL驱动
	"gorm.io/gorm"                     // GORM ORM
)

// 1. 用户模型
type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Email     string    `gorm:"size:255;uniqueIndex" json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 2. 数据库配置
const (
	mysqlDSN = "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	redisDSN = "localhost:6379"
)

// 3. 标准库SQL操作示例
func sqlExample() {
	fmt.Println("\n=== 标准库SQL示例 ===")

	// 连接数据库
	db, err := sql.Open("mysql", mysqlDSN)
	if err != nil {
		log.Printf("连接数据库错误: %v\n", err)
		return
	}
	defer db.Close()

	// 设置连接池
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	// 创建表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE,
			age INT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Printf("创建表错误: %v\n", err)
		return
	}

	// 插入数据
	result, err := db.Exec(
		"INSERT INTO users (name, email, age) VALUES (?, ?, ?)",
		"张三", "zhangsan@example.com", 25,
	)
	if err != nil {
		log.Printf("插入数据错误: %v\n", err)
		return
	}

	id, _ := result.LastInsertId()
	fmt.Printf("插入用户ID: %d\n", id)

	// 查询单行
	var user User
	err = db.QueryRow("SELECT id, name, email, age FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err != nil {
		log.Printf("查询数据错误: %v\n", err)
		return
	}
	fmt.Printf("查询到的用户: %+v\n", user)

	// 事务示例
	tx, err := db.Begin()
	if err != nil {
		log.Printf("开始事务错误: %v\n", err)
		return
	}

	// 在事务中执行操作
	_, err = tx.Exec("UPDATE users SET age = ? WHERE id = ?", 26, id)
	if err != nil {
		tx.Rollback()
		log.Printf("更新数据错误: %v\n", err)
		return
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		log.Printf("提交事务错误: %v\n", err)
		return
	}
}

// 4. GORM示例
func gormExample() {
	fmt.Println("\n=== GORM示例 ===")

	// 连接数据库
	db, err := gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{})
	if err != nil {
		log.Printf("连接数据库错误: %v\n", err)
		return
	}

	// 自动迁移
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Printf("自动迁移错误: %v\n", err)
		return
	}

	// 创建用户
	user := User{
		Name:  "李四",
		Email: "lisi@example.com",
		Age:   30,
	}

	result := db.Create(&user)
	if result.Error != nil {
		log.Printf("创建用户错误: %v\n", result.Error)
		return
	}
	fmt.Printf("创建用户成功，ID: %d\n", user.ID)

	// 查询用户
	var users []User
	result = db.Where("age > ?", 20).Find(&users)
	if result.Error != nil {
		log.Printf("查询用户错误: %v\n", result.Error)
		return
	}
	fmt.Printf("查询到 %d 个用户\n", len(users))

	// 更新用户
	result = db.Model(&user).Updates(User{Age: 31})
	if result.Error != nil {
		log.Printf("更新用户错误: %v\n", result.Error)
		return
	}

	// 删除用户
	result = db.Delete(&user)
	if result.Error != nil {
		log.Printf("删除用户错误: %v\n", result.Error)
		return
	}
}

// 5. Redis示例
func redisExample() {
	fmt.Println("\n=== Redis示例 ===")

	// 创建Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisDSN,
		Password: "", // 无密码
		DB:       0,  // 默认DB
	})

	ctx := rdb.Context()

	// 设置键值
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Printf("设置键值错误: %v\n", err)
		return
	}

	// 获取值
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		log.Printf("获取值错误: %v\n", err)
		return
	}
	fmt.Printf("key = %s\n", val)

	// 设置带过期时间的键值
	err = rdb.Set(ctx, "temp", "temporary value", time.Second*30).Err()
	if err != nil {
		log.Printf("设置临时键值错误: %v\n", err)
		return
	}

	// Hash操作
	user := User{
		Name:  "王五",
		Email: "wangwu@example.com",
		Age:   35,
	}

	// 将结构体转换为map
	userJSON, _ := json.Marshal(user)
	err = rdb.HSet(ctx, "user:1", "data", userJSON).Err()
	if err != nil {
		log.Printf("设置Hash错误: %v\n", err)
		return
	}

	// 获取Hash
	userdata, err := rdb.HGet(ctx, "user:1", "data").Result()
	if err != nil {
		log.Printf("获取Hash错误: %v\n", err)
		return
	}
	fmt.Printf("用户数据: %s\n", userdata)

	// 列表操作
	rdb.RPush(ctx, "list", "first")
	rdb.RPush(ctx, "list", "second")
	rdb.RPush(ctx, "list", "third")

	// 获取列表
	list, err := rdb.LRange(ctx, "list", 0, -1).Result()
	if err != nil {
		log.Printf("获取列表错误: %v\n", err)
		return
	}
	fmt.Printf("列表内容: %v\n", list)
}

func main() {
	// 运行SQL示例
	sqlExample()

	// 运行GORM示例
	gormExample()

	// 运行Redis示例
	redisExample()
}

/* 数据库操作最佳实践

1. 连接管理
   - 正确配置连接池
   - 使用合适的最大连接数
   - 设置连接超时和生命周期
   - 实现健康检查

2. 查询优化
   - 使用预处理语句
   - 避免N+1查询问题
   - 合理使用索引
   - 分页查询大数据集

3. 事务处理
   - 保持事务简短
   - 避免长事务
   - 正确处理回滚
   - 使用事务隔离级别

4. 错误处理
   - 优雅处理连接错误
   - 实现重试机制
   - 记录详细错误信息
   - 处理死锁情况

5. 性能优化
   - 使用批量操作
   - 实现缓存策略
   - 优化查询语句
   - 监控数据库性能

6. 安全性
   - 防止SQL注入
   - 加密敏感数据
   - 实现访问控制
   - 定期备份数据

7. ORM使用建议
   - 合理使用预加载
   - 避免反射开销
   - 使用原生SQL处理复杂查询
   - 正确处理关联关系
*/
