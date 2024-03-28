package controller

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"github.com/gofiber/fiber/v2"
)

type BMIController struct {
	service.BMIService
	configuration.Config
}

func NewBMIController(bmiService *service.BMIService, config configuration.Config) *BMIController {
	return &BMIController{BMIService: *bmiService, Config: config}
}
func (controller BMIController) Route(app *fiber.App) {
	app.Post("/v1/api/bmi", controller.Create)
}

// Create func create bmi.
// @Description create bmi.
// @Summary create bmi
// @Tags bmi
// @Accept json
// @Produce json
// @Param request body model.ProductCreateOrUpdateModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/bmi [post]
func (controller BMIController) Create(c *fiber.Ctx) error {
	var request model.BMICreateModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	response := controller.BMIService.Create(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}
