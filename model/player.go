package model

import (
	"time"

	"gorm.io/gorm"
)

type PlayerModel struct {
	Id             int64          `json:"id"`
	UserId         int64          `json:"userId" gorm:"unique;not null"`
	Gender         string         `json:"gender"`
	Date           string         `json:"date"`
	CountryOfBirth string         `json:"countryOfBirth"`
	PlaceOfBirth   string         `json:"placeOfBirth"`
	LocalLanguage  string         `json:"localLanguage"`
	Nationality    string         `json:"nationality"`
	Status         string         `json:"status"`
	Level          string         `json:"level"`
	User           *UserModel     `json:"user,omitempty" gorm:"foreignKey:UserId"`
	CreatedAt      *time.Time     `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt      *time.Time     `json:"updatedAt,omitempty" gorm:"autoUpdateTime:true"`
	DeletedAt      gorm.DeletedAt `json:"-" swaggerignore:"true"`
} //@name PlayerModel

func (PlayerModel) TableName() string {
	return "players"
}
