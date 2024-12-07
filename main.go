package main

import (
	"gy-go-aes-server/handlers"
	"log"
	"net/http"
)

func main() {
	// 路由定义
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/encrypt", handlers.EncryptHandler)
	http.HandleFunc("/decrypt", handlers.DecryptHandler)

	// 启动服务器
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
