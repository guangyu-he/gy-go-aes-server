package handlers

import (
	"encoding/json"
	"fmt"
	"gy-go-aes-server/aes"
	"gy-go-aes-server/bundesliga"
	"io"
	"log"
	"net/http"
	"strconv"
)

type RequestData struct {
	Key  string `json:"key"`
	Text string `json:"text"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	log.Println(r.URL.Path, r.URL.Query(), r.Method, r.RemoteAddr)

	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err := fmt.Fprintf(w, "Hello, World!")
	if err != nil {
		return
	}
}

func EncryptHandler(w http.ResponseWriter, r *http.Request) {

	log.Println(r.URL.Path, r.URL.Query(), r.Method, r.RemoteAddr)

	type ResponseData struct {
		Text          string `json:"text"`
		EncryptedText string `json:"encrypted_text"`
		Status        string `json:"status"`
	}

	if r.Method != http.MethodPost {
		log.Println("Method not allowed")
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body")
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
		log.Println("Error parsing JSON")
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	if requestData.Key == "" {
		log.Println("AES Key is empty")
		http.Error(w, "AES Key is empty", http.StatusBadRequest)
		return
	}

	encryptedText, err := aes.AESEncrypt([]byte(requestData.Text), []byte(requestData.Key))
	if err != nil {
		log.Println("Error encrypting text")
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
		log.Println("Error generating JSON response")
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println("Error writing JSON response")
		http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
		return
	}
}

func DecryptHandler(w http.ResponseWriter, r *http.Request) {

	log.Println(r.URL.Path, r.URL.Query(), r.Method, r.RemoteAddr)

	type ResponseData struct {
		Text          string `json:"text"`
		DecryptedText string `json:"decrypted_text"`
		Status        string `json:"status"`
	}

	if r.Method != http.MethodPost {
		log.Println("Method not allowed")
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body")
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
		log.Println("Error parsing JSON")
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	if requestData.Key == "" {
		log.Println("AES Key is empty")
		http.Error(w, "AES Key is empty", http.StatusBadRequest)
		return
	}

	decryptedText, err := aes.AESDecrypt(requestData.Text, []byte(requestData.Key))
	if err != nil {
		log.Println("Error decrypting text")
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
		log.Println("Error generating JSON response")
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
	log.Println(r.URL.Path, r.URL.Query(), r.Method, r.RemoteAddr)

	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	MatchDay := r.URL.Query().Get("matchday")
	//MatchDayInt, err := strconv.Atoi(MatchDay)
	//if err != nil {
	//	log.Println("Invalid MatchDay")
	//	http.Error(w, "Invalid MatchDay", http.StatusBadRequest)
	//	return
	//}
	result, err := bundesliga.MatchDayInfo(MatchDay)
	if err != nil {
		log.Printf("Error fetching matchday %s: %s\n", MatchDay, err)
		http.Error(w, fmt.Sprintf("Error fetching matchday %s: %s", MatchDay, err), http.StatusInternalServerError)
		return
	}
	log.Printf("MatchDay %s - Results fetched\n", MatchDay)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(result))
	if err != nil {
		return
	}
}

func BundesLigaGuessHandler(w http.ResponseWriter, r *http.Request) {

	type ResponseData struct {
		Json   interface{} `json:"json"`
		Status string      `json:"status"`
	}

	log.Println(r.URL.Path, r.URL.Query(), r.Method, r.RemoteAddr)

	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	TeamID := r.URL.Query().Get("teamid")
	TeamIDInt, err := strconv.Atoi(TeamID)
	if err != nil {
		log.Println("Invalid TeamID")
		http.Error(w, "Invalid TeamID", http.StatusBadRequest)
		return
	}
	nextGameResult, err := bundesliga.NextGame(TeamIDInt)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("Error fetching next game for team %s: %s", TeamID, err), http.StatusInternalServerError)
	}
	log.Printf("Next game for team %s - Prediction fetched\n", TeamID)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	responseData := ResponseData{
		Json:   nextGameResult,
		Status: "success",
	}
	jsonResponse, err := json.Marshal(responseData)
	if err != nil {
		log.Println("Error generating JSON response")
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		return
	}

}
