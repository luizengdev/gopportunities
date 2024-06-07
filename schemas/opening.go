package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string `json:"role" gorm:"not null"`
	Company  string `json:"company" gorm:"not null"`
	Location string `json:"location" gorm:"not null"`
	Remote   bool   `json:"remote" gorm:"not null"`
	Link     string `json:"link" gorm:"not null"`
	Salary   int64  `json:"salary" gorm:"not null"`
}
