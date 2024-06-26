package utils

import (
	"fmt"
)

func SendNotification(webHookUrl string) {
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
			PostToDiscord(webHookUrl, Service[i])
		}
	}
}
