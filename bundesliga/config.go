package bundesliga

import (
	"log"
	"os"
)

var ApiKey string

func init() {
	ApiKey = os.Getenv("FOOTBALL_DATA_API")
	if ApiKey == "" {
		log.Fatalf("FOOTBALL_DATA_API is not set")
	}
}
