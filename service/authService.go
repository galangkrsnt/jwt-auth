package service

import (
	"errors"
	"jwt-auth/model"
	"jwt-auth/repository"
	"jwt-auth/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	RegisterService(u model.User) error
	LoginCheck(username string, password string) (string, error)
	GetUserByID(uid uint) (model.User, error)
}

type AuthService struct {
	userRepo repository.UserRepoInterface
}

func NewAuthService() AuthServiceInterface {
	return &AuthService{repository.NewUserRepo()}
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (a AuthService) GetUserByID(uid uint) (model.User, error) {

	u, err := a.userRepo.GetUserByID(uid)
	if err != nil {
		return u, errors.New("user not found")
	}

	return u, nil
}

func (a AuthService) LoginCheck(username string, password string) (string, error) {

	u, err := a.userRepo.Login(username, password)

	if err != nil {
		return " ", err
	}

	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (a AuthService) RegisterService(u model.User) error {
	err := a.userRepo.AddUser(u)
	if err != nil {
		return err
	}

	return nil

}
