package handler

import "github.com/gofiber/fiber/v2"

func PingHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, Gorm!",
	})
}

func GetAccountDetailsHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Getting Account Details",
	})
}

func GetUserDetails(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "User!",
	})
}

func GetAccountBalance(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Acc balance",
	})
}

func CreateAccount(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Create Acc",
	})
}

func UpdateBalance(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Update Balance",
	})
}
