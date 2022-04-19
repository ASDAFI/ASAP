package farm



import (
	"gorm.io/gorm"
)

type Farm struct {
	gorm.Model
	FarmName string `gorm:"column:company_name;"`
}

func (Farm) TableName() string {
	return "farm"
}

