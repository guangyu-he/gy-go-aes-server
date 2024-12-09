package bundesliga

import (
	"log"
	"os"
)

var ApiKey string

func init() {
	ApiKey = os.Getenv("FOOTBALL_DATA_API")
	if ApiKey == "" {
		log.Printf("FOOTBALL_DATA_API is not set")
		ApiKey = "adfac7e310f6495f99f1c38883718fd0"
	}
}
