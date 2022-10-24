package repository

import (
	"errors"
	"html"
	"jwt-auth/config"
	"jwt-auth/model"
	"log"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type UserRepoInterface interface {
	AddUser(u model.User) error
	Login(username string, password string) (model.User, error)
	GetUserByID(uid uint) (model.User, error)
}

type UserRepo struct {
}

func NewUserRepo() UserRepoInterface {
	return &UserRepo{}
}

func (d UserRepo) GetUserByID(uid uint) (model.User, error) {

	var u model.User

	err := config.DB.First(&u, uid).Error

	if err != nil {
		return u, errors.New("user not found")
	}

	u.Password = ""

	return u, nil

}

func (d UserRepo) Login(username string, password string) (model.User, error) {
	u := model.User{}

	err := config.DB.Model(model.User{}).Where("username = ? ", username).Take(&u).Error

	if err != nil {
		return u, err
	}

	return u, nil
}

func BeforeSaveUser(u model.User) model.User {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal("Failed hashing password")
	}

	u.Password = string(hashedPassword)

	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return u
}

func (d UserRepo) AddUser(u model.User) error {

	config.AutoMigrateDataBase()

	hashed := BeforeSaveUser(u)
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading env file")
	}

	if err != nil {
		log.Fatal("Failed Connect to Database")
	}

	err = config.DB.Create(&hashed).Error
	if err != nil {
		return err
	}

	return nil

}
