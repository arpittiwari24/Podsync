package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func main() {
	fmt.Println("Jai Shree ram !!")
	getSpotify()
	getYoutube()
}

func getSpotify() {
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println("Error loading .env file")
	// }

	// clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	// clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	// client := spotify.NewClient(client *http.Client) client 

}

func getYoutube() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	api_key := os.Getenv("API_KEY")
	ctx := context.Background()

	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(api_key))
	if err != nil {
		fmt.Println("Error creating youtube service")
	}
	fmt.Println(youtubeService)
	data , err := youtubeService.Search.List([]string{"golang"}).Q("golang").Do()

	if err != nil {
		fmt.Println("Error fetching data")
	}
	fmt.Println(data)
}