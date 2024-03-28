package model

type BMIModel struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	WeightKG float64 `json:"weightKG"`
	HightM   float64 `json:"hightM"`
	BMI      float64 `json:"bmi"`
}

type BMICreateModel struct {
	Name     string  `json:"name" validate:"required"`
	WeightKG float64 `json:"weightKG" validate:"required"`
	HightM   float64 `json:"hightM" validate:"required"`
}
