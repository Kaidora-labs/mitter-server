package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaidora-labs/mitter-server/database"
	"github.com/kaidora-labs/mitter-server/models"
)

func PostUserHandler(c *fiber.Ctx) error {
	body := new(models.CreateUserDTO)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	newUser := models.User{
		Id:           body.Id,
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		PhoneNumber:  body.PhoneNumber,
		EmailAddress: body.EmailAddress,
	}

	userModel := models.NewUserModel(database.DB)
	user, err := userModel.Save(&newUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User created successfully",
		"data":    user,
	})
}

func GetUserHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	userModel := models.NewUserModel(database.DB)
	user, err := userModel.Find(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User retrieved successfully",
		"data":    user,
	})
}

func GetUsersHandler(c *fiber.Ctx) error {
	userModel := models.NewUserModel(database.DB)
	users, err := userModel.FindAll()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Users retrieved successfully",
		"data":    users,
	})
}
