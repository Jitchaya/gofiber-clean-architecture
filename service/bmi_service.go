package service

import (
	"context"

	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
)

type BMIService interface {
	Create(ctx context.Context, model model.BMICreateModel) model.BMIModel
}
