package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 包级变量
var (
	debugMode = false
	logger    = log.New(os.Stdout, "[包管理示例] ", log.Ldate|log.Ltime)
)

// 包的初始化函数
func init() {
	// init函数在main函数之前自动执行
	logger.Println("包初始化...")
	rand.Seed(time.Now().UnixNano())
}

// 示例结构体
type Config struct {
	Name     string   `json:"name"`
	Version  string   `json:"version"`
	Features []string `json:"features"`
}

// 导出的函数（首字母大写）
func PrintConfig(cfg Config) {
	fmt.Printf("配置名称: %s\n", cfg.Name)
	fmt.Printf("版本: %s\n", cfg.Version)
	fmt.Printf("功能列表: %v\n", cfg.Features)
}

// 未导出的函数（首字母小写）
func validateConfig(cfg Config) error {
	if cfg.Name == "" {
		return fmt.Errorf("配置名称不能为空")
	}
	if cfg.Version == "" {
		return fmt.Errorf("版本号不能为空")
	}
	return nil
}

// 演示标准库的使用
func demonstrateStdLib() {
	// 1. strings包
	fmt.Println("\n--- strings包示例 ---")
	text := "  Hello, Go!  "
	fmt.Printf("原始文本: %q\n", text)
	fmt.Printf("去除空格: %q\n", strings.TrimSpace(text))
	fmt.Printf("转换为小写: %q\n", strings.ToLower(text))

	// 2. filepath包
	fmt.Println("\n--- filepath包示例 ---")
	path := "dir/subdir/file.txt"
	fmt.Printf("路径分析:\n")
	fmt.Printf("目录: %s\n", filepath.Dir(path))
	fmt.Printf("文件名: %s\n", filepath.Base(path))
	fmt.Printf("扩展名: %s\n", filepath.Ext(path))

	// 3. json包
	fmt.Println("\n--- json包示例 ---")
	cfg := Config{
		Name:     "MyApp",
		Version:  "1.0.0",
		Features: []string{"特性1", "特性2", "特性3"},
	}

	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		logger.Printf("JSON编码错误: %v\n", err)
		return
	}
	fmt.Printf("JSON格式化输出:\n%s\n", string(jsonData))

	// 4. http包
	fmt.Println("\n--- http包示例 ---")
	// 创建一个简单的HTTP处理函数
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go!")
	}
	http.HandleFunc("/hello", handler)
	// 在实际应用中会启动服务器
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

// 演示错误处理最佳实践
type AppError struct {
	Err     error
	Message string
	Code    int
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s (代码: %d): %v", e.Message, e.Code, e.Err)
}

// 演示包的文档注释
/*
Package main 演示了Go语言的包管理和模块化特性。

主要功能：
  - 包的基本使用
  - 模块依赖管理
  - 标准库的使用示例
  - 错误处理最佳实践
  - 包的文档注释规范

使用示例：
    go run main.go

更多信息请参考：https://golang.org/doc/
*/

func main() {
	fmt.Println("=== 包管理和模块化示例 ===")

	// 1. 配置对象的使用
	cfg := Config{
		Name:     "示例应用",
		Version:  "1.0.0",
		Features: []string{"包管理", "模块化", "文档化"},
	}

	// 2. 验证配置
	if err := validateConfig(cfg); err != nil {
		logger.Printf("配置验证失败: %v\n", err)
		return
	}

	// 3. 打印配置
	PrintConfig(cfg)

	// 4. 演示标准库的使用
	demonstrateStdLib()

	// 5. 包级变量的使用
	if debugMode {
		logger.Println("调试模式已启用")
	}

	// 6. 错误处理示例
	if err := doSomething2(); err != nil {
		switch e := err.(type) {
		case *AppError:
			logger.Printf("应用错误: %v\n", e)
		default:
			logger.Printf("未知错误: %v\n", err)
		}
	}
}

// 示例函数
func doSomething2() error {
	// 模拟一个错误
	return &AppError{
		Err:     fmt.Errorf("操作失败"),
		Message: "处理请求时发生错误",
		Code:    500,
	}
}
