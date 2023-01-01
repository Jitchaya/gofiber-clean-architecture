package controller

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/middleware"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"github.com/gofiber/fiber/v2"
)

type TransactionDetailController struct {
	service.TransactionDetailService
}

func NewTransactionDetailController(transactionDetailService *service.TransactionDetailService) *TransactionDetailController {
	return &TransactionDetailController{TransactionDetailService: *transactionDetailService}
}

func (controller TransactionDetailController) Route(app *fiber.App) {
	app.Get("/v1/api/transaction-detail/:id", middleware.AuthenticateJWT("ROLE_USER"), controller.FindById)
}

// FindById func gets one exists transaction detail.
// @Description Get one exists transaction detail.
// @Summary get one exists transaction detail
// @Tags Transaction Detail
// @Accept json
// @Produce json
// @Param id path string true "Transaction Detail Id"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/transaction-detail/{id} [get]
func (controller TransactionDetailController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	result := controller.TransactionDetailService.FindById(c.Context(), id)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
