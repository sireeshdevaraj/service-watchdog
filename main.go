package main

import (
	"github.com/sireeshdevaraj/service-watchdog/utils"
	"time"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	for {
		utils.SendNotification() // Checking is done inside the function.
		time.Sleep(time.Hour)
	}
}
