package main

import (
	"encoding/json"
	"log"
	"toester/database"
	"toester/dto"
	"toester/model"

	"gorm.io/gorm"
)

func main() {
	db := database.Must()
	defer func() {
		selfDb, _ := db.DB()
		selfDb.Close()
	}()

	// database.Seeder(db)

	Page(db)

	_ = db
}

func Page(db *gorm.DB) {
	var players []*dto.PlayerUserResponse
	err := db.Debug().
		Model(model.PlayerModel{}).
		Scopes(func(tx *gorm.DB) *gorm.DB {
			return tx.Joins("INNER JOIN users u ON u.id = players.user_id").
				Select(
					"players.id AS id",
					"players.gender AS gender",
					"players.date AS date",
					"players.level AS level",
					"players.local_language AS local_language",
					"players.nationality AS nationality",
					"players.place_of_birth AS place_of_birth",
					"players.status AS status",
					"players.created_at AS created_at",

					"u.id AS user_id",
					"u.first_name AS first_name",
					"u.middle_name AS middle_name",
					"u.last_name AS last_name",
					"u.passport_number AS passport_number",
					"u.pin AS pin",
					"u.username AS username",
					"u.date_of_birth as date_of_birth",

					"COUNT(players.id) OVER() AS total",
				).Limit(2)

		}).
		Scan(&players).Error

	data, _ := json.MarshalIndent(players, " ", "    ")

	log.Println(string(data))

	log.Println(err)
}
