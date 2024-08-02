package database

import (
	"toester/model"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) {
	for i := 0; i < 10000; i++ {
		user := model.UserModel{
			Pin:            gofakeit.Password(true, true, true, false, false, 10),
			Username:       gofakeit.Username(),
			Valid:          "true",
			PassportNumber: gofakeit.Password(true, true, true, false, false, 10),
			FirstName:      gofakeit.FirstName(),
			LastName:       gofakeit.LastName(),
			MiddleName:     gofakeit.FirstName(),
			DateOfBirth:    gofakeit.Date().Format("2006-01-02"),
			UserType:       "player",
		}

		db.Create(&user)

		player := model.PlayerModel{
			UserId:         user.Id,
			Gender:         gofakeit.Gender(),
			Date:           gofakeit.Date().Format("2006-01-02"),
			CountryOfBirth: gofakeit.Country(),
			PlaceOfBirth:   gofakeit.City(),
			LocalLanguage:  gofakeit.Language(),
			Nationality:    gofakeit.Country(),
			Status:         "active",
			Level:          "beginner",
		}

		db.Create(&player)
	}
}
