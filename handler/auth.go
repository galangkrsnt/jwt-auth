package handler

import (
	"jwt-auth/model"
	"jwt-auth/service"
	"jwt-auth/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandlerInterface interface {
	RegisterHandler(c *gin.Context)
	LoginHandler(c *gin.Context)
	CurrentUser(c *gin.Context)
}

type AuthHandler struct {
	userService service.AuthServiceInterface
}

func NewAuthHandler() AuthHandlerInterface {
	return &AuthHandler{
		userService: service.NewAuthService(),
	}
}

func (a AuthHandler) CurrentUser(c *gin.Context) {

	user_id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := a.userService.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func (a AuthHandler) LoginHandler(c *gin.Context) {
	var loginInput model.User

	err := c.ShouldBindJSON(&loginInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := a.userService.LoginCheck(loginInput.Username, loginInput.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (a AuthHandler) RegisterHandler(c *gin.Context) {

	var input model.User

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = a.userService.RegisterService(input)
	if err != nil {
		log.Fatal("error save to db")
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registration Success"})

}
