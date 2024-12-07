package handlers

import (
	"encoding/json"
	"fmt"
	"gy-go-aes-server/aes"
	"gy-go-aes-server/bundesliga"
	"io"
	"net/http"
	"strconv"
)

type RequestData struct {
	Key  string `json:"key"`
	Text string `json:"text"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err := fmt.Fprintf(w, "Hello, World!")
	if err != nil {
		return
	}
}

func EncryptHandler(w http.ResponseWriter, r *http.Request) {

	type ResponseData struct {
		Text          string `json:"text"`
		EncryptedText string `json:"encrypted_text"`
		Status        string `json:"status"`
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(r.Body)

	var requestData RequestData
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	if requestData.Key == "" {
		http.Error(w, "AES Key is empty", http.StatusBadRequest)
		return
	}

	encryptedText, err := aes.AESEncrypt([]byte(requestData.Text), []byte(requestData.Key))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encrypting text: %s", err), http.StatusInternalServerError)
		return
	}

	responseData := ResponseData{
		Text:          requestData.Text,
		EncryptedText: encryptedText,
		Status:        "success",
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	jsonResponse, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
		return
	}
}

func DecryptHandler(w http.ResponseWriter, r *http.Request) {

	type ResponseData struct {
		Text          string `json:"text"`
		DecryptedText string `json:"decrypted_text"`
		Status        string `json:"status"`
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(r.Body)

	var requestData RequestData
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	if requestData.Key == "" {
		http.Error(w, "AES Key is empty", http.StatusBadRequest)
		return
	}

	decryptedText, err := aes.AESDecrypt(requestData.Text, []byte(requestData.Key))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encrypting text: %s", err), http.StatusInternalServerError)
		return
	}

	responseData := ResponseData{
		Text:          requestData.Text,
		DecryptedText: decryptedText,
		Status:        "success",
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	jsonResponse, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
		return
	}

}

func BundesLigaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 从request中获取matchday参数
	matchday := r.URL.Query().Get("matchday")
	// 将matchday转换为int类型
	// 如果转换失败，返回错误信息
	matchdayInt, err := strconv.Atoi(matchday)
	if err != nil {
		http.Error(w, "Invalid matchday", http.StatusBadRequest)
		return
	}
	result := bundesliga.MatchInfo(matchdayInt)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err = w.Write([]byte(result))
	if err != nil {
		return
	}
}
