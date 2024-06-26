package main

import (
	"github.com/joho/godotenv"
	"github.com/sireeshdevaraj/service-watchdog/utils"
	"os"
	"time"
)

func main() {
	godotenv.Load()
	webHookUrl := os.Getenv("WebhookUrl")
	fmt.Println("Application is online...")
	for {
		utils.SendNotification(webHookUrl) // Checking is done inside the function.
		time.Sleep(time.Hour)
	}
}
