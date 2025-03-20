package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// 模拟服务端处理函数，返回模拟的事件流
func handleWatch(w http.ResponseWriter, r *http.Request) {
	// 设置响应头，让客户端知道这是一个事件流
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// 模拟不断发送数据变化
	for {
		// 模拟生成一个事件数据
		eventData := fmt.Sprintf("data: %s\n\n", time.Now().Format(time.RFC3339))

		// 将数据写入响应流
		if _, err := io.WriteString(w, eventData); err != nil {
			log.Printf("Error writing to client: %v", err)
			return
		}

		// 刷新缓冲区，确保数据立即发送给客户端
		if flusher, ok := w.(http.Flusher); ok {
			flusher.Flush()
		}

		// 模拟一定时间间隔发送数据变化
		time.Sleep(2 * time.Second)
	}
}

func main() {
	// 启动模拟服务端
	http.HandleFunc("/watch", handleWatch)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
