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

type s3 struct {
	ID          string
	SECRET_KEY  string
	BUCKET_NAME string
	REGION      string
}

var S3 = &s3{}

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

	S3.ID = os.Getenv("S3_ID")
	S3.SECRET_KEY = os.Getenv("S3_SECRET_KEY")
	S3.BUCKET_NAME = os.Getenv("S3_BUCKET_NAME")
	S3.REGION = os.Getenv("S3_REGION")
}

func GetDBAdd() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s", DB.User, DB.Password, DB.Host, DB.Port, DB.Name, DB.Params)
}
