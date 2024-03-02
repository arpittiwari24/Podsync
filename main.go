package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func main() {
	fmt.Println("Jai Shree ram !!")
	getSpotify()
	getYoutube()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/youtube-data", func(c *fiber.Ctx) error {
		return c.SendString("Youtube")
	})

	log.Fatal(app.Listen(":4002"))
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

	data , err := youtubeService.Search.List([]string{"snippet"}).Q("golang").Do()

	if err != nil {
		fmt.Println("Error fetching data")
	}
	// marshal,err := data.MarshalJSON()

	// if err != nil {
	// 	fmt.Println("error")
	// }
	
	object := data.Items[0]
	fmt.Println(object.Snippet.Title)
}