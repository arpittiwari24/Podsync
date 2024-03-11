package api

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type Youtube struct {
	VideoId string
	ThumbnailUrl string
	Title string
	PublishedAt string
}

func Handler(c *fiber.Ctx) error {
	data , err := getYoutube()

	if err != nil {
		return c.SendString("Error fetching podcasts")
	}

	return c.Render("index", fiber.Map{ "data": data })
}

func getYoutube() ([]Youtube, error) {
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

	data, err := youtubeService.Search.List([]string{"snippet"}).MaxResults(50).Q("Trending Podcasts").Do()

	if err != nil {
		fmt.Println("Error fetching data")
	}
	var videos []Youtube

	for _, item := range data.Items {
		video := Youtube{
			VideoId:      item.Id.VideoId,
			ThumbnailUrl: item.Snippet.Thumbnails.Default.Url,
			Title:        item.Snippet.Title,
			PublishedAt:  item.Snippet.PublishedAt,
		}
		videos = append(videos, video)
	}

	return videos, nil
}

// func getSpotify() {
// 	// err := godotenv.Load()
// 	// if err != nil {
// 	// 	fmt.Println("Error loading .env file")
// 	// }

// 	// clientId := os.Getenv("SPOTIFY_CLIENT_ID")
// 	// clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

// 	// client := spotify.NewClient(client *http.Client) client 

// }