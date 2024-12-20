package farm


import (
	"gorm.io/gorm"
)

type Farm struct {
	gorm.Model
	FarmName string 	`gorm:"column:company_name;"`
	Production string 	`gorm:"column:production;"`
}


type CreateFarmParameters struct {
	FarmName string
	Production string
}

func NewFarm(params CreateFarmParameters) (*Farm, error) {
	farm := &Farm{
		FarmName:  params.FarmName,
		Production: params.Production,
	}
	return farm, farm.validateForCreateNewInstance()
}

func (Farm) TableName() string {
	return "farm"
}


