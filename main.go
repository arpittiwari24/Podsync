package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"main.go/api"
)


 func main() {
	fmt.Println("Jai Shree ram !!")
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", api.Handler)

	log.Fatal(app.Listen(":4002"))
}