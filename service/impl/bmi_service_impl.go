package impl

import (
	"context"
	"fmt"

	"github.com/RizkiMufrizal/gofiber-clean-architecture/common"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/repository"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"github.com/go-redis/redis/v9"
)

func NewBMIServiceImpl(bmiRepository *repository.BMIRepository, cache *redis.Client) service.BMIService {
	return &bmiServiceImpl{BMIRepository: *bmiRepository, Cache: cache}
}

type bmiServiceImpl struct {
	repository.BMIRepository
	Cache *redis.Client
}

func (service *bmiServiceImpl) Create(ctx context.Context, bmiModel model.BMICreateModel) model.BMIModel {
	common.Validate(bmiModel)
	exception.Info("Create BMI")
	CalBMI := BMI(bmiModel.WeightKG, bmiModel.HightM)
	bmi := entity.BMI{
		Name:     bmiModel.Name,
		WeightKG: bmiModel.WeightKG,
		HightM:   bmiModel.HightM,
		BMI:      CalBMI,
	}
	fmt.Println(CalBMI)
	bmiSend := model.BMIModel{
		Name:     bmiModel.Name,
		WeightKG: bmiModel.WeightKG,
		HightM:   bmiModel.HightM,
		BMI:      CalBMI,
	}
	service.BMIRepository.Insert(ctx, bmi)
	return bmiSend
}
func BMI(weightKG, hightM float64) (bmi float64) {
	bmi = weightKG / (hightM * hightM)
	return bmi
}
