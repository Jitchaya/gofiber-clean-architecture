package controller

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"github.com/gofiber/fiber/v2"
)

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: *userService}
}

type UserController struct {
	service.UserService
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/v1/api/authentication", controller.Authentication)
}

// Authentication func Authenticate user.
// @Description authenticate user.
// @Summary authenticate user
// @Tags Authenticate user
// @Accept json
// @Produce json
// @Param request body model.UserModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/authentication [post]
func (controller UserController) Authentication(c *fiber.Ctx) error {
	var request model.UserModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	result := controller.UserService.Authentication(c.Context(), request)
	var userRoles []map[string]interface{}
	for _, userRole := range result.UserRoles {
		userRoles = append(userRoles, map[string]interface{}{
			"role": userRole.Role,
		})
	}
	tokenJwtResult := configuration.GenerateToken(result.Username, userRoles)
	resultWithToken := map[string]interface{}{
		"token":    tokenJwtResult,
		"username": result.Username,
		"role":     userRoles,
	}
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    resultWithToken,
	})
}
