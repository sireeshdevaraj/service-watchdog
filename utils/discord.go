package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Thumbnail struct {
	Url string `json:"url"`
}

type Embed struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	Color       string    `json:"color"`
	Thumbnail   Thumbnail `json:"thumbnail"`
}

type Discord struct {
	Content string  `json:"content"`
	Embeds  []Embed `json:"embeds"`
}

func PostToDiscord(WebHookUrl string, Service string) {
	thumbnail := Thumbnail{
		Url: "https://c.tenor.com/usccWzYPsGUAAAAC/anime-discord.gif",
	}
	embeds := []Embed{{
		Title:       "Services are down, might need to restart services!! ",
		Description: "Service Name: " + Service,
		Url:         "https://kuuhaku.space",
		Color:       "5758966",
		Thumbnail:   thumbnail,
	}}
	discordData := Discord{
		Content: RoleToPing,
		Embeds:  embeds,
	}
	jsonData, err := json.Marshal(discordData)

	if err != nil {
		fmt.Println("Error in Json Encoding", err)
		return
	}
	resp, err := http.Post(WebHookUrl,
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Println("Error Occurred when sending request", err)
		return
	}
	defer resp.Body.Close()
}
