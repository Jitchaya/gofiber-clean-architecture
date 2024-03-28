package repository

import (
	"context"

	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
)

type BMIRepository interface {
	Insert(ctx context.Context, bmi entity.BMI)
}
