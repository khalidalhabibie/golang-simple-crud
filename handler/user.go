package handler

import (
	"fmt"
	"net/http"

	"golang-simple-crud/model"
	"golang-simple-crud/repository"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	AddUser(*gin.Context)
	GetProfile(*gin.Context)
	SignInUser(*gin.Context)
	UpdateUser(*gin.Context)
	UpdatePassword(*gin.Context)
	GetAllUser(*gin.Context)
	UploadPhoto(*gin.Context)
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
	}
	dbUser, err := h.repo.GetUser(user.Username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "User Not Found"})
		return

	}
	if isTrue := comparePassword(dbUser.Password, user.Password); isTrue {
		fmt.Println("user before", dbUser.ID, dbUser.Username, dbUser.Admin)
		token := GenerateToken(dbUser.ID, dbUser.Username, dbUser.Admin)
		ctx.JSON(http.StatusOK, gin.H{"message": "Successfully", "token": token})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Wrong Password"})
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

	if err := ctx.ShouldBindJSON(&user); err != nil {
	}

	const BearerSchema string = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BearerSchema):]

	ID := ParsingTokentoID(tokenString)
	username := ParsingTokentoUsername(tokenString)
	user.Admin = false

	user.ID = uint(ID)
	user.Username = username
	hashPassword(&user.Password)
	fmt.Println(user)
	data, err := h.repo.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, data)

}

func (h *userHandler) UpdatePassword(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
	}

	const BearerSchema string = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BearerSchema):]

	ID := ParsingTokentoID(tokenString)
	username := ParsingTokentoUsername(tokenString)
	user.Admin = false

	user.ID = uint(ID)
	user.Username = username
	hashPassword(&user.Password)

	fmt.Println(user.Password)
	data, err := h.repo.UpdatePassword(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, data)

}

func (h *userHandler) GetAllUser(ctx *gin.Context) {
	user, err := h.repo.GetAllUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": ctx.Writer.Status(), "error": err.Error()})
		return

	}
	codeStatus := ctx.Writer.Status()
	ctx.JSON(http.StatusOK, gin.H{"Data": user, "code": codeStatus})

}

func (h *userHandler) UploadPhoto(ctx *gin.Context) {

	file, err := ctx.FormFile("photo")
	if err != nil {
		fmt.Println(err)
	}

	const BearerSchema string = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BearerSchema):]
	username := ParsingTokentoUsername(tokenString)
	file.Filename = username
	fmt.Println(file.Filename)

	err = ctx.SaveUploadedFile(file, file.Filename+".png")
	if err != nil {
		fmt.Println(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	log.Print(dir)
	var user model.User

	user.Image = file.Filename
	h.repo.UpdateUser(user)

	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
