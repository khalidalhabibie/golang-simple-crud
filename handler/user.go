package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khalidalhabibie/depatu/model"
	"github.com/khalidalhabibie/depatu/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	AddUser(*gin.Context)
	GetProfile(*gin.Context)
	SignInUser(*gin.Context)
	UpdateUser(*gin.Context)
	//DeleteUser(*gin.Context)
	//GetTask(*gin.Context)
}

type userHandler struct {
	repo repository.UserRepository
}

func NewUserHandler() UserHandler {
	return &userHandler{
		repo: repository.NewUserRepository(),
	}
}

func hashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPass)
}

func comparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}

func (h *userHandler) SignInUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	dbUser, err := h.repo.GetUser(user.Username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "User Not Found"})
		return

	}
	if isTrue := comparePassword(dbUser.Password, user.Password); isTrue {
		fmt.Println("user before", dbUser.ID, dbUser.Username)
		token := GenerateToken(dbUser.ID, dbUser.Username)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully", "token": token})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Wrong Password"})
	return

}

func (h *userHandler) AddUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashPassword(&user.Password)
	user, err := h.repo.AddUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	user.Password = ""
	ctx.JSON(http.StatusOK, user)

}
func (h *userHandler) GetProfile(ctx *gin.Context) {
	const BearerSchema string = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BearerSchema):]
	username := ParsingTokentoUsername(tokenString)

	user, err := h.repo.GetUser(username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, user)

}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	var user model.User

	err := ctx.ShouldBindJSON(&user)
	fmt.Println(user)
	const BearerSchema string = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BearerSchema):]
	ID := ParsingTokentoID(tokenString)

	user.ID = uint(ID)
	/*
		user, err := h.repo.UpdateUser(user)
	*/
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, user)

}
