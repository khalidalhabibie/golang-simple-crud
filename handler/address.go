package handler

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/khalidalhabibie/depatu/model"
	"github.com/khalidalhabibie/depatu/repository"
)

type AddressHandler interface {
	AddAddress(*gin.Context)
	GetAddressbyUser(*gin.Context)
	GetAddressUserAdmin(*gin.Context)
	UpdateAddress(*gin.Context)
	DeleteAddress(*gin.Context)
}

type addressHandler struct {
	repo repository.AddressRepository
}

func NewAddressHandler() AddressHandler {
	return &addressHandler{
		repo: repository.NewAddressRepository(),
	}
}

func (h *addressHandler) AddAddress(ctx *gin.Context) {
	var address model.Address
	if err := ctx.ShouldBindJSON(&address); err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	const BearerSchema string = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BearerSchema):]
	username := ParsingTokentoUsername(tokenString)
	//ID := ParsingTokentoID(tokenString)
	address.Username = username
	fmt.Println(address)
	addAddress, err := h.repo.AddAddress(address)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, addAddress)

}

func (h *addressHandler) GetAddressbyUser(ctx *gin.Context) {
	var address model.Address
	if err := ctx.ShouldBindJSON(&address); err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	const BearerSchema string = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BearerSchema):]
	username := ParsingTokentoUsername(tokenString)
	//ID := ParsingTokentoID(tokenString)
	address.Username = username
	fmt.Println(address)
	addressuser, err := h.repo.GetAddressByUsername(address.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": ctx.Writer.Status(), "error": err.Error()})
		return

	}
	codeStatus := ctx.Writer.Status()
	ctx.JSON(http.StatusOK, gin.H{"Data": addressuser, "code": codeStatus})

}

func (h *addressHandler) GetAddressUserAdmin(ctx *gin.Context) {
	var address model.Address
	if err := ctx.ShouldBindJSON(&address); err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	addressuser, err := h.repo.GetAddressByUsername(address.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": ctx.Writer.Status(), "error": err.Error()})
		return

	}
	codeStatus := ctx.Writer.Status()
	ctx.JSON(http.StatusOK, gin.H{"Data": addressuser, "code": codeStatus})

}


func (h *addressHandler) UpdateAddress(ctx *gin.Context) {
	var address model.Address

	if err := ctx.ShouldBindJSON(&address); err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	const BearerSchema string = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BearerSchema):]

	
	username := ParsingTokentoUsername(tokenString)
	
	address.Username = username

	fmt.Println(address)
	data, err := h.repo.UpdateAddress(address)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	
	ctx.JSON(http.StatusOK, data)

}


func (h *addressHandler) DeleteAddress(ctx *gin.Context) {
	var address model.Address

	if err := ctx.ShouldBindJSON(&address); err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	const BearerSchema string = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BearerSchema):]

	
	username := ParsingTokentoUsername(tokenString)
	
	address.Username = username

	fmt.Println(address)
	data, err := h.repo.DeleteAddress(address)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	
	ctx.JSON(http.StatusOK, data)

}