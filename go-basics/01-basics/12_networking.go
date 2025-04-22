// 网络编程示例
//
// 这个示例展示了Go语言中网络编程的基本用法，
// 包括HTTP服务器/客户端、TCP/UDP通信、WebSocket等。
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

// 1. 简单的HTTP服务器
func startHTTPServer(wg *sync.WaitGroup) {
	defer wg.Done()

	// 定义处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, 这是首页!")
	})

	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		// 设置响应头
		w.Header().Set("Content-Type", "application/json")

		// 返回JSON响应
		response := map[string]string{"message": "Hello, World!"}
		json.NewEncoder(w).Encode(response)
	})

	// 启动服务器
	fmt.Println("HTTP服务器启动在 :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 2. HTTP客户端
func httpClientExample() {
	fmt.Println("\n=== HTTP客户端示例 ===")

	// 2.1 基本GET请求
	resp, err := http.Get("http://example.com")
	if err != nil {
		fmt.Printf("HTTP GET错误: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应错误: %v\n", err)
		return
	}
	fmt.Printf("响应状态: %s\n", resp.Status)
	fmt.Printf("响应长度: %d bytes\n", len(body))

	// 2.2 自定义HTTP客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 创建请求
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		fmt.Printf("创建请求错误: %v\n", err)
		return
	}

	// 添加请求头
	req.Header.Add("User-Agent", "Go-Tutorial-Client")

	// 发送请求
	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("发送请求错误: %v\n", err)
		return
	}
	defer resp.Body.Close()
}

// 3. TCP服务器
func startTCPServer(wg *sync.WaitGroup) {
	defer wg.Done()

	// 监听端口
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Printf("TCP服务器启动错误: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP服务器启动在 :8081...")

	for {
		// 接受连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("接受连接错误: %v\n", err)
			continue
		}

		// 处理连接
		go handleTCPConnection(conn)
	}
}

// 处理TCP连接
func handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	// 设置超时
	conn.SetDeadline(time.Now().Add(10 * time.Second))

	// 读取数据
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("读取数据错误: %v\n", err)
		return
	}

	// 响应客户端
	message := fmt.Sprintf("收到 %d 字节的数据", n)
	conn.Write([]byte(message))
}

// 4. TCP客户端
func tcpClientExample() {
	fmt.Println("\n=== TCP客户端示例 ===")

	// 连接服务器
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Printf("连接服务器错误: %v\n", err)
		return
	}
	defer conn.Close()

	// 发送数据
	_, err = conn.Write([]byte("Hello, TCP Server!"))
	if err != nil {
		fmt.Printf("发送数据错误: %v\n", err)
		return
	}

	// 读取响应
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("读取响应错误: %v\n", err)
		return
	}

	fmt.Printf("服务器响应: %s\n", buffer[:n])
}

// 5. UDP服务器
func startUDPServer(wg *sync.WaitGroup) {
	defer wg.Done()

	// 解析UDP地址
	addr, err := net.ResolveUDPAddr("udp", ":8082")
	if err != nil {
		fmt.Printf("解析UDP地址错误: %v\n", err)
		return
	}

	// 创建UDP连接
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Printf("UDP服务器启动错误: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP服务器启动在 :8082...")

	buffer := make([]byte, 1024)
	for {
		// 读取数据
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("读取UDP数据错误: %v\n", err)
			continue
		}

		// 响应客户端
		message := fmt.Sprintf("收到来自 %v 的 %d 字节数据", remoteAddr, n)
		conn.WriteToUDP([]byte(message), remoteAddr)
	}
}

// 6. UDP客户端
func udpClientExample() {
	fmt.Println("\n=== UDP客户端示例 ===")

	// 解析服务器地址
	addr, err := net.ResolveUDPAddr("udp", "localhost:8082")
	if err != nil {
		fmt.Printf("解析UDP地址错误: %v\n", err)
		return
	}

	// 创建UDP连接
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Printf("连接UDP服务器错误: %v\n", err)
		return
	}
	defer conn.Close()

	// 发送数据
	_, err = conn.Write([]byte("Hello, UDP Server!"))
	if err != nil {
		fmt.Printf("发送UDP数据错误: %v\n", err)
		return
	}

	// 读取响应
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Printf("读取UDP响应错误: %v\n", err)
		return
	}

	fmt.Printf("服务器响应: %s\n", buffer[:n])
}

// 7. 中间件示例
type Middleware func(http.HandlerFunc) http.HandlerFunc

// 日志中间件
func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		fmt.Printf("[%s] %s %s %v\n",
			r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	}
}

// 认证中间件
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "未授权", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

// 链式调用中间件
func chainMiddleware(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}

func main() {
	var wg sync.WaitGroup

	// 启动HTTP服务器
	wg.Add(1)
	go startHTTPServer(&wg)

	// 启动TCP服务器
	wg.Add(1)
	go startTCPServer(&wg)

	// 启动UDP服务器
	wg.Add(1)
	go startUDPServer(&wg)

	// 等待服务器启动
	time.Sleep(time.Second)

	// 运行客户端示例
	httpClientExample()
	tcpClientExample()
	udpClientExample()

	// 中间件示例
	fmt.Println("\n=== 中间件示例 ===")
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Protected Resource")
	}

	// 应用中间件
	protectedHandler := chainMiddleware(handler,
		loggerMiddleware,
		authMiddleware,
	)

	http.HandleFunc("/protected", protectedHandler)

	fmt.Println("服务器正在运行...")
	wg.Wait()
}

/* 网络编程最佳实践

1. HTTP服务器
   - 始终设置适当的超时
   - 使用合适的内容类型头
   - 实现优雅关闭
   - 使用中间件处理横切关注点

2. HTTP客户端
   - 重用http.Client
   - 设置合理的超时
   - 正确关闭响应体
   - 使用context控制请求

3. TCP/UDP
   - 实现心跳机制
   - 处理连接断开
   - 使用缓冲区池
   - 实现重连机制

4. 错误处理
   - 记录详细错误信息
   - 实现优雅降级
   - 使用超时控制
   - 实现重试机制

5. 性能优化
   - 使用连接池
   - 实现请求限流
   - 使用压缩
   - 实现缓存

6. 安全性
   - 使用TLS/SSL
   - 实现认证授权
   - 防止SQL注入
   - 防止XSS攻击
*/
