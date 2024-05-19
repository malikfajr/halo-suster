package entity

import (
	"time"
)

type AddMedicalRecordPayload struct {
	UserId      string `json:"-"`
	UserNip     int    `json:"-"`
	PatientId   int    `json:"identityNumber" validate:"required,idNumber"`
	Symptoms    string `json:"symptoms" validate:"required,min=1,max=2000"`
	Medications string `json:"medications" validate:"required,min=1,max=2000"`
}

type MedicalRecordQueryParam struct {
	PatientId string `query:"identityDetail.identityNumber"`
	UserId    string `query:"createdBy.userId"`
	UserNip   string `query:"createdBy.nip"`
	CreatedAt string `query:"createdAt"`
	Limit     int    `query:"limit"`
	Offset    int    `query:"offset"`
}

type MedicalRecord struct {
	IdentityDetail interface{} `json:"identityDetail"`
	CreatedBy      interface{} `json:"createdBy"`
	Symptoms       string      `json:"symptoms"`
	Medications    string      `json:"medications"`
	CreatedAt      *time.Time  `json:"createdAt"`
}
