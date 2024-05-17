package config

import (
	"fmt"
	"os"
	"strconv"
)

type database struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
	Params   string
}

var DB = &database{}

type bcrypt struct {
	Salt int
}

var Bcrypt = &bcrypt{}

type jwt struct {
	SECRET string
}

var JWT = &jwt{}

func init() {
	DB.Name = os.Getenv("DB_NAME")
	DB.Port = os.Getenv("DB_PORT")
	DB.Host = os.Getenv("DB_HOST")
	DB.User = os.Getenv("DB_USER")
	DB.Password = os.Getenv("DB_PASSWORD")
	DB.Params = os.Getenv("DB_PARAMS")

	saltStr := os.Getenv("BCRYPT_SALT")
	if salt, err := strconv.Atoi(saltStr); err != nil {
		Bcrypt.Salt = 8
	} else {
		Bcrypt.Salt = salt
	}

	JWT.SECRET = os.Getenv("JWT_SECRET")
}

func GetDBAdd() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s", DB.User, DB.Password, DB.Host, DB.Port, DB.Name, DB.Params)
}
