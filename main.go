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
	http.HandleFunc("/bundesliga", handlers.BundesLigaHandler)
	http.HandleFunc("/bundesligaguess", handlers.BundesLigaGuessHandler)

	// 启动服务器
	log.Println("Starting server on :8888...")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
