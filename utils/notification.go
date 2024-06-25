package utils

import (
	"fmt"
	"os"
)

func SendNotification() {
	for i := 0; i < len(Service); i++ {
		result, err := CheckStatus(Service[i])
		if err != nil {
			fmt.Println("Error in Command execution", err)
			continue
		}
		if result == true {
			continue
		} else {
			fmt.Println("The database is in-active:", Service[i])
			WebHookUrl := os.Getenv("WebHookUrl")
			PostToDiscord(WebHookUrl, Service[i])
		}
	}
}
