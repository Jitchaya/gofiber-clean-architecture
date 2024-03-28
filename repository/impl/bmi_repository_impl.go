package impl

import (
	"context"

	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewBMIRepositoryImpl(DB *gorm.DB) repository.BMIRepository {
	return &BMIRepositoryImpl{DB: DB}
}

type BMIRepositoryImpl struct {
	*gorm.DB
}

func (repository *BMIRepositoryImpl) Insert(ctx context.Context, bmi entity.BMI) {
	bmi.Id = uuid.New()
	err := repository.DB.WithContext(ctx).Create(&bmi).Error
	exception.PanicLogging(err)
}
