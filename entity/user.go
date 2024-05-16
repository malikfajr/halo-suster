package entity

import "time"

type User struct {
	ID        string     `json:"userId"`
	Nip       int        `json:"nip"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"createdAt"`
}

type UserParam struct {
	UserId    string `query:"userId"`
	Limit     int    `query:"limit"`
	Offset    int    `query:"offset"`
	Name      string `query:"name"`
	Nip       string `query:"nip"`
	Role      string `query:"role"`
	CreatedAt string `query:"createdAt"`
}

type ITStaff struct {
	ID       string
	Nip      int
	Password string
	Name     string
}

type ITStaffRegister struct {
	Nip      int    `json:"nip" validate:"required"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type AuthLogin struct {
	Nip      int    `json:"nip" validate:"required"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type AuthResponse struct {
	ID          string `json:"userId"`
	Nip         int    `json:"nip"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type NurseResponse struct {
	UserId string `json:"userId"`
	Nip    int    `json:"nip"`
	Name   string `json:"name"`
}

type AddNursePayload struct {
	Id                  string `json:"-"`
	Nip                 int    `json:"nip" validate:"required"`
	Name                string `json:"name" validate:"required,min=5,max=50"`
	IdentityCardScanImg string `json:"identityCardScanImg" validate:"required,url"`
}
