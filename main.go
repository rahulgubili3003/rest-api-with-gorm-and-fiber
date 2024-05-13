package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rahulgubili3003/rest-api-with-gorm-and-fiber/handler"
	"log"
	_ "net/http"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")

	// Routes
	api.Get("/ping", handler.PingHandler)
	api.Get("/getAccountDetails", handler.GetAccountDetailsHandler)
	api.Get("/getUserDetails", handler.GetUserDetails)
	api.Get("/getAccountBalance", handler.GetAccountBalance)
	api.Post("/createAccount", handler.CreateAccount)
	api.Post("/updateBalance", handler.UpdateBalance)

	// Listen on Port 3000
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
