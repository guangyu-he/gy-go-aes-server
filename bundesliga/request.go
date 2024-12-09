package bundesliga

import (
	"io"
	"log"
	"net/http"
)

func RequestGet(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	req.Header.Set("X-Auth-Token", ApiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Request failed with status：%d\n", resp.StatusCode)
	}

	return body, nil
}
