package controllers

import "github.com/gofiber/fiber/v2"

//  show hello world
func Hello(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello, World!",
	})
}


//  calculate result
func MultiplyHandler(c *fiber.Ctx) error {
	type Request struct {
		Number1 int `json:"a"`
		Number2 int `json:"b"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	result := req.Number1 * req.Number2
	return c.JSON(fiber.Map{
		"result": result,
	})
}


