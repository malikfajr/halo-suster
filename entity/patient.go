package entity

import "time"

type Patient struct {
	IdNumber    int        `json:"identityNumber" db:"id"`
	PhoneNumber string     `json:"phoneNumber" db:"phone_number"`
	Name        string     `json:"name" db:"name"`
	BirthDate   *time.Time `json:"birthDate" db:"birth_date"`
	Gender      string     `json:"gender" db:"gender"`
	ImageCard   string     `json:"identityCardScanImg" db:"card_img"`
	CreatedAt   *time.Time `json:"createdAt" db:"created_at"`
}

type AddPatientPayload struct {
	IdNumber    int        `json:"identityNumber" validate:"required,idNumber"`
	PhoneNumber string     `json:"phoneNumber" validate:"required,id_phone"`
	Name        string     `json:"name" validate:"required,min=3,max=30"`
	BirthDate   *time.Time `json:"birthDate" validate:"required"`
	Gender      string     `json:"gender" validate:"required,oneof=male female"`
	ImageCard   string     `json:"identityCardScanImg" validate:"required,imageUrl"`
}

type PatientQueryParam struct {
	IdNumber    string `query:"idNumber"`
	Name        string `query:"name"`
	PhoneNumber string `query:"phoneNumber"`
	CreatedAt   string `query:"createdAt"`
	Limit       int    `query:"limit"`
	Offset      int    `query:"offset"`
}
