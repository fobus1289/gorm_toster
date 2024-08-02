package model

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	Id             int64           `json:"id" gorm:"primaryKey"`
	Pin            string          `json:"pin" gorm:"column:pin;uniqueIndex;not null"`
	Username       string          `json:"username" gorm:"column:username;not null"`
	Valid          string          `json:"valid" gorm:"column:valid;not null"`
	PassportNumber string          `json:"passportNumber" gorm:"column:passport_number;not null"`
	FirstName      string          `json:"firstName" gorm:"column:first_name;not null"`
	LastName       string          `json:"lastName" gorm:"column:last_name;not null"`
	MiddleName     string          `json:"middleName" gorm:"column:middle_name;not null"`
	DateOfBirth    string          `json:"dateOfBirth" gorm:"column:date_of_birth;not null"`
	UserType       string          `json:"userType" gorm:"column:user_type;not null"`
	LastVisitTime  *time.Time      `json:"lastVisitTime" gorm:"column:last_visit_time;"`
	CreatedAt      *time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt      *time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true,default:null"`
	DeletedAt      *gorm.DeletedAt `json:"-" swaggerignore:"true"`
} //@name UserModel

func (UserModel) TableName() string {
	return "users"
}
