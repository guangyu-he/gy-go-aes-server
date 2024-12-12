package main

import (
	"gy-go-aes-server/handlers"
	"log"
	"net/http"
	"os"
)

func tokenAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		if token == "" {
			token = r.Header.Get("Authorization")
		}
		if token != ValidToken {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	if _, err := os.Stat(PackageDir); os.IsNotExist(err) {
		if err := os.MkdirAll(PackageDir, 0755); err != nil {
			log.Fatalf("Failed to create package directory: %v", err)
		}
	}

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/encrypt", handlers.EncryptHandler)
	http.HandleFunc("/decrypt", handlers.DecryptHandler)
	http.HandleFunc("/bundesliga", handlers.BundesLigaHandler)
	http.HandleFunc("/bundesligaguess", handlers.BundesLigaGuessHandler)
	http.Handle("/pip/", tokenAuth(http.HandlerFunc(PackageHandler)))
	http.Handle("/packages/", tokenAuth(http.StripPrefix("/packages/", http.HandlerFunc(StaticFileHandler))))

	log.Println("Starting server on :8888...")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
