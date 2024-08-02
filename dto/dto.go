package dto

import "time"

type PlayerUserResponse struct {
	Id             int64      `json:"id"`
	UserId         int64      `json:"userId"`
	Gender         string     `json:"gender"`
	Date           string     `json:"date"`
	CountryOfBirth string     `json:"countryOfBirth"`
	PlaceOfBirth   string     `json:"placeOfBirth"`
	LocalLanguage  string     `json:"localLanguage"`
	Nationality    string     `json:"nationality"`
	Status         string     `json:"status"`
	Level          string     `json:"level"`
	CreatedAt      *time.Time `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	Pin            string     `json:"pin"`
	Username       string     `json:"username"`
	PassportNumber string     `json:"passportNumber"`
	FirstName      string     `json:"firstName"`
	LastName       string     `json:"lastName"`
	MiddleName     string     `json:"middleName"`
	DateOfBirth    string     `json:"dateOfBirth"`
	Total          int64      `json:"-"`
}
