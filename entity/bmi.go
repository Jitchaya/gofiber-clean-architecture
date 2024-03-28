package entity

import "github.com/google/uuid"

type BMI struct {
	Id       uuid.UUID `gorm:"primaryKey;column:bmi_id;type:varchar(36)"`
	Name     string    `gorm:"primaryKey;column:name;type:varchar(100)"`
	WeightKG float64   `gorm:"column:weightKG;type:float"`
	HightM   float64   `gorm:"column:hightM;type:float"`
	BMI      float64   `gorm:"column:bmi;type:float"`
}

func (BMI) TableName() string {
	return "tb_bmi"
}
