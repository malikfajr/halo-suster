package entity

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

type ITStaffLogin struct {
	Nip      int    `json:"nip" validate:"required"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type ITStaffResponse struct {
	ID          string `json:"userId"`
	Nip         int    `json:"nip"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
